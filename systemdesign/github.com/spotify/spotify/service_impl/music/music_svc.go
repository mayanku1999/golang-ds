package music

import (
	"github.com/spotify/spotify/config"
	"github.com/spotify/spotify/dao"
	"github.com/spotify/spotify/dao/impl"
	"github.com/spotify/spotify/dao/model"
	"github.com/spotify/spotify/enums"
)

type MusicSvcImpl struct {
	conf     *config.AppConfig
	MusicDao dao.IMusic
}

func NewMusicSvcImpl(appConfig *config.AppConfig) *MusicSvcImpl {
	return &MusicSvcImpl{
		conf:     appConfig,
		MusicDao: impl.NewMusicRepo(),
	}
}

func (m *MusicSvcImpl) Create(title string, genre enums.Genre, artists []string) {
	m.MusicDao.Create(title, genre, artists)
}

func (m *MusicSvcImpl) GetAll() []*model.Music {
	return m.MusicDao.GetAll()
}

func (m *MusicSvcImpl) GetAllByGenre(genre enums.Genre) []*model.Music {
	return m.MusicDao.GetAllByGenre(genre)
}

func (m *MusicSvcImpl) GetAllByArtist(artist string) []*model.Music {
	return m.MusicDao.GetAllByArtist(artist)
}
