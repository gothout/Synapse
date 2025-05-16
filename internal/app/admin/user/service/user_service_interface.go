package user

import user "Synapse/internal/app/admin/user/model"

type Service interface {
	Create(user *user.User) (*user.User, error)
	ReadAllUser(enterpriseId, page int64) (*[]user.User, error)
	ReadByEmail(email string) (*user.User, error)
	//Read()
	//Update()
	//Delete()
}
