package artist

type TrackInfoResponse struct {
}

type AlbumInfoResponse struct {
}

type InfoResponse struct {
	Id     uint64
	Name   string
	Avatar string
	Track  []TrackInfoResponse
	Album  []AlbumInfoResponse
}

type Response struct {
	Id     uint64 `json:"Id" example:"1"`
	Name   string `json:"Name" example:"ArtistName"`
	Avatar string `json:"Avatar" example:"ArtistAvatar"`
}

type UseCase interface {
	GetArtistInfo(artistId uint64) (InfoResponse, error)
}

type Repository interface {
	GetByTrackId(trackId uint64) ([]Response, error)
	GetByAlbumId(albumId uint64) (Response, error)
	GetTracks(artistId uint64) ([]TrackInfoResponse, error)
	GetAlbums(artistId uint64) ([]AlbumInfoResponse, error)
	Get(artistId uint64) (Response, error)
}
