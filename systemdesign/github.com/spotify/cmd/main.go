package main

import (
	"fmt"
	"github.com/spotify/spotify/config"
	"github.com/spotify/spotify/enums"
	"github.com/spotify/spotify/service_impl"
	"github.com/spotify/spotify/service_impl/music"
	"github.com/spotify/spotify/service_impl/music/searchable"
)

func main() {
	conf := config.GetSingleConfigInstance()
	musicClient := music.NewMusicSvcImpl(conf)
	facSearchable := searchable.NewFactorySearchable(musicClient)
	userClient := service_impl.NewUserSvcImpl(conf, facSearchable)

	musicClient.Create("Ram Aayenge", enums.Genre_BHAKTI, []string{"Maithili Thakur"})
	musicClient.Create("Mere Kanha", enums.Genre_BHAKTI, []string{"Zubin Nautiyal"})
	musicClient.Create("Better than revenge", enums.Genre_ROCK, []string{"Taylor Swift"})
	musicClient.Create("You are losing me", enums.Genre_LOVE, []string{"Taylor Swift"})
	musicClient.Create("Gaddi ch play", enums.Genre_POP, []string{"Deep Jandu", "Bohemia"})
	musicClient.Create("Same Beef", enums.Genre_HIPHOP, []string{"Bohemia"})

	for _, v := range musicClient.GetAll() {
		fmt.Println(v.GetId(), v.GetTitle(), v.GetGenre(), v.GetArtists())
	}

	userClient.Create("Mayank")
	userClient.Create("Sainath")
	for _, v := range userClient.GetAll() {
		fmt.Println(v.GetId(), v.GetName())
	}

	fmt.Println("search by genre")
	for _, v := range userClient.SearchSong("GENRE", "BHAKTI") {
		fmt.Println(v.GetId(), v.GetTitle(), v.GetGenre(), v.GetArtists())
	}

	fmt.Println("search by artist")
	for _, v := range userClient.SearchSong("ARTIST", "Maithili Thakur") {
		fmt.Println(v.GetId(), v.GetTitle(), v.GetGenre(), v.GetArtists())
	}
}
