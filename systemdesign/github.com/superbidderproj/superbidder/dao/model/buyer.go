package model

type Buyer struct {
	Id                    string
	Name                  string
	AuctionParticipatedIn []string
}

func NewBuyer() *Buyer {
	return &Buyer{}
}

func (b *Buyer) GetId() string {
	return b.Id
}

func (b *Buyer) GetName() string {
	return b.Name
}

func (b *Buyer) GetAuctionParticipatedIn() []string {
	return b.AuctionParticipatedIn
}

func (b *Buyer) SetId(v string) *Buyer {
	b.Id = v
	return b
}

func (b *Buyer) SetName(v string) *Buyer {
	b.Name = v
	return b
}

func (b *Buyer) SetAuctionParticipatedIn(v []string) *Buyer {
	b.AuctionParticipatedIn = v
	return b
}
