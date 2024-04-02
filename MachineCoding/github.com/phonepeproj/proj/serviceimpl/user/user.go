package user

import (
	"fmt"
	//"github.com/phonepeproj/api/proj/controllers"
	"github.com/phonepeproj/proj/config"
	"github.com/phonepeproj/proj/dao"
	"github.com/phonepeproj/proj/dao/daoimpl"
	"github.com/phonepeproj/proj/dao/model"
	"github.com/phonepeproj/proj/enums"
)

type UserSvcImpl struct {
	conf    *config.AppConfig
	UserDao dao.IUserDao
}

func NewUserSvcImpl(conf *config.AppConfig) *UserSvcImpl {
	return &UserSvcImpl{
		conf:    conf,
		UserDao: daoimpl.NewUserDaoImpl(),
	}
}

//var _ controllers.IUserSvc = &UserSvcImpl{}

func (u *UserSvcImpl) Signup(userName, password string, role enums.UserRole) (*model.User, error) {
	valid, err := u.validateUserNamePassword(userName, password)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, fmt.Errorf("username , password can not be emppty")
	}

	newUser, err := u.UserDao.Create(userName, password, role)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u *UserSvcImpl) Login(userId, userName, password string) (*model.User, error) {
	valid, err := u.validateUserNamePassword(userName, password)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, fmt.Errorf("username , password can not be emppty")
	}
	loggedInUser, err := u.UserDao.Login(userId, userName, password)
	if err != nil {
		return nil, err
	}
	return loggedInUser, nil
}

func (u *UserSvcImpl) Get(userId string) (*model.User, error) {
	user, err := u.UserDao.Get(userId)
	if err != nil {
		return nil, fmt.Errorf("user does not exist")
	}
	isUserLoggedIn, err := u.isUserSessionValid(user.GetSessionValidTill())
	if err != nil {
		return nil, err
	}
	if isUserLoggedIn {
		return user, nil
	}
	return nil, fmt.Errorf("user is not logged in")
}

func (u *UserSvcImpl) CheckUserEligibilityForFeature(userId, feature string) (bool, error) {
	user, err := u.UserDao.Get(userId)
	if err != nil {
		return false, fmt.Errorf("user does not exist")
	}
	userGroupMapping, ok := u.conf.FeatureConstraint[feature]
	if !ok {
		return false, fmt.Errorf("feature does not exist")
	}
	isFeatEnabled, ok := userGroupMapping.UserGroupConstraint[enums.UserRole_name[int32(user.GetRole())]]
	if !ok {
		return false, nil
	}
	return isFeatEnabled, nil
}

func (u *UserSvcImpl) IsUserSessionValid(userId string) (bool, error) {
	user, err := u.UserDao.Get(userId)
	if err != nil {
		return false, fmt.Errorf("user does not exist")
	}
	isUserLoggedIn, err := u.isUserSessionValid(user.GetSessionValidTill())
	if err != nil {
		return false, err
	}
	return isUserLoggedIn, nil
}
