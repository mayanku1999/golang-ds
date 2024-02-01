package serviceimpl

import (
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/config"
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/dao"
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/dao/impl"
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/dao/model"
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/enums"
)

type BuyerSvcImpl struct {
	conf     *config.AppConfig
	BuyerDao dao.IBuyerDao
}

func NewBuyerSvcImpl(conf *config.AppConfig) *BuyerSvcImpl {
	return &BuyerSvcImpl{
		BuyerDao: impl.NewBuyerDaoImpl(),
		conf:     conf,
	}
}

func (b *BuyerSvcImpl) Create(id, name string) *model.Buyer {
	return b.BuyerDao.CreateBuyer(id, name)
}

func (b *BuyerSvcImpl) GetAll() map[string]*model.Buyer {
	return b.BuyerDao.GetAll()
}

func (b *BuyerSvcImpl) UpdateAuctionParticipateIn(buyerId, auctionId string, op enums.UpdateAucParticipatedOperationForBuyer) error {
	var err error
	switch op {
	case enums.CREATED_BID:
		err = b.BuyerDao.AddParticipatedAuction(buyerId, auctionId)
	case enums.WITHDRAW_BID:
		err = b.BuyerDao.WithdrawParticipatedAuction(buyerId, auctionId)
	}
	if err != nil {
		return err
	}
	return nil
}
