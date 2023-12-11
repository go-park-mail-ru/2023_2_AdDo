package activity_repository

import (
	"encoding/json"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/sirupsen/logrus"
	"main/internal/pkg/activity"
)

type MemCached struct {
	mcClient *memcache.Client
	logger   *logrus.Logger
}

func NewMemCached(mc *memcache.Client, l *logrus.Logger) MemCached {
	return MemCached{mcClient: mc, logger: l}
}

func (m MemCached) SaveActivityAndCountCheck(action activity.UserTrackAction, count uint8) (bool, error) {
	m.logger.Infoln("Set All Recent Activity")

	data, err := json.Marshal(action)
	if err != nil {
		m.logger.Errorln("Marshalling data  error ", err)
		return false, err
	}

	err = m.mcClient.Set(&memcache.Item{Key: action.UserId, Value: data, Expiration: 3600})
	if err != nil {
		m.logger.Errorln("Get All Recent Activity error ", err)
		return false, err
	}

	result, err := m.GetAllRecentActivity(action.UserId)
	if err != nil {
		m.logger.Errorln("Get All Recent Activity error ", err)
		return false, err
	}

	return len(result) == int(count), nil
}

func (m MemCached) GetAllRecentActivity(userId string) ([]activity.UserTrackAction, error) {
	m.logger.Infoln("Get All Recent Activity")

	result, err := m.mcClient.Get(userId)
	if err != nil {
		m.logger.Errorln("Get All Recent Activity error ", err)
		return nil, err
	}

	activities := make([]activity.UserTrackAction, 0)
	err = json.Unmarshal(result.Value, &activities)
	if err != nil {
		m.logger.Errorln("Get All Recent Activity error ", err)
		return nil, err
	}

	return activities, nil
}

func (m MemCached) CleanLastActivityForUser(userId string) error {
	m.logger.Infoln("Clean All Recent Activity For User")

	err := m.mcClient.Delete(userId)
	if err != nil {
		m.logger.Errorln("Get All Recent Activity error ", err)
		return err
	}

	return nil
}
