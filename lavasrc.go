package lavasrc

type PlaylistType string

const (
	PlaylistTypeAlbum           PlaylistType = "album"
	PlaylistTypePlaylist        PlaylistType = "playlist"
	PlaylistTypeArtist          PlaylistType = "artist"
	PlaylistTypeRecommendations PlaylistType = "recommendations"
)

type PlaylistInfo struct {
	Type       PlaylistType `json:"type"`
	URL        string       `json:"url"`
	ArtworkURL string       `json:"artworkUrl"`
	Author     string       `json:"author"`
}

type TrackInfo struct {
	AlbumName        string `json:"albumName"`
	AlbumURL         string `json:"albumUrl"`
	ArtistURL        string `json:"artistUrl"`
	ArtistArtworkURL string `json:"artistArtworkUrl"`
	PreviewURL       string `json:"previewUrl"`
	IsPreview        bool   `json:"isPreview"`
}
