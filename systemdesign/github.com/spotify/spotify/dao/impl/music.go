package impl

import (
	"github.com/google/uuid"
	"github.com/spotify/spotify/dao/model"
	"github.com/spotify/spotify/enums"
	"golang.org/x/exp/maps"
)

type MusicRepository struct {
	MusicCatalogue map[uuid.UUID]*model.Music
}

func NewMusicRepo() *MusicRepository {
	return &MusicRepository{
		MusicCatalogue: make(map[uuid.UUID]*model.Music),
	}

}

func (m *MusicRepository) Create(title string, genre enums.Genre, artists []string) *model.Music {
	newMusic := model.NewMusic().SetId(uuid.New()).SetTitle(title).SetArtists(artists).SetGenre(genre)
	m.MusicCatalogue[newMusic.GetId()] = newMusic
	return newMusic
}

func (m *MusicRepository) GetAll() []*model.Music {
	return maps.Values(m.MusicCatalogue)
}
func (m *MusicRepository) GetAllByArtist(artistName string) []*model.Music {
	var res []*model.Music
	for _, v := range maps.Values(m.MusicCatalogue) {
		for _, artist := range v.GetArtists() {
			if artist == artistName {
				res = append(res, v)
			}
		}
	}
	return res
}

func (m *MusicRepository) GetAllByGenre(genreName enums.Genre) []*model.Music {
	var res []*model.Music
	for _, v := range maps.Values(m.MusicCatalogue) {
		if v.GetGenre() == genreName {
			res = append(res, v)
		}
	}
	return res
}
