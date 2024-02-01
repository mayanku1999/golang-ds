package highest_uniq_bid

import (
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/dao/model"
	"errors"
	"math"
)

type HighestUniqBid struct {
	Auction *model.Auction
}

func (l *HighestUniqBid) GetWinner(auctionId string) (string, error) {
	buyersBidMap := l.Auction.GetBuyersBidMap()
	if len(buyersBidMap) == 0 {
		return "", errors.New("winner buyer can not be found")
	}
	bidSet := map[int]struct{}{}
	var toBeRemoved []int
	maxVal := math.MinInt
	for _, v := range buyersBidMap {
		if _, ok := bidSet[v]; ok {
			toBeRemoved = append(toBeRemoved, v)
		} else {
			bidSet[v] = struct{}{}
		}
	}
	for k, v := range buyersBidMap {
		for _, bid := range toBeRemoved {
			if v == bid {
				delete(buyersBidMap, k)
			}
		}
	}
	for _, v := range buyersBidMap {
		if v > maxVal {
			maxVal = v
		}
	}
	for k, v := range buyersBidMap {
		if v == maxVal {
			return k, nil
		}
	}

	return "NO WINNERS", nil
}
