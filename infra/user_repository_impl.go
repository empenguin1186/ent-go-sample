package infra

import (
	"context"
	"ent-go-sample/domain/model"
	"ent-go-sample/ent"
	"ent-go-sample/ent/user"
	"github.com/prometheus/common/log"
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
	user, err := u.client.User.
		Query().
		Select(user.FieldID, user.FieldAge, user.FieldName).
		Where(user.IDEQ(id)).
		First(u.context)

	if err != nil {
		log.Errorf("failed to findbyId on user table, err = %w", err)
		return nil, err
	}

	return &model.User{ID: user.ID, Name: user.Name, Age: user.Age}, nil
}

func (u *UserRepositoryImpl) Delete(id int) error {
	return u.client.User.DeleteOneID(id).Exec(u.context)
}

func (u *UserRepositoryImpl) Update(user *model.User) error {
	input := ent.User{ID: user.ID, Age: user.Age, Name: user.Name}
	return u.client.User.UpdateOne(&input).Exec(u.context)
}
