package daoimpl

import (
	"fmt"
	"github.com/google/uuid"
	time2 "github.com/phonepeproj/pkg/time"
	"github.com/phonepeproj/proj/dao/model"
	"github.com/phonepeproj/proj/enums"
	"time"
)

type UserDaoImpl struct {
	Users map[string]*model.User
}

func NewUserDaoImpl() *UserDaoImpl {
	return &UserDaoImpl{
		Users: map[string]*model.User{},
	}
}

func (u *UserDaoImpl) Create(userName, password string, role enums.UserRole) (*model.User, error) {
	newUser := model.NewUserModel().
		SetId(uuid.NewString()).
		SetUserName(userName).
		SetPassword(password).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetRole(role)
	u.Users[newUser.GetId()] = newUser
	return newUser, nil
}

func (u *UserDaoImpl) Login(userId, userName, password string) (*model.User, error) {
	us, ok := u.Users[userId]
	if !ok {
		return nil, fmt.Errorf("invalid user login")
	}
	if us.UserName == userName && us.Password == password {
		us.SessionValidTill = time2.NullTime{Time: time.Now().Add(1 * time.Hour), Valid: true}
		return us, nil
	}
	return nil, fmt.Errorf("invalid username/password")
}

func (u *UserDaoImpl) Get(userId string) (*model.User, error) {
	us, ok := u.Users[userId]
	if !ok {
		return nil, fmt.Errorf("invalid user login")
	}
	return us, nil
}
