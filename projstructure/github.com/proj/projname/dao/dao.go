package dao

import "Users/mayank/Documents/learn3/github.com/proj/projname/dao/model"

type IPlayerDao interface {
	CreatePlayer() *model.Player
}
