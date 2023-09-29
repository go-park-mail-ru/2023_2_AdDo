package track

type Album struct {
	Id         uint64
	Name       string
	Release    []uint64
	FKArtistId uint64
	ImagePath  string
}
