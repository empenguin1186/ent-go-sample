package repository

import "ent-go-sample/domain/model"

type UserRepository interface {
	Create(user *model.User) error
	FindById(id int) (*model.User, error)
	Delete(id int) error
	Update(user *model.User) error
}
