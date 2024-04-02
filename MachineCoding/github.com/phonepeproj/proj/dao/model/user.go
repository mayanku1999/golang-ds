package model

import (
	time2 "github.com/phonepeproj/pkg/time"
	"github.com/phonepeproj/proj/enums"
	"time"
)

type User struct {
	Id               string
	UserName         string
	Password         string
	Role             string
	SessionValidTill time2.NullTime
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time2.NullTime
}

func NewUserModel() *User {
	return &User{}
}

// Getter methods
func (u *User) GetId() string {
	return u.Id
}

func (u *User) GetUserName() string {
	return u.UserName
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *User) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

func (u *User) GetDeletedAt() time2.NullTime {
	return u.DeletedAt
}

// Setter methods returning a pointer to User
func (u *User) SetId(id string) *User {
	u.Id = id
	return u
}

func (u *User) SetUserName(userName string) *User {
	u.UserName = userName
	return u
}

func (u *User) SetPassword(password string) *User {
	u.Password = password
	return u
}

func (u *User) SetCreatedAt(createdAt time.Time) *User {
	u.CreatedAt = createdAt
	return u
}

func (u *User) SetUpdatedAt(updatedAt time.Time) *User {
	u.UpdatedAt = updatedAt
	return u
}

func (u *User) SetDeletedAt(deletedAt time.Time) *User {
	u.DeletedAt = time2.NullTime{Time: deletedAt, Valid: true}
	return u
}

func (u *User) GetRole() enums.UserRole {
	return enums.UserRole(enums.UserRole_value[u.Role])
}

func (u *User) SetRole(role enums.UserRole) *User {
	u.Role = enums.UserRole_name[int32(role)]
	return u
}

func (u *User) GetSessionValidTill() time2.NullTime {
	return u.SessionValidTill
}

func (u *User) SetSessionValidTill(time time2.NullTime) *User {
	u.SessionValidTill = time
	return u
}
