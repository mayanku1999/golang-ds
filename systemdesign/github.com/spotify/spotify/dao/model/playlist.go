package model

import (
	"github.com/google/uuid"
	"github.com/spotify/spotify/enums"
)

type Playlist struct {
	Id            uuid.UUID
	Accessibility enums.PlayListAccessibility
	MusicList     []string
	Owner         string
}

func (p *Playlist) GetId() uuid.UUID {
	return p.Id
}

func (p *Playlist) GetAccessibility() enums.PlayListAccessibility {
	return p.Accessibility
}

func (p *Playlist) GetMusicList() []string {
	return p.MusicList
}

func (p *Playlist) GetOwner() string {
	return p.Owner
}

func (p *Playlist) SetId(id uuid.UUID) *Playlist {
	p.Id = id
	return p
}

func (p *Playlist) SetAccessibility(id enums.PlayListAccessibility) *Playlist {
	p.Accessibility = id
	return p
}

func (p *Playlist) SetMusicList(id []string) *Playlist {
	p.MusicList = id
	return p
}

func (p *Playlist) SetOwner(id string) *Playlist {
	p.Owner = id
	return p
}
