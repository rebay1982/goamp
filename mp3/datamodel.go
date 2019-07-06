package mp3

type mediaIndex struct {
	SortedArtists []string
	Artists       map[string]artistInfo
}

type artistInfo struct {
	Artist       string
	SortedAlbums []string
	Albums       map[string]albumInfo
}

type albumInfo struct {
	Album        string
	SortedTracks []string
	Tracks       []TrackInfo
}

// TrackInfo Struct containing the track information.
type TrackInfo struct {
	Artist      string
	Album       string
	Title       string
	TrackNumber string
	Filename    string
}
