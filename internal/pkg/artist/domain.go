package artist

type Artist struct {
	Id     uint64
	Name   string
	Avatar string
	Album  []uint64
	Track  []uint64
}

type Response struct {
	Id     uint64
	Name   string
	Avatar string
}

type Repository interface {
	GetByTrackId(trackId uint64) ([]Response, error)
	GetByAlbumId(albumId uint64) (Response, error)
}
