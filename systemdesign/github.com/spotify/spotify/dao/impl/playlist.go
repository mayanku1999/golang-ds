package impl

import (
	"github.com/google/uuid"
	"github.com/spotify/spotify/dao/model"
)

type PlayListRepository struct {
	Playlist map[uuid.UUID]*model.Playlist
}

func NewPlayListRepo() *PlayListRepository {
	return &PlayListRepository{
		Playlist: make(map[uuid.UUID]*model.Playlist),
	}

}
