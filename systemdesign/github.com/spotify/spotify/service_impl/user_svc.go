package service_impl

import (
	"github.com/spotify/spotify/config"
	"github.com/spotify/spotify/dao"
	"github.com/spotify/spotify/dao/impl"
	"github.com/spotify/spotify/dao/model"
	"github.com/spotify/spotify/service_impl/music/searchable"
)

type UserSvcImpl struct {
	conf             *config.AppConfig
	UserDao          dao.IUser
	SearchableClient *searchable.FactorySearchable
}

func NewUserSvcImpl(conf *config.AppConfig, factorySearchable *searchable.FactorySearchable) *UserSvcImpl {
	return &UserSvcImpl{
		conf:             conf,
		UserDao:          impl.NewUserRepository(),
		SearchableClient: factorySearchable,
	}
}

func (u *UserSvcImpl) Create(name string) {
	u.UserDao.Create(name)
}

func (u *UserSvcImpl) GetAll() []*model.User {
	return u.UserDao.GetAll()
}

func (u *UserSvcImpl) SearchSong(searchType string, searchVal string) []*model.Music {
	searchableImpl := u.SearchableClient.GetSearchImpl(searchType)
	return searchableImpl.Search(searchVal)
}
