package serviceimpl

import (
	"Users/mayank/Documents/learn3/github.com/proj/api/proj_name"
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/config"
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/dao"
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/dao/impl"
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/dao/model"
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/enums"
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/serviceimpl/winningstrategy"
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/serviceimpl/winningstrategy/highest_uniq_bid"
	"errors"
)

type AuctionSvcImpl struct {
	conf           *config.AppConfig
	AuctionDao     dao.IAuctionDao
	SellerSvc      proj_name.ISellerSvc
	BuyerSvcClient proj_name.IBuyerSvc
}

func (a *AuctionSvcImpl) GetAllAuctions() map[string]*model.Auction {
	return a.AuctionDao.GetAll()
}

func NewAuctionSvcImpl(conf *config.AppConfig, sellerSvcClient proj_name.ISellerSvc, buyerSvcClient proj_name.IBuyerSvc) *AuctionSvcImpl {
	return &AuctionSvcImpl{
		AuctionDao:     impl.NewAuctionDaoImpl(),
		conf:           conf,
		SellerSvc:      sellerSvcClient,
		BuyerSvcClient: buyerSvcClient,
	}
}

func (a *AuctionSvcImpl) Create(id, name, seller string, minBid, maxBid int) (*model.Auction, error) {
	newAuction := a.AuctionDao.Create(id, name, seller, minBid, maxBid)
	_, err := a.SellerSvc.UpdateAuctionForSeller(seller, newAuction.GetId())
	if err != nil {
		return nil, err
	}
	return newAuction, nil
}

func (a *AuctionSvcImpl) CreateOrUpdateBid(buyerId, auctionId string, bidAmt int) error {
	auction, err := a.AuctionDao.GetById(auctionId)
	if err != nil {
		return err
	}

	if bidAmt > auction.MaxBid() || bidAmt < auction.MinBid() {
		return errors.New("invalid bid amount ")
	}

	updateBuyerBidMapErr := a.AuctionDao.UpdateBuyersBidMap(buyerId, auctionId, bidAmt)
	if updateBuyerBidMapErr != nil {
		return updateBuyerBidMapErr
	}
	updateAuctionParticipateInErr := a.BuyerSvcClient.UpdateAuctionParticipateIn(buyerId, auctionId, enums.CREATED_BID)
	if updateAuctionParticipateInErr != nil {
		return updateAuctionParticipateInErr
	}
	return nil
}

func (a *AuctionSvcImpl) Close(auctionId string) (string, int, error) {
	auction, err := a.AuctionDao.GetById(auctionId)
	if err != nil || len(auction.GetBuyersBidMap()) == 0 {
		return "", 0, err
	}
	aucWinner := winningstrategy.AuctionWinner{AucWinner: &highest_uniq_bid.HighestUniqBid{Auction: auction}}
	winnerBuyerId, winnerBuyerIdErr := aucWinner.DeclareWinner(auctionId)
	if winnerBuyerIdErr != nil {
		return "", 0, winnerBuyerIdErr
	}

	winnerBuyerBidVal, ok := auction.GetBuyersBidMap()[winnerBuyerId]
	if !ok && winnerBuyerId != "NO WINNERS" {
		return "", 0, errors.New("buyer bid map can not be found")
	}
	return winnerBuyerId, winnerBuyerBidVal, nil
}

func (a *AuctionSvcImpl) WithdrawBid(buyerId, auctionId string) error {
	_, err := a.AuctionDao.GetById(auctionId)
	if err != nil {
		return err
	}
	deleteBuyerFromBuyersBidMapErr := a.AuctionDao.DeleteBuyerFromBuyersBidMap(buyerId, auctionId)
	if deleteBuyerFromBuyersBidMapErr != nil {
		return deleteBuyerFromBuyersBidMapErr
	}
	updateAuctionParticipateInErr := a.BuyerSvcClient.UpdateAuctionParticipateIn(buyerId, auctionId, enums.WITHDRAW_BID)
	if updateAuctionParticipateInErr != nil {
		return updateAuctionParticipateInErr
	}
	return nil
}
