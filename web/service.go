package web

import (
	"fmt"

	"github.com/rebay1982/goamp/mp3"
)

// Service a Goamp service instance
type Service struct {
	isInitialized bool
	tracks        []mp3.TrackInfo
}

// NewService returns a new service instance
func NewService() *Service {
	return &Service{
		isInitialized: false,
	}
}

// StartService fires up the service, indexes MP3
func (s *Service) StartService() {
	mp3.RefreshLibrary("./test")

	s.tracks = mp3.GetTracks()
	s.isInitialized = true
}

// GetArtists retrive all artists
func (s *Service) GetArtists() []string {
	var artists []string

	for i := range s.tracks {
		var artist = s.tracks[i].Artist

		var skip = false
		for j := range artists {

			if artists[j] == artist {
				skip = true
				break // No need to go through the rest.
			}
		}

		if !skip {
			artists = append(artists, artist)
		}
	}

	return artists
}

// GetArtistAlbums retrive an artist's albums
func (s *Service) GetArtistAlbums(artist string) []string {
	var albums []string

	for i := range s.tracks {

		if s.tracks[i].Artist == artist {
			skip := false
			album := s.tracks[i].Album

			for j := range albums {
				if albums[j] == album {
					skip = true
					break
				}
			}

			if !skip {
				fmt.Printf("Catch the album %s", album)
				albums = append(albums, album)
			}

		} else {
			// Skip it if artist doesn't match
			continue
		}
	}

	return albums
}

// DumpTracks test method to dump track information
func (s *Service) DumpTracks() []mp3.TrackInfo {
	return s.tracks
}
