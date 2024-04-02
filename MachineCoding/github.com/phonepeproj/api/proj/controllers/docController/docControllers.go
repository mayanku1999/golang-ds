package docController

import (
	"github.com/phonepeproj/proj/dao/model"
	"github.com/phonepeproj/proj/enums"
	"github.com/phonepeproj/proj/serviceimpl/document"
)

type IDocSvc interface {
	GetDocById(docId string) (*model.Document, error)
	GetDocsModeAndUser(userId string, mode enums.PublishedMode) ([]*model.Document, error)
	Create(authorID, title, content string, pubMode enums.PublishedMode) (*model.Document, error)
	Update(authorID, docID string, newContent string) (*model.Document, error)
	Delete(authorId, docId string) error
	GetLatestVersion(authorId, docId string) (*model.Version, error)
	RevertToVersion(authorId, docId string, versionNum string) error
	ShareDoc(docShareSubList []*document.DocShareSubject) error
}
