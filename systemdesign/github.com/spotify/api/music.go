package api

import (
	"github.com/spotify/spotify/dao/model"
	"github.com/spotify/spotify/enums"
)

type MusicSvc interface {
	Create(title string, genre enums.Genre, artists []string)
	GetAll() []*model.Music
	GetAllByGenre(genre enums.Genre) []*model.Music
	GetAllByArtist(artist string) []*model.Music
}
