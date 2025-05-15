package user

import user "Synapse/internal/app/admin/user/model"

type Repository interface {
	Create(user *user.User) (*user.User, error)
}
