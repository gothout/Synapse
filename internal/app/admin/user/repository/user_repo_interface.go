package user

import (
	user "Synapse/internal/app/admin/user/model"
)

type Repository interface {
	Create(user *user.User) (*user.User, error)
	ReadAllUser(enterpriseId, page int64) (*[]user.User, error)
	ReadByEmail(email string) (*user.User, error)
	ReadByID(userID int64) (*user.User, error)
	UpdateUserByID(UserID int64, updated *user.User) (*user.User, error)
}
