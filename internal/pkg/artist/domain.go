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
	Create(track Artist) (uint64, error)
	GetById(id uint64) (Artist, error)
	GetByTrackId(trackId uint64) ([]Response, error)
	GetByAlbumId(albumId uint64) (Response, error)
}
