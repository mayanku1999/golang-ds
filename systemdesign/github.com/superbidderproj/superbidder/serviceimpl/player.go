package serviceimpl

import (
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/config"
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/dao"
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/dao/impl"
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/dao/model"
)

type PlayersSvc struct {
	conf      *config.AppConfig
	PlayerDao dao.IPlayerDao
	players   []*model.Player
}

func (p *PlayersSvc) Players() []*model.Player {
	return p.players
}

func (p *PlayersSvc) SetPlayers(players []*model.Player) {
	p.players = players
}

func NewPlayersSvc(conf *config.AppConfig) *PlayersSvc {
	return &PlayersSvc{
		PlayerDao: impl.NewPlayerDaoImpl(),
		players:   []*model.Player{},
		conf:      conf,
	}
}

func (p *PlayersSvc) Create() *model.Player {
	newPlayer := p.PlayerDao.CreatePlayer()
	if p.players == nil {
		p.players = []*model.Player{}
	}
	p.players = append(p.players, newPlayer)
	return p.PlayerDao.CreatePlayer()
}
