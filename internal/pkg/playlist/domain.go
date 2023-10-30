package playlist

import (
	"context"
	"io"
	"main/internal/pkg/track"
)

type Id struct {
	Id uint64 `json:"Id" example:"1"`
}

type Base struct {
	Id       uint64 `json:"Id" example:"1"`
	Name     string `json:"Name" example:"PlaylistName"`
	AuthorId string `json:"AuthorId" example:"sdfa-asdf-adsf"`
	Preview  string `json:"Preview" example:"PlaylistPreview"`
}

type Response struct {
	Id       uint64 `json:"Id" example:"1"`
	Name     string `json:"Name" example:"PlaylistName"`
	AuthorId string `json:"AuthorId" example:"sdfa-asdf-adsf"`
	Preview  string `json:"Preview" example:"PlaylistPreview"`
	Tracks   []track.Response
}

type ToTrackId struct {
	PlaylistId uint64 `json:"PlaylistId" example:"1"`
	TrackId    uint64 `json:"TrackId" example:"1"`
}

type UseCase interface {
	Create(playlist Base) error
	Get(playlistId uint64) (Response, error)
	GetUserPlaylists(userId string) ([]Base, error)
	AddTrack(playlistId, trackId uint64) error
	Like(userId string, playlistId uint64) error
	RemoveTrack(playlistId, trackId uint64) error
	UpdatePreview(playlistId uint64, src io.Reader, size int64) error
	RemovePreview(playlistId uint64) error
	DeleteById(playlistId uint64) error
}

type Repository interface {
	Create(ctx context.Context, playlist Base) error
	Get(ctx context.Context, playlistId uint64) (Base, error)
	GetByCreatorId(ctx context.Context, userId string) ([]Base, error)
	AddTrack(ctx context.Context, playlistId, trackId uint64) error
	RemoveTrack(ctx context.Context, playlistId, trackId uint64) error
	UpdateImage(ctx context.Context, playlistId uint64, image string) error
	RemovePreviewPath(ctx context.Context, playlistId uint64) (string, error)
	Delete(ctx context.Context, playlistId uint64) error
	CreateLike(ctx context.Context, userId string, playlistId uint64) error
}
