package document

import (
	"errors"
	"fmt"
	"github.com/phonepeproj/api/proj/controllers/userController"
	"github.com/phonepeproj/proj/config"
	"github.com/phonepeproj/proj/dao"
	"github.com/phonepeproj/proj/dao/daoimpl"
	"github.com/phonepeproj/proj/dao/model"
	"github.com/phonepeproj/proj/enums"
	"github.com/phonepeproj/proj/serviceimpl/notification"
	"github.com/phonepeproj/proj/serviceimpl/user"
)

type DocSvcImpl struct {
	conf             *config.AppConfig
	DocDao           dao.IDocDao
	UserSvcClient    userController.IUserSvc
	NotificationPub1 *notification.NotificationSubjectSvcImpl
}

func NewDocSvcImpl(appConfig *config.AppConfig, impl *user.UserSvcImpl) *DocSvcImpl {
	return &DocSvcImpl{conf: appConfig, DocDao: daoimpl.NewDocDaoImpl(), UserSvcClient: impl, NotificationPub1: notification.NewNotificationSvcImpl()}
}

type DocShareSubject struct {
	UserId        string
	DocId         string
	DocAccessMode enums.AccessMode
}

func (d *DocSvcImpl) GetDocById(docId string) (*model.Document, error) {
	docs, err := d.DocDao.GetById(docId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("error getting doc by id: %v", err))
	}
	return docs, nil
}

func (d *DocSvcImpl) GetDocsModeAndUser(userId string, mode enums.PublishedMode) ([]*model.Document, error) {
	docs, err := d.DocDao.GetByModeAndUser(userId, mode)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("error getting doc by model and user: %v", err))
	}
	return docs, nil
}

func (d *DocSvcImpl) Create(authorID, title, content string, pubMode enums.PublishedMode) (*model.Document, error) {
	if title == "" || content == "" {
		return nil, errors.New("title and content cannot be empty")
	}
	_, err := d.UserSvcClient.Get(authorID)
	if err != nil {
		return nil, err
	}

	_, err = d.UserSvcClient.IsUserSessionValid(authorID)
	if err != nil {
		return nil, err
	}

	newDoc, err := d.DocDao.Create(authorID, title, content, pubMode)
	if err != nil {
		return nil, err
	}
	return newDoc, nil
}

func (d *DocSvcImpl) Update(authorID, docID string, newContent string) (*model.Document, error) {
	if newContent == "" || docID == "" || authorID == "" {
		return nil, errors.New("newContent/authorId, docId cannot be empty")
	}
	_, err := d.UserSvcClient.Get(authorID)
	if err != nil {
		return nil, err
	}
	_, err = d.UserSvcClient.IsUserSessionValid(authorID)
	if err != nil {
		return nil, err
	}

	updDoc, err := d.DocDao.Update(authorID, docID, newContent)
	if err != nil {
		return nil, err
	}
	return updDoc, nil
}

func (d *DocSvcImpl) Delete(authorId, docId string) error {
	if docId == "" || authorId == "" {
		return errors.New("docId/authorId cannot be empty")
	}
	_, err := d.UserSvcClient.Get(authorId)
	if err != nil {
		return err
	}
	_, err = d.UserSvcClient.IsUserSessionValid(authorId)
	if err != nil {
		return err
	}
	err = d.DocDao.Delete(authorId, docId)
	if err != nil {
		return err
	}
	return nil
}

func (d *DocSvcImpl) GetLatestVersion(authorId, docId string) (*model.Version, error) {
	if docId == "" || authorId == "" {
		return nil, errors.New("docId/authorId cannot be empty")
	}
	_, err := d.UserSvcClient.Get(authorId)
	if err != nil {
		return nil, err
	}
	_, err = d.UserSvcClient.IsUserSessionValid(authorId)
	if err != nil {
		return nil, err
	}
	latestVersion, err := d.DocDao.GetLatestVersion(docId)
	if err != nil {
		return nil, err
	}
	return latestVersion, nil
}

func (d *DocSvcImpl) RevertToVersion(authorId, docId string, versionNum string) error {
	if docId == "" || authorId == "" {
		return errors.New("docId/authorId cannot be empty")
	}
	_, err := d.UserSvcClient.Get(authorId)
	if err != nil {
		return err
	}
	_, err = d.UserSvcClient.IsUserSessionValid(authorId)
	if err != nil {
		return err
	}
	if versionNum == "LATEST" {
		doc, err := d.DocDao.GetById(docId)
		if err != nil {
			return err
		}
		versions := doc.GetVersionControls().Versions
		if len(versions) == 1 {
			return fmt.Errorf("no previous version exist")
		}
		if len(versions) >= 2 {
			previousVersion := versions[len(versions)-2]
			hardDeleteLatVersion := versions[:len(versions)-1]
			doc.SetVersionControls(hardDeleteLatVersion).SetCurrentContent(previousVersion.Content)
			return nil
		}

	}
	return fmt.Errorf("unimplemented")
}

func (d *DocSvcImpl) ShareDoc(docShareSubList []*DocShareSubject) error {
	if len(docShareSubList) == 0 {
		return nil
	}
	for _, doc := range docShareSubList {
		// update the doc privileges
		us, err := d.UserSvcClient.Get(doc.UserId)
		if err != nil {
			return err
		}

		d.NotificationPub1.Subscribe(&notification.DocShareConsumerImpl{User: us, UserId: us.GetId(), DocId: doc.DocId})
	}
	d.NotificationPub1.NotifyAllOnDocShare()
	return nil
}
