package impl

import (
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/dao/model"
	"errors"
)

type BuyerDaoImpl struct {
	buyers map[string]*model.Buyer
}

func NewBuyerDaoImpl() *BuyerDaoImpl {
	return &BuyerDaoImpl{
		buyers: make(map[string]*model.Buyer),
	}
}

func (b *BuyerDaoImpl) CreateBuyer(id, name string) *model.Buyer {
	buyer := model.NewBuyer().SetId(id).SetName(name)
	b.buyers[buyer.GetId()] = buyer
	return buyer
}

func (b *BuyerDaoImpl) GetAll() map[string]*model.Buyer {
	return b.buyers
}

func (b *BuyerDaoImpl) AddParticipatedAuction(buyerId string, auctionId string) error {
	buyer, ok := b.buyers[buyerId]
	if !ok {
		return errors.New("record not found")
	}
	buyer.AuctionParticipatedIn = append(buyer.AuctionParticipatedIn, auctionId)
	return nil
}

func (b *BuyerDaoImpl) WithdrawParticipatedAuction(buyerId string, auctionId string) error {
	buyer, ok := b.buyers[buyerId]
	if !ok || len(buyer.GetAuctionParticipatedIn()) == 0 {
		return errors.New("record not found or auction is not present for buyer")
	}
	buyer.AuctionParticipatedIn = removeValueInPlace(buyer.GetAuctionParticipatedIn(), auctionId)
	return nil
}

func removeValueInPlace(slice []string, value string) []string {
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			// Remove the element at index i
			slice = append(slice[:i], slice[i+1:]...)
			i-- // Decrement i to account for the removed element
		}
	}
	if len(slice) == 0 {
		return nil
	}
	return slice
}
