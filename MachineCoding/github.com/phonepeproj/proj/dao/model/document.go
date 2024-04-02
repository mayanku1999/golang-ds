package model

import (
	"github.com/phonepeproj/proj/enums"
	"sync"
	"time"
)

type Document struct {
	Id          string
	Name        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	CreatedBy   string
	PublishMode enums.PublishedMode
	*Privileges
	CurrentContent string
	*VersionsControl
	Lock *sync.Mutex
}

func NewDocument() *Document {
	return &Document{Lock: &sync.Mutex{}}
}

// Getter methods
func (d *Document) GetId() string {
	return d.Id
}

func (d *Document) GetName() string {
	return d.Name
}

func (d *Document) GetCreatedAt() time.Time {
	return d.CreatedAt
}

func (d *Document) GetUpdatedAt() time.Time {
	return d.UpdatedAt
}

func (d *Document) GetDeletedAt() time.Time {
	return d.DeletedAt
}

func (d *Document) GetCreatedBy() string {
	return d.CreatedBy
}

func (d *Document) GetPublishMode() enums.PublishedMode {
	return d.PublishMode
}

func (d *Document) GetPrivileges(userId string) *PrivilegeInfo {
	return d.Privileges.Privileges[userId]
}

func (d *Document) GetCurrentContent() string {
	return d.CurrentContent
}

func (d *Document) GetVersionControls() *VersionsControl {
	return d.VersionsControl
}

// Setter methods returning a pointer to Document
func (d *Document) SetId(id string) *Document {
	d.Id = id
	return d
}

func (d *Document) SetName(name string) *Document {
	d.Name = name
	return d
}

func (d *Document) SetCreatedAt(createdAt time.Time) *Document {
	d.CreatedAt = createdAt
	return d
}

func (d *Document) SetUpdatedAt(updatedAt time.Time) *Document {
	d.UpdatedAt = updatedAt
	return d
}

func (d *Document) SetDeletedAt(deletedAt time.Time) *Document {
	d.DeletedAt = deletedAt
	return d
}

func (d *Document) SetCreatedBy(createdBy string) *Document {
	d.CreatedBy = createdBy
	return d
}

func (d *Document) SetPublishMode(publishMode enums.PublishedMode) *Document {
	d.PublishMode = publishMode
	return d
}

func (d *Document) SetPrivileges(userId string, accessMode enums.AccessMode, status enums.AccessStatus) *Document {
	d.Privileges.Privileges[userId] = &PrivilegeInfo{
		AccessStatus: status,
		AccessMode:   accessMode,
	}
	return d
}

func (d *Document) SetCurrentContent(currentContent string) *Document {
	d.CurrentContent = currentContent
	return d
}

func (d *Document) SetVersionControls(versionControls []*Version) *Document {
	d.Versions = versionControls
	return d
}
