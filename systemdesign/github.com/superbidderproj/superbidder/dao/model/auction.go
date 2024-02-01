package model

import (
	"time"
)

type Auction struct {
	Id           string
	Name         string
	Seller       string
	BuyersBidMap map[string]int
	minBid       int
	maxBid       int
	isActive     bool
	CreatedAt    time.Time
}

func (a *Auction) GetIsActive() bool {
	return a.isActive
}

func (a *Auction) SetIsActive(isActive bool) *Auction {
	a.isActive = isActive
	return a
}

func NewAuction() *Auction {
	return &Auction{}
}

func (a *Auction) GetId() string {
	return a.Id
}

func (a *Auction) SetId(Id string) *Auction {
	a.Id = Id
	return a
}

func (a *Auction) GetName() string {
	return a.Name
}

func (a *Auction) SetName(Name string) *Auction {
	a.Name = Name
	return a
}

func (a *Auction) GetSeller() string {
	return a.Seller
}

func (a *Auction) SetSeller(seller string) *Auction {
	a.Seller = seller
	return a
}

func (a *Auction) GetBuyersBidMap() map[string]int {
	return a.BuyersBidMap
}

func (a *Auction) SetBuyersBidMap(Buyers map[string]int) *Auction {
	a.BuyersBidMap = Buyers
	return a
}

func (a *Auction) MinBid() int {
	return a.minBid
}

func (a *Auction) SetMinBid(minBid int) *Auction {
	a.minBid = minBid
	return a
}

func (a *Auction) MaxBid() int {
	return a.maxBid
}

func (a *Auction) SetMaxBid(maxBid int) *Auction {
	a.maxBid = maxBid
	return a
}

func (a *Auction) GetCreatedAt() time.Time {
	return a.CreatedAt
}

func (a *Auction) SetCreatedAt(CreatedAt time.Time) *Auction {
	a.CreatedAt = CreatedAt
	return a
}
