package impl

import "Users/mayank/Documents/learn3/github.com/proj/superbidder/dao/model"

type PlayerDaoImpl struct {
}

func NewPlayerDaoImpl() *PlayerDaoImpl {
	return &PlayerDaoImpl{}
}

func (p *PlayerDaoImpl) CreatePlayer() *model.Player {
	newPlayer := model.NewPlayer()
	return newPlayer
}
