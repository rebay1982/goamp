package mp3

import (
	"os"
	"strings"

	// https://github.com/mikkyang/id3-go
	id3 "github.com/mikkyang/id3-go"

	"path/filepath"
)

// TrackInfo Struct containing the track information.
type TrackInfo struct {
	Artist      string
	Album       string
	Title       string
	TrackNumber string
	Filename    string
}

var tracks []TrackInfo

func cleanUpTagValue(originalValue string) string {

	cleanValue := strings.TrimSuffix(originalValue, "\u0000")
	cleanValue = strings.TrimSpace(cleanValue)

	return cleanValue
}

func readTrackInfo(filename string) TrackInfo {
	mp3File, _ := id3.Open(filename)
	defer mp3File.Close()

	tagVersion := mp3File.Version()
	var track TrackInfo
	track.Artist = cleanUpTagValue(mp3File.Artist())
	track.Album = cleanUpTagValue(mp3File.Album())
	track.Title = cleanUpTagValue(mp3File.Title())

	// Skip id3 v1 for now.
	track.TrackNumber = "NA"
	if tagVersion >= "2.2.0" {
		track.TrackNumber = cleanUpTagValue(mp3File.Frame("TRCK").String())
	}

	track.Filename = filename

	return track
}

func indexMp3(path string, info os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if strings.HasSuffix(path, ".mp3") {
		tracks = append(tracks, readTrackInfo(path))
	}

	return nil
}

// RefreshLibrary refresh the artist library.
func RefreshLibrary(sourceDir string) {
	filepath.Walk(sourceDir, indexMp3)
}

// GetTracks returns tracks found in library
func GetTracks() []TrackInfo {
	return tracks
}
