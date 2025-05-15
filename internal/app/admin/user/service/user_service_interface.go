package user

import user "Synapse/internal/app/admin/user/model"

type Service interface {
	Create(user *user.User) (*user.User, error)
	//Read()
	//Update()
	//Delete()
}
