package dao

import "Users/mayank/Documents/learn3/github.com/proj/superbidder/dao/model"

type IPlayerDao interface {
	CreatePlayer() *model.Player
}

type ISellerDao interface {
	CreateSeller(id, name string) *model.Seller
	UpdateAuctionForGivenSellerId(sellerId, auctionId string) (*model.Seller, error)
	GetAll() map[string]*model.Seller
}

type IBuyerDao interface {
	CreateBuyer(id, name string) *model.Buyer
	GetAll() map[string]*model.Buyer
	AddParticipatedAuction(buyerId string, auctionId string) error
	WithdrawParticipatedAuction(buyerId string, auctionId string) error
}

type IAuctionDao interface {
	Create(id, name, seller string, minBid, maxBid int) *model.Auction
	GetAll() map[string]*model.Auction
	GetById(id string) (*model.Auction, error)
	UpdateBuyersBidMap(buyerId, auctionId string, bidAmount int) error
	DeleteBuyerFromBuyersBidMap(buyerId, auctionId string) error
}
