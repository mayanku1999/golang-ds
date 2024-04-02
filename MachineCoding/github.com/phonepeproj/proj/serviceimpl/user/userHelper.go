package user

import (
	"fmt"
	time2 "github.com/phonepeproj/pkg/time"
	"time"
)

func (u *UserSvcImpl) isUserSessionValid(sessionInfo time2.NullTime) (bool, error) {

	if !sessionInfo.Valid {
		return false, fmt.Errorf("user has not logged in ever")
	}

	if sessionInfo.Time.Before(time.Now()) {
		return false, fmt.Errorf("user session is expired, login again")
	}
	return true, nil
}

func (u *UserSvcImpl) validateUserNamePassword(userName, password string) (bool, error) {
	if userName == "" || password == "" {
		return false, fmt.Errorf("username , password can not be emppty")
	}
	return true, nil
}
