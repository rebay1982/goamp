package mp3

import (
	"os"
	"sort"
	"strings"

	"path/filepath"
)

var musicLibraryIndex []artistInfo

var sortedArtists []string
var tracks []TrackInfo

// RefreshLibrary refresh the artist library.
func RefreshLibrary(sourceDir string) {
	filepath.Walk(sourceDir, indexMp3)

	// Generate sorted artist list.
	for i := range musicLibraryIndex {
		sortedArtists = append(sortedArtists, musicLibraryIndex[i].Artist)
		sort.Strings(sortedArtists)
	}
}

// GetTracks returns tracks found in library
func GetTracks() []TrackInfo {
	return tracks
}

// GetArtists retrieve the list of indexed artists, sorted alphabetically
func GetArtists() []string {
	return sortedArtists
}

func indexMp3(path string, info os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if strings.HasSuffix(path, ".mp3") {
		trackInfo := readTrackInfo(path)

		//getArtistStructure(trackInfo.Artist)

		tracks = append(tracks, trackInfo)
	}

	return nil
}

// func getArtistStructure(artist string) artistInfo {

// 	var create = true
// 	var artInfo artistInfo

// 	for i := range musicLibraryIndex {
// 		if musicLibraryIndex[i].Artist == artist {
// 			artInfo = musicLibraryIndex[i]
// 			create = false
// 			break
// 		}
// 	}

// 	if create {
// 		artInfo = artistInfo{artist, make([]albumInfo, 1)}
// 		musicLibraryIndex = append(musicLibraryIndex, artInfo)
// 	}

// 	return artInfo
// }
