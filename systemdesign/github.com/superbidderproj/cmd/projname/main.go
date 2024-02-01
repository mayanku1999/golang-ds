package main

import (
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/config"
	"Users/mayank/Documents/learn3/github.com/proj/superbidder/serviceimpl"
	"fmt"
)

func main() {
	conf := config.LoadConfig()
	//fmt.Println(conf)

	//players := superbidder.NewPlayersSvc(conf)
	//players.Create()
	//players.Create()
	//fmt.Println((players.Players()))

	sellerClient := serviceimpl.NewSellerSvcImpl(conf)
	buyerClient := serviceimpl.NewBuyerSvcImpl(conf)
	auctionClient := serviceimpl.NewAuctionSvcImpl(conf, sellerClient, buyerClient)

	buyerClient.Create("b1", "b1")
	buyerClient.Create("b2", "b2")
	buyerClient.Create("b3", "b3")

	fmt.Println("all buyers", buyerClient.GetAll())

	sellerClient.Create("s1", "s1")
	sellerClient.Create("s2", "s2")

	fmt.Println("all sellers", sellerClient.GetAll())

	auc1, auc1Err := auctionClient.Create("a1", "a1", "s1", 10, 50)

	fmt.Println("all auctions before", auctionClient.GetAllAuctions())

	if auc1Err != nil {
		fmt.Printf(auc1Err.Error())
	}
	fmt.Println("auction created: ", auc1)
	bidErr := auctionClient.CreateOrUpdateBid("b1", "a1", 77)
	if bidErr != nil {
		fmt.Printf(bidErr.Error())
	}
	bidErr = auctionClient.CreateOrUpdateBid("b2", "a1", 15)
	if bidErr != nil {
		fmt.Printf(bidErr.Error())
	}
	bidErr = auctionClient.CreateOrUpdateBid("b2", "a1", 19)
	if bidErr != nil {
		fmt.Printf(bidErr.Error())
	}
	bidErr = auctionClient.CreateOrUpdateBid("b3", "a1", 19)
	if bidErr != nil {
		fmt.Printf(bidErr.Error())
	}
	bidErr = auctionClient.CreateOrUpdateBid("b4", "a1", 17)
	if bidErr != nil {
		fmt.Printf(bidErr.Error())
	}
	fmt.Println("all auctions 2 ", auctionClient.GetAllAuctions())
	fmt.Println("all buyers", buyerClient.GetAll())
	bidErr = auctionClient.WithdrawBid("b4", "a1")
	if bidErr != nil {
		fmt.Printf(bidErr.Error())
	}
	fmt.Println("all auctions 3 ", auctionClient.GetAllAuctions())

	winningBuyer, winningBid, closeAucErr := auctionClient.Close(auc1.GetId())
	if closeAucErr != nil {
		fmt.Printf(closeAucErr.Error())
	}

	fmt.Println(winningBuyer, winningBid)
}
