package genre_searcher

import (
	"github.com/spotify/spotify/dao/model"
	"github.com/spotify/spotify/enums"
	"github.com/spotify/spotify/service_impl/music"
)

type GenreSearcher struct {
	MusicClient *music.MusicSvcImpl
}

func (a *GenreSearcher) Search(genre string) []*model.Music {
	genreEnum := enums.Genre_value[genre]
	return a.MusicClient.GetAllByGenre(enums.Genre(genreEnum))
}
