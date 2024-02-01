package impl

import (
	"github.com/google/uuid"
	"github.com/spotify/spotify/dao/model"
	"golang.org/x/exp/maps"
)

type UserRepository struct {
	Users map[uuid.UUID]*model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		Users: make(map[uuid.UUID]*model.User),
	}
}

func (u *UserRepository) Create(name string) *model.User {
	newUser := model.NewUser().SetId(uuid.New()).SetName(name)
	u.Users[newUser.GetId()] = newUser
	return newUser
}

func (u *UserRepository) GetAll() []*model.User {
	return maps.Values(u.Users)
}
