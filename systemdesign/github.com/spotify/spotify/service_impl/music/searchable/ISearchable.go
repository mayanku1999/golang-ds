package searchable

import (
	"github.com/spotify/spotify/dao/model"
	"github.com/spotify/spotify/service_impl/music"
	"github.com/spotify/spotify/service_impl/music/searchable/artist_searcher"
	"github.com/spotify/spotify/service_impl/music/searchable/genre_searcher"
)

type ISearchable interface {
	Search(query string) []*model.Music
}

type FactorySearchable struct {
	genreSearcher  *genre_searcher.GenreSearcher
	artistSearcher *artist_searcher.ArtistSearcher
}

func NewFactorySearchable(musicClient *music.MusicSvcImpl) *FactorySearchable {
	return &FactorySearchable{
		genreSearcher: &genre_searcher.GenreSearcher{
			MusicClient: musicClient,
		},
		artistSearcher: &artist_searcher.ArtistSearcher{
			MusicClient: musicClient,
		},
	}
}

func (f *FactorySearchable) GetSearchImpl(searchType string) ISearchable {
	switch searchType {
	case "GENRE":
		return f.genreSearcher
	case "ARTIST":
		return f.artistSearcher
	}
	return nil
}
