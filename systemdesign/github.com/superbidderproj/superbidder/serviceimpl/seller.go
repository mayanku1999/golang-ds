package serviceimpl

import (
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/config"
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/dao"
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/dao/impl"
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/dao/model"
)

type SellerSvcImpl struct {
	conf      *config.AppConfig
	SellerDao dao.ISellerDao
}

//func (a *SellerSvcImpl) GetSellers() map[string]*model.Seller {
//	return a.sellers
//}
//func (s *SellerSvcImpl) GetSellerById(id string) (*model.Seller, error) {
//	seller, ok := s.sellers[id]
//	if !ok {
//		return nil, errors.New("seller record not found")
//	}
//	return seller, nil
//}

func NewSellerSvcImpl(conf *config.AppConfig) *SellerSvcImpl {
	return &SellerSvcImpl{
		SellerDao: impl.NewSellerDaoImpl(),
		conf:      conf,
	}
}

func (a *SellerSvcImpl) Create(id, name string) *model.Seller {
	return a.SellerDao.CreateSeller(id, name)
}

func (a *SellerSvcImpl) UpdateAuctionForSeller(sellerId, auctionId string) (*model.Seller, error) {
	seller, err := a.SellerDao.UpdateAuctionForGivenSellerId(sellerId, auctionId)
	if err != nil {
		return nil, err
	}
	return seller, nil
}

func (a *SellerSvcImpl) GetAll() map[string]*model.Seller {
	return a.SellerDao.GetAll()
}
