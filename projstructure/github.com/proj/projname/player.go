package projname

import (
	"Users/mayank/Documents/learn3/github.com/proj/projname/dao"
	"Users/mayank/Documents/learn3/github.com/proj/projname/dao/model"
)

type PlayersSvc struct {
	PlayerDao dao.IPlayerDao
	players   []*model.Player
}

func (p *PlayersSvc) Players() []*model.Player {
	return p.players
}

func (p *PlayersSvc) SetPlayers(players []*model.Player) {
	p.players = players
}

func NewPlayersSvc() *PlayersSvc {
	return &PlayersSvc{
		PlayerDao: dao.NewPlayerDaoImpl(),
		players:   []*model.Player{},
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
