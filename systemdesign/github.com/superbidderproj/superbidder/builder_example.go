package main

import (
	"fmt"
	"strings"
)

// Song represents a song entity.
type Song struct {
	Title  string
	Artist string
	Genre  string
}

// SongDatabase represents a database of songs.
type SongDatabase struct {
	songs []Song
}

// SongSearchStrategy defines the strategy interface for searching songs.
type SongSearchStrategy interface {
	Search(songs []Song, query string) []Song
}

// TitleSearchStrategy implements the SongSearchStrategy for searching by title.
type TitleSearchStrategy struct{}

func (t *TitleSearchStrategy) Search(songs []Song, query string) []Song {
	var results []Song
	for _, song := range songs {
		if strings.Contains(strings.ToLower(song.Title), strings.ToLower(query)) {
			results = append(results, song)
		}
	}
	return results
}

// ArtistSearchStrategy implements the SongSearchStrategy for searching by artist.
type ArtistSearchStrategy struct{}

func (a *ArtistSearchStrategy) Search(songs []Song, query string) []Song {
	var results []Song
	for _, song := range songs {
		if strings.Contains(strings.ToLower(song.Artist), strings.ToLower(query)) {
			results = append(results, song)
		}
	}
	return results
}

// GenreSearchStrategy implements the SongSearchStrategy for searching by genre.
type GenreSearchStrategy struct{}

func (g *GenreSearchStrategy) Search(songs []Song, query string) []Song {
	var results []Song
	for _, song := range songs {
		if strings.Contains(strings.ToLower(song.Genre), strings.ToLower(query)) {
			results = append(results, song)
		}
	}
	return results
}

// SongSearchBuilder builds complex search queries.
type SongSearchBuilder struct {
	strategies []SongSearchStrategy
}

func NewSongSearchBuilder() *SongSearchBuilder {
	return &SongSearchBuilder{}
}

func (b *SongSearchBuilder) AddStrategy(strategy SongSearchStrategy) *SongSearchBuilder {
	b.strategies = append(b.strategies, strategy)
	return b
}

func (b *SongSearchBuilder) Search(songs []Song, query string) []Song {
	var results []Song
	for _, strategy := range b.strategies {
		results = append(results, strategy.Search(songs, query)...)
	}
	return results
}

func main() {
	// Sample data
	songs := []Song{
		{"Song1", "Artist1", "Genre1"},
		{"Song2", "Artist2", "Genre2"},
		{"Song3", "Artist3", "Genre3"},
	}

	// Using the Builder pattern to create a search with multiple strategies
	builder := NewSongSearchBuilder().
		AddStrategy(&TitleSearchStrategy{}).
		AddStrategy(&ArtistSearchStrategy{}).
		AddStrategy(&GenreSearchStrategy{})

	// Performing a search with multiple criteria
	results := builder.Search(songs, "Song")
	fmt.Println("Search Results:", results)
}
