package track_usecase

import (
	"main/internal/pkg/track"
)

type Default struct {
	repo track.Repository
}

func NewDefault(trackRepo track.Repository) Default {
	return Default{
		repo: trackRepo,
	}
}

func (useCase *Default) Add(t track.Track) (uint64, error) {
	id, err := useCase.repo.Create(t)
	if err != nil {
		return 0, track.ErrTrackAlreadyExist
	}
	return id, nil
}

func (useCase *Default) GetAll() ([]track.Track, error) {
	tracks, err := useCase.repo.GetAll()
	if err != nil {
		return nil, track.ErrNoTracks
	}
	return tracks, nil
}

func (useCase *Default) GetPopular() ([]track.Track, error) {
	//tracks, err := useCase.repo.GetMostLiked()
	//if err != nil {
	//	return nil, track.ErrNoTracks
	//}
	return nil, nil
}

func (useCase *Default) GetFavourite(userId uint64) ([]track.Track, error) {
	//tracks, err := useCase.repo.GetByUserId(userId)
	//if err != nil {
	//	return nil, track.ErrNoTracks
	//}
	return nil, nil
}
