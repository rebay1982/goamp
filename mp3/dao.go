package mp3

import (
	"strings"

	id3 "github.com/mikkyang/id3-go"
)

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
		trackValue := cleanUpTagValue(mp3File.Frame("TRCK").String())
		track.TrackNumber = strings.Split(trackValue, "/")[0]
	}

	track.Filename = filename

	return track
}

func cleanUpTagValue(originalValue string) string {

	cleanValue := strings.TrimSuffix(originalValue, "\u0000")
	cleanValue = strings.TrimSpace(cleanValue)

	return cleanValue
}
