package dao

import (
	"github.com/phonepeproj/proj/dao/model"
	"github.com/phonepeproj/proj/enums"
	"github.com/phonepeproj/proj/serviceimpl/docversion"
)

type IUserDao interface {
	Create(userName, password string, role enums.UserRole) (*model.User, error)
	Login(userId, userName, password string) (*model.User, error)
	Get(userId string) (*model.User, error)
}

type IDocDao interface {
	GetById(docId string) (*model.Document, error)
	GetByModeAndUser(userId string, mode enums.PublishedMode) ([]*model.Document, error)
	Create(authorId, name, content string, publishMode enums.PublishedMode) (*model.Document, error)
	Update(authorID, docId string, newContent string) (*model.Document, error)
	GetLatestVersion(docId string) (*docversion.Version, error)
	Delete(authorID, docId string) error
	//Update()
}
