package daily_playlist

import "main/internal/pkg/playlist"

type Repository interface {
	GetUserDailyPlaylist(userId string) (playlist.Response, error)
	SetUserPlaylist(userId string, response playlist.Response) error
}

type ServiceUseCase interface {
	GetUserDailyPlaylist(userId string) (playlist.Response, error)
}

type WorkerUseCase interface {
	Run() error
}
