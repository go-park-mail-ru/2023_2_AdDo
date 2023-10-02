package album

type Album struct {
	Id      uint64
	Name    string
	Preview string
	artist  uint64
	track   []uint64
}

type Response struct {
	Id      uint64
	Name    string
	Preview string
}

type Repository interface {
	GetByTrackId(trackId uint64) ([]Response, error)
	GetByArtistId(artistId uint64) ([]Response, error)
}
