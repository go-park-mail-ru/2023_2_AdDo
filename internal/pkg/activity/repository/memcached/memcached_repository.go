package activity_repository

import (
	"main/internal/pkg/activity"
)

type MemCached struct {
}

func (m MemCached) SaveActivityAndCountCheck(action activity.UserTrackAction, count uint8) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m MemCached) GetAllRecentActivity(userId string) ([]activity.UserTrackAction, error) {
	//TODO implement me
	panic("implement me")
}

func (m MemCached) CleanLastActivityForUser(userId string) error {
	//TODO implement me
	panic("implement me")
}

func NewMemCached() MemCached {
	return MemCached{}
}
