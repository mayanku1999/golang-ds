package main

import (
	"Users/mayank/Documents/learn3/github.com/proj/projname"
	"Users/mayank/Documents/learn3/github.com/proj/projname/config"
	"fmt"
)

func main() {
	conf := config.LoadConfig()
	fmt.Println(conf)
	players := projname.NewPlayersSvc()
	players.Create()
	players.Create()
	fmt.Println((players.Players()))
}
