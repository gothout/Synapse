package user

import (
	enterpriseRepo "Synapse/internal/app/admin/enterprise/repository"
	security "Synapse/internal/app/admin/pkg/security"
	ruleRepo "Synapse/internal/app/admin/rule/repository"
	userModel "Synapse/internal/app/admin/user/model"
	userRepo "Synapse/internal/app/admin/user/repository"
	"fmt"
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
