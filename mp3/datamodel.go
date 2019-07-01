package mp3

type artistInfo struct {
	Artist    string
	AlbumTree []albumInfo
}

type albumInfo struct {
	Album  string
	Tracks []TrackInfo
}

// TrackInfo Struct containing the track information.
type TrackInfo struct {
	Artist      string
	Album       string
	Title       string
	TrackNumber string
	Filename    string
}
