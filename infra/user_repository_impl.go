package infra

import (
	"context"
	"ent-go-sample/domain/model"
	"ent-go-sample/ent"
)

type UserRepositoryImpl struct {
	client  *ent.Client
	context context.Context
}

func NewUserRepositoryImpl(client *ent.Client, context context.Context) *UserRepositoryImpl {
	return &UserRepositoryImpl{client: client, context: context}
}

func (u *UserRepositoryImpl) Create(user *model.User) error {
	_, err := u.client.User.
		Create().
		SetAge(user.Age).
		SetName(user.Name).
		Save(u.context)
	return err
}

func (u *UserRepositoryImpl) FindById(id int) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepositoryImpl) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepositoryImpl) Update(user *model.User) error {
	//TODO implement me
	panic("implement me")
}
