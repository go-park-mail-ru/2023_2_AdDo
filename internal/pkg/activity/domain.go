package activity

const SkipAction = "SKIP"
const ListenAction = "LISTEN"
const LikeAction = "LIKE"
const RotationAction = "ROTATION"
const RecentActivityNeedToRecreateTrackPool = 5

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

type UserTrackAction struct {
	UserId  string
	TrackId uint64
	Action  string
}

type ConsumerRepository interface {
	PopLikeTrack(out chan<- UserTrackAction)
	PopLikeAlbum(out chan<- UserTrackAction)
	PopLikeArtist(out chan<- UserTrackAction)
	PopLikeGenre(out chan<- UserTrackAction)
	PopSkipTrack(out chan<- UserTrackAction)
	PopListenTrack(out chan<- UserTrackAction)
}

type ProducerRepository interface {
	PushLikeTrack(userId string, trackId uint64) error
	PushLikeAlbum(userId string, albumId uint64) error
	PushLikeArtist(userId string, artistId uint64) error
	PushLikeGenre(userId string, genreId uint64) error
	PushListenTrack(userId string, trackId uint64, dur uint32) error
	PushSkipTrack(userId string, trackId uint64, dur uint32) error
}

type KeyValueRepository interface {
	SetAndCheck(action UserTrackAction, count uint8) (bool, error)
	CleanLastAndMerge(userId string) error
	GetAllActivity(userId string) ([]UserTrackAction, []UserTrackAction, error)
}

type WorkerUseCase interface {
	Run()
}
