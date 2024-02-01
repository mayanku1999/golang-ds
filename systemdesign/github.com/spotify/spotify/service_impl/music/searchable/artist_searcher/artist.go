package artist_searcher

import (
	"github.com/spotify/spotify/dao/model"
	"github.com/spotify/spotify/service_impl/music"
)

type ArtistSearcher struct {
	MusicClient *music.MusicSvcImpl
}

func (a *ArtistSearcher) Search(artistName string) []*model.Music {
	return a.MusicClient.GetAllByArtist(artistName)
}
