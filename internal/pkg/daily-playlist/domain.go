package daily_playlist

import "main/internal/pkg/track"

type Response struct {
	Id      uint64
	OwnerId string
	Tracks  []track.Response
}

type Repository interface {
	GetUserDailyPlaylist(userId string) (Response, error)
	SetUserPlaylist(userId string, response Response) error
}

type ServiceUseCase interface {
	GetUserDailyPlaylist(userId string) (Response, error)
}

type WorkerUseCase interface {
	CreateDailyPlaylistForUsers()
}
