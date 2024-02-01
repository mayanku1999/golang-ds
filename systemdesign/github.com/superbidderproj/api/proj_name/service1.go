package proj_name

import (
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/dao/model"
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/enums"
)

type IPlayerSvc interface {
	Create() *model.Player
}
type ISellerSvc interface {
	Create(id, name string) *model.Seller
	UpdateAuctionForSeller(sellerId, auctionId string) (*model.Seller, error)
	GetAll() map[string]*model.Seller
}

type IAuctionSvc interface {
	GetAllAuctions() map[string]*model.Auction
	Create(id, name, seller string, minBid, maxBid int) (*model.Auction, error)
	CreateOrUpdateBid(buyerId, auctionId string, bidAmt int) error
	WithdrawBid(buyerId, auctionId string) error
}

type IBuyerSvc interface {
	Create(id, name string) *model.Buyer
	GetAll() map[string]*model.Buyer
	UpdateAuctionParticipateIn(buyerId, auctionId string, op enums.UpdateAucParticipatedOperationForBuyer) error
}
