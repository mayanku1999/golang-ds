package daoimpl

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/phonepeproj/proj/dao"
	"github.com/phonepeproj/proj/dao/model"
	"github.com/phonepeproj/proj/enums"
	"github.com/phonepeproj/proj/serviceimpl/docversion"
	maps "golang.org/x/exp/maps"
	"time"
)

type DocDaoImpl struct {
	Docs map[string]*model.Document
}

func NewDocDaoImpl() *DocDaoImpl {
	return &DocDaoImpl{
		Docs: make(map[string]*model.Document),
	}
}

var _ dao.IDocDao = &DocDaoImpl{}

func (d *DocDaoImpl) GetById(docId string) (*model.Document, error) {
	doc, ok := d.Docs[docId]
	if !ok {
		return nil, fmt.Errorf("invalid doc id/doc does not exist")
	}
	return doc, nil
}
func (d *DocDaoImpl) GetByModeAndUser(userId string, mode enums.PublishedMode) ([]*model.Document, error) {
	res := []*model.Document{}
	for _, doc := range maps.Values(d.Docs) {
		if doc.CreatedBy == userId && doc.PublishMode == mode {
			res = append(res, doc)
		}
	}
	return res, nil
}

func (d *DocDaoImpl) Create(authorId, name, content string, publishMode enums.PublishedMode) (*model.Document, error) {
	newDoc := model.NewDocument().
		SetId(uuid.NewString()).
		SetName(name).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetCreatedBy(authorId).
		SetPublishMode(publishMode).
		SetCurrentContent(content)
	newDoc.VersionsControl = &docversion.VersionsControl{Versions: []*docversion.Version{{Content: content, CreatedAt: time.Now()}}}
	if d.Docs == nil {
		d.Docs = map[string]*model.Document{
			newDoc.GetId(): newDoc,
		}
	} else {
		d.Docs[newDoc.GetId()] = newDoc
	}
	privilegeInfo := &model.PrivilegeInfo{
		AccessStatus: enums.AccessStatus_ACTIVE,
		AccessMode:   enums.AccessMode_WRITE_READ,
	}
	newDoc.Privileges = &model.Privileges{Privileges: map[string]*model.PrivilegeInfo{authorId: privilegeInfo}}
	return newDoc, nil
}

func (d *DocDaoImpl) Update(authorID, docId string, newContent string) (*model.Document, error) {
	doc, ok := d.Docs[docId]
	if !ok {
		return nil, fmt.Errorf("invalid doc id/doc does not exist")
	}

	privilegeinfo := doc.GetPrivileges(authorID)
	if privilegeinfo == nil {
		return nil, fmt.Errorf("permission denied")
	}
	if !((privilegeinfo.AccessMode == enums.AccessMode_WRITE_READ || privilegeinfo.AccessMode == enums.AccessMode_WRITE) && privilegeinfo.AccessStatus == enums.AccessStatus_ACTIVE) {
		return nil, fmt.Errorf("permission denied")
	}

	// Acquire lock
	doc.Lock.Lock()
	defer doc.Lock.Unlock()

	updatedVersions := append(doc.GetVersionControls().Versions, &docversion.Version{Content: newContent, CreatedAt: time.Now()})
	updatedDoc := doc.SetCurrentContent(newContent).SetVersionControls(updatedVersions)
	return updatedDoc, nil
}

func (d *DocDaoImpl) Delete(authorID, docId string) error {
	doc, ok := d.Docs[docId]
	if !ok {
		return fmt.Errorf("invalid doc id/doc does not exist")
	}

	privilegeinfo := doc.GetPrivileges(authorID)
	if !((privilegeinfo.AccessMode == enums.AccessMode_WRITE_READ || privilegeinfo.AccessMode == enums.AccessMode_WRITE) && privilegeinfo.AccessStatus == enums.AccessStatus_ACTIVE) {
		return fmt.Errorf("permission denied")
	}

	// Acquire lock
	doc.Lock.Lock()
	defer doc.Lock.Unlock()

	delete(d.Docs, docId)
	return nil
}

func (d *DocDaoImpl) GetLatestVersion(docId string) (*docversion.Version, error) {
	doc, ok := d.Docs[docId]
	if !ok {
		return nil, fmt.Errorf("invalid doc id/doc does not exist")
	}
	versions := doc.VersionsControl.Versions
	latestVersion := versions[len(versions)-1]
	return latestVersion, nil
}
