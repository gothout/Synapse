package user

import user "Synapse/internal/app/admin/user/model"

type Service interface {
	Create(user *user.User) (*user.User, error)
	ReadAllUser(enterpriseId, page int64) (*[]user.User, error)
	ReadByEmail(email string) (*user.User, error)
	UpdateUserByID(UserID int64, updated *user.User) (*user.User, error)
	DeleteUserByID(UserID int64) error
	CreateTokenUser(email string, senha string) (*user.User, string, error)

	//Read()
	//Update()
	//Delete()
}
