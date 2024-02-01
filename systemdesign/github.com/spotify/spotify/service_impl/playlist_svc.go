package service_impl

import (
	"github.com/spotify/spotify/config"
	"github.com/spotify/spotify/dao"
)

type PlaylistSvcImpl struct {
	conf        *config.AppConfig
	PlaylistDao dao.IPlaylist
}
