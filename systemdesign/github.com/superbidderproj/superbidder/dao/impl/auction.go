package impl

import (
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/dao/model"
	"errors"
	"time"
)

type AuctionDaoImpl struct {
	auction map[string]*model.Auction
}

func NewAuctionDaoImpl() *AuctionDaoImpl {
	return &AuctionDaoImpl{
		auction: make(map[string]*model.Auction),
	}
}

func (a *AuctionDaoImpl) Create(id, name, seller string, minBid, maxBid int) *model.Auction {
	newAuction := model.NewAuction().SetId(id).SetName(name).SetSeller(seller).SetMinBid(minBid).
		SetMaxBid(maxBid).SetIsActive(true).SetCreatedAt(time.Now())
	a.auction[id] = newAuction
	return newAuction
}

func (a *AuctionDaoImpl) GetAll() map[string]*model.Auction {
	return a.auction
}

func (a *AuctionDaoImpl) GetById(id string) (*model.Auction, error) {
	auc, ok := a.auction[id]
	if !ok {
		return nil, errors.New("record not found")
	}
	return auc, nil
}

func (a *AuctionDaoImpl) UpdateBuyersBidMap(buyerId, auctionId string, bidAmount int) error {
	_, ok := a.auction[auctionId]
	if !ok {
		return errors.New("record not found")
	}
	if a.auction[auctionId].GetBuyersBidMap() == nil {
		a.auction[auctionId].BuyersBidMap = map[string]int{
			buyerId: bidAmount,
		}
	} else {
		a.auction[auctionId].BuyersBidMap[buyerId] = bidAmount
	}
	return nil
}

func (a *AuctionDaoImpl) DeleteBuyerFromBuyersBidMap(buyerId, auctionId string) error {
	_, ok := a.auction[auctionId]
	if !ok || a.auction[auctionId].GetBuyersBidMap() == nil {
		return errors.New("record not found")
	}
	delete(a.auction[auctionId].BuyersBidMap, buyerId)
	return nil
}
