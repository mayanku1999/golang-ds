package main

import (
	"Users/mayank/Documents/learn3/github.com/proj/projname"
	"Users/mayank/Documents/learn3/github.com/proj/projname/config"
	"fmt"
)

func main() {
	conf := config.GetSingleConfigInstance()
	fmt.Println(conf)
	players := projname.NewPlayersSvc(conf)
	players.Create()
	players.Create()
	fmt.Println((players.Players()))
}
