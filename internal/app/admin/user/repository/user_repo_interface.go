package user

import (
	user "Synapse/internal/app/admin/user/model"
	"context"
	"time"
)

type Repository interface {
	Create(user *user.User) (*user.User, error)
	ReadAllUser(enterpriseId, page int64) (*[]user.User, error)
	ReadByEmail(email string) (*user.User, error)
	ReadByID(userID int64) (*user.User, error)
	UpdateUserByID(UserID int64, updated *user.User) (*user.User, error)
	DeleteUserByID(UserID int64) error
	// Autenticação
	ValidateCredentials(ctx context.Context, email, senha string) (*user.User, error)
	// Token
	SaveToken(ctx context.Context, userID int64, token string, expireAt time.Time) error
	GetValidToken(ctx context.Context, userID int64) (string, error)

	//Read()
	//Update()
	//Delete()
}
