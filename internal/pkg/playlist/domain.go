package playlist

import (
	"context"
	"github.com/asaskevich/govalidator"
	xssvalidator "github.com/infiniteloopcloud/xss-validator"
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
	Id         uint64 `json:"Id" example:"1"`
	Name       string `json:"Name" example:"PlaylistName"`
	AuthorId   string `json:"AuthorId" example:"sdfa-asdf-adsf"`
	AuthorName string `json:"AuthorName" example:"username"`
	Preview    string `json:"Preview" example:"PlaylistPreview"`
	Tracks     []track.Response
}

type Name struct {
	Name string `valid:"length(1|30), required, printableascii" json:"Name" example:"PlaylistName"`
}

type ToTrackId struct {
	PlaylistId uint64 `json:"PlaylistId" example:"1"`
	TrackId    uint64 `json:"TrackId" example:"1"`
}

type IsCreator struct {
	IsCreator bool `json:"IsCreator" example:"true"`
}

func (u *Name) Validate() error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	err = xssvalidator.Validate(u.Name, xssvalidator.DefaultRules...)
	if err != nil {
		return err
	}

	return nil
}

type UseCase interface {
	Create(pl Base) (Response, error)
	Get(playlistId uint64) (Response, error)
	IsCreator(userId string, playlistId uint64) (bool, error)
	GetUserPlaylists(userId string) ([]Base, error)
	CollectionPlaylists(userId string) ([]Base, error)
	AddTrack(playlistId, trackId uint64) error
	Like(userId string, playlistId uint64) error
	IsLike(userId string, playlistId uint64) (bool, error)
	Unlike(userId string, playlistId uint64) error
	RemoveTrack(playlistId, trackId uint64) error
	UpdatePreview(playlistId uint64, src io.Reader, size int64) error
	RemovePreview(playlistId uint64) error
	DeleteById(playlistId uint64) error
	HasModifyAccess(userId string, playlistId uint64) (bool, error)
	HasReadAccess(playlistId uint64) (bool, error)
	MakePrivate(playlistId uint64) error
	MakePublic(playlistId uint64) error
	UpdateName(playlistId uint64, newName string) error
}

type Repository interface {
	Create(ctx context.Context, pl Base) (Response, error)
	Get(ctx context.Context, playlistId uint64) (Base, error)
	GetByCreatorId(ctx context.Context, userId string) ([]Base, error)
	GetByUserId(ctx context.Context, userId string) ([]Base, error)
	AddTrack(ctx context.Context, playlistId, trackId uint64) error
	RemoveTrack(ctx context.Context, playlistId, trackId uint64) error
	UpdateImage(ctx context.Context, playlistId uint64, image string) error
	RemovePreviewPath(ctx context.Context, playlistId uint64) (string, error)
	Delete(ctx context.Context, playlistId uint64) error
	CreateLike(ctx context.Context, userId string, playlistId uint64) error
	CheckLike(ctx context.Context, userId string, playlistId uint64) (bool, error)
	DeleteLike(ctx context.Context, userId string, playlistId uint64) error
	IsCreator(ctx context.Context, userId string, playlistId uint64) (bool, error)
	IsPrivate(ctx context.Context, playlistId uint64) (bool, error)
	MakePublic(ctx context.Context, playlistId uint64) error
	MakePrivate(ctx context.Context, playlistId uint64) error
	Search(ctx context.Context, text string) ([]Base, error)
	UpdateName(ctx context.Context, playlistId uint64, title string) error
}
