package artist

type Artist struct {
	Id     uint64
	Name   string
	Avatar string
	Album  []uint64
	Track  []uint64
}

type Response struct {
	Id     uint64 `json:"Id" example:"1"`
	Name   string `json:"Name" example:"ArtistName"`
	Avatar string `json:"Avatar" example:"ArtistAvatar"`
}

type Repository interface {
	GetByTrackId(trackId uint64) ([]Response, error)
	GetByAlbumId(albumId uint64) (Response, error)
}
