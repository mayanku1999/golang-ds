package model

type Player struct {
	id         string
	playerType string
}

func (p *Player) PlayerType() string {
	return p.playerType
}

func (p *Player) SetPlayerType(playerType string) {
	p.playerType = playerType
}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) Id() string {
	return p.id
}

func (p *Player) SetId(id string) {
	p.id = id
}
