package impl

import (
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/dao/model"
	"errors"
)

type SellerDaoImpl struct {
	sellers map[string]*model.Seller
}

func NewSellerDaoImpl() *SellerDaoImpl {
	return &SellerDaoImpl{
		sellers: make(map[string]*model.Seller),
	}
}

func (b *SellerDaoImpl) CreateSeller(id, name string) *model.Seller {
	seller := model.NewSeller().SetId(id).SetName(name)
	b.sellers[seller.GetId()] = seller
	return seller
}

func (b *SellerDaoImpl) UpdateAuctionForGivenSellerId(sellerId, auctionId string) (*model.Seller, error) {
	_, ok := b.sellers[sellerId]
	if !ok {
		return nil, errors.New("record not found")
	}
	b.sellers[sellerId].CreatedAuctions = append(b.sellers[sellerId].CreatedAuctions, auctionId)
	return b.sellers[sellerId], nil
}

func (b *SellerDaoImpl) GetAll() map[string]*model.Seller {
	return b.sellers
}
