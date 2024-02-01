package model

import (
	"github.com/google/uuid"
	"github.com/spotify/spotify/enums"
)

type Music struct {
	Id      uuid.UUID
	Title   string
	Artists []string
	Genre   enums.Genre
}

func NewMusic() *Music {
	return &Music{}
}

func (m *Music) GetId() uuid.UUID {
	return m.Id
}

func (m *Music) GetTitle() string {
	return m.Title
}

func (m *Music) GetArtists() []string {
	return m.Artists
}

func (m *Music) GetGenre() enums.Genre {
	return m.Genre
}

func (m *Music) SetId(id uuid.UUID) *Music {
	m.Id = id
	return m
}

func (m *Music) SetTitle(id string) *Music {
	m.Title = id
	return m
}

func (m *Music) SetArtists(id []string) *Music {
	m.Artists = id
	return m
}

func (m *Music) SetGenre(id enums.Genre) *Music {
	m.Genre = id
	return m
}
