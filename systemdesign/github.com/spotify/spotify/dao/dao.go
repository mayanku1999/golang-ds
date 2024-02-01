package dao

import (
	"github.com/spotify/spotify/dao/model"
	"github.com/spotify/spotify/enums"
)

type IUser interface {
	Create(name string) *model.User
	GetAll() []*model.User
}
type IMusic interface {
	Create(title string, genre enums.Genre, artists []string) *model.Music
	GetAll() []*model.Music
	GetAllByGenre(genreName enums.Genre) []*model.Music
	GetAllByArtist(artistName string) []*model.Music
}
type IPlaylist interface {
	CreatePlaylist()
}
