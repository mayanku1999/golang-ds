package userController

import (
	"github.com/phonepeproj/proj/dao/model"
	"github.com/phonepeproj/proj/enums"
)

type IUserSvc interface {
	Signup(userName, password string, role enums.UserRole) (*model.User, error)
	Login(userId, userName, password string) (*model.User, error)
	CheckUserEligibilityForFeature(userId, feature string) (bool, error)
	Get(userId string) (*model.User, error)
	IsUserSessionValid(userId string) (bool, error)
}
