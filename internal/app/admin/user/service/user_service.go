package user

import (
	enterpriseRepo "Synapse/internal/app/admin/enterprise/repository"
	security "Synapse/internal/app/admin/pkg/security"
	ruleRepo "Synapse/internal/app/admin/rule/repository"
	userModel "Synapse/internal/app/admin/user/model"
	userRepo "Synapse/internal/app/admin/user/repository"
	"context"
	"fmt"
	"strings"
)

// service implementa a interface Service
type service struct {
	userRepo       userRepo.Repository
	enterpriseRepo enterpriseRepo.Repository
	ruleRepo       ruleRepo.Repository
}

// NewService cria uma nova instância do serviço de usuário
func NewService(
	uRepo userRepo.Repository,
	eRepo enterpriseRepo.Repository,
	rRepo ruleRepo.Repository,
) Service {
	return &service{
		userRepo:       uRepo,
		enterpriseRepo: eRepo,
		ruleRepo:       rRepo,
	}
}

// Create valida e cria um novo usuário
func (s *service) Create(user *userModel.User) (*userModel.User, error) {
	// Verifica se a empresa existe
	_, err := s.enterpriseRepo.ReadByID(user.EnterpriseID)
	if err != nil {
		return nil, fmt.Errorf("empresa com ID %d não encontrada", user.EnterpriseID)
	}

	// Verifica se a regra existe
	_, err = s.ruleRepo.FindRuleByID(user.RuleID)
	if err != nil {
		return nil, fmt.Errorf("regra com ID %d não encontrada", user.RuleID)
	}

	hashed, err := security.HashPassword(user.Senha)
	if err != nil {
		return nil, fmt.Errorf("erro ao criptografar senha")
	}
	user.Senha = hashed
	// Cria o usuário
	return s.userRepo.Create(user)
}

// ReadAllUser retorna todos os usuários de uma empresa com paginação.
func (s *service) ReadAllUser(enterpriseID int64, page int64) (*[]userModel.User, error) {
	// Opcional: validar se a empresa existe antes de listar
	_, err := s.enterpriseRepo.ReadByID(enterpriseID)
	if err != nil {
		return nil, fmt.Errorf("empresa com ID não encontrada")
	}

	return s.userRepo.ReadAllUser(enterpriseID, page)
}

// ReadByEmail retorna um usuário com base no e-mail.
func (s *service) ReadByEmail(email string) (*userModel.User, error) {
	user, err := s.userRepo.ReadByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("usuário com e-mail não encontrado")
	}
	return user, nil
}

// UpdateUserByID atualiza os dados de um usuário com base no ID
func (s *service) UpdateUserByID(userID int64, updated *userModel.User) (*userModel.User, error) {
	existingUser, err := s.userRepo.ReadByID(userID)
	if err != nil {
		return nil, fmt.Errorf("usuário com ID %d não encontrado", userID)
	}

	// --- EMAIL ---
	email := strings.TrimSpace(updated.Email)
	if email != "" && email != existingUser.Email {
		userWithEmail, err := s.userRepo.ReadByEmail(email)
		if err == nil && userWithEmail.ID != userID {
			return nil, fmt.Errorf("o e-mail informado já está em uso por outro usuário")
		}
		existingUser.Email = email
	}

	// --- NOME ---
	nome := strings.TrimSpace(updated.Nome)
	if nome != "" && nome != existingUser.Nome {
		existingUser.Nome = nome
	}

	// --- NÚMERO ---
	numero := strings.TrimSpace(updated.Numero)
	if numero != "" && numero != existingUser.Numero {
		existingUser.Numero = numero
	}

	// --- RULE_ID ---
	if updated.RuleID != 0 && updated.RuleID != existingUser.RuleID {
		_, err := s.ruleRepo.FindRuleByID(updated.RuleID)
		if err != nil {
			return nil, fmt.Errorf("regra com ID %d não encontrada", updated.RuleID)
		}
		existingUser.RuleID = updated.RuleID
	}

	// --- ENTERPRISE_ID ---
	if updated.EnterpriseID != 0 && updated.EnterpriseID != existingUser.EnterpriseID {
		_, err := s.enterpriseRepo.ReadByID(updated.EnterpriseID)
		if err != nil {
			return nil, fmt.Errorf("empresa com ID %d não encontrada", updated.EnterpriseID)
		}
		existingUser.EnterpriseID = updated.EnterpriseID
	}

	// --- SENHA ---
	senha := strings.TrimSpace(updated.Senha)
	if senha != "" {
		if err := security.ComparePassword(existingUser.Senha, senha); err != nil {
			// Senha é diferente da atual, então criptografa e atualiza
			hashed, err := security.HashPassword(senha)
			if err != nil {
				return nil, fmt.Errorf("erro ao criptografar senha")
			}
			existingUser.Senha = hashed
		}
	}

	// Atualiza no banco
	return s.userRepo.UpdateUserByID(userID, existingUser)
}

func (s *service) DeleteUserByID(userID int64) error {
	return s.userRepo.DeleteUserByID(userID)
}

// 🔐 Autentica e gera token JWT
func (s *service) CreateTokenUser(email string, senha string) (*userModel.User, string, error) {
	ctx := context.Background()

	user, err := s.userRepo.ValidateCredentials(ctx, email, senha)
	if err != nil {
		return nil, "", fmt.Errorf("credenciais inválidas: %w", err)
	}

	// Verifica token válido antes de gerar
	existingToken, err := s.userRepo.GetValidToken(ctx, user.ID)
	if err != nil {
		return nil, "", fmt.Errorf("erro ao verificar token existente: %w", err)
	}
	if existingToken != "" {
		return user, existingToken, nil
	}

	// Se não existe, gera e salva novo
	tokenData, err := security.GenerateToken(user.ID)
	if err != nil {
		return nil, "", fmt.Errorf("erro ao gerar token: %w", err)
	}

	// ✅ Agora passando também o tokenData.ExpiresAt
	err = s.userRepo.SaveToken(ctx, user.ID, tokenData.Token, tokenData.ExpiresAt)
	if err != nil {
		return nil, "", fmt.Errorf("erro ao salvar token no banco: %w", err)
	}

	return user, tokenData.Token, nil
}
