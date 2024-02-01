package winningstrategy

type IWinningStrategy interface {
	GetWinner(auctionId string) (string, error)
}

type AuctionWinner struct {
	AucWinner IWinningStrategy
}

func (w *AuctionWinner) DeclareWinner(auctionId string) (string, error) {
	return w.AucWinner.GetWinner(auctionId)
}
