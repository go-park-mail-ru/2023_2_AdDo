package activity

type UserLikeTrack struct {
	UserId  string
	TrackId uint64
}

type UserLikeGenre struct {
	UserId  string
	GenreId uint64
}

type UserLikeAlbum struct {
	UserId  string
	AlbumId uint64
}

type UserLikeArtist struct {
	UserId   string
	ArtistId uint64
}

type UserListenTrack struct {
	UserId    string
	DurationS uint32
	TrackId   uint64
}

type UserSkipTrack struct {
	UserId    string
	DurationS uint32
	TrackId   uint64
}

type ConsumerRepository interface {
	PullLikeTrack() (UserLikeTrack, error)
	PullLikeAlbum() (UserLikeAlbum, error)
	PullLikeArtist() (UserLikeArtist, error)
	PullLikeGenre() (UserLikeGenre, error)
	PullSkipTrack() (UserSkipTrack, error)
	PullListenTrack() (UserListenTrack, error)
}

type ProducerRepository interface {
	PushLikeTrack(userId string, trackId uint64) error
	PushLikeAlbum(userId string, albumId uint64) error
	PushLikeArtist(userId string, artistId uint64) error
	PushLikeGenre(userId string, genreId uint64) error
	PushListenTrack(userId string, trackId uint64, dur uint32) error
	PushSkipTrack(userId string, trackId uint64, dur uint32) error
}

// Worker, который будет делать логику чтения и обработки активности
type WorkerUseCase interface {
	Run()
}
