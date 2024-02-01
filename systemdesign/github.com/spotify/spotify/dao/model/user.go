package model

import "github.com/google/uuid"

type User struct {
	Id            uuid.UUID
	Name          string
	Dob           string
	Playlists     []string
	SongPublished []string
}

func NewUser() *User {
	return &User{}
}
func (u *User) GetId() uuid.UUID {
	return u.Id
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetDob() string {
	return u.Dob
}

func (u *User) GetPlaylists() []string {
	return u.Playlists
}

func (u *User) GetSongPublished() []string {
	return u.SongPublished
}

func (u *User) SetId(id uuid.UUID) *User {
	u.Id = id
	return u
}

func (u *User) SetName(id string) *User {
	u.Name = id
	return u
}

func (u *User) SetDob(id string) *User {
	u.Dob = id
	return u
}

func (u *User) SetPlaylists(id []string) *User {
	u.Playlists = id
	return u
}

func (u *User) SetSongPublished(id []string) *User {
	u.SongPublished = id
	return u
}
