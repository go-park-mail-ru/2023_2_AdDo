package activity_repository

import (
	"encoding/json"
	"errors"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/sirupsen/logrus"
	"main/internal/pkg/activity"
)

type MemCached struct {
	mcClient *memcache.Client
	logger   *logrus.Logger
}

const Activity = "ACTIVITY"
const Skip = "SKIP"
const New = "NEW"
const Old = "OLD"

func NewMemCached(mc *memcache.Client, l *logrus.Logger) MemCached {
	return MemCached{mcClient: mc, logger: l}
}

func (m MemCached) GetRecentSkip(userId string) ([]activity.UserTrackAction, error) {
	return m.GetTarget(userId, Skip, New)
}

func (m MemCached) GetRecentActivity(userId string) ([]activity.UserTrackAction, error) {
	return m.GetTarget(userId, Activity, New)
}

func (m MemCached) GetOldSkip(userId string) ([]activity.UserTrackAction, error) {
	return m.GetTarget(userId, Skip, Old)
}

func (m MemCached) GetOldActivity(userId string) ([]activity.UserTrackAction, error) {
	return m.GetTarget(userId, Activity, Old)
}

func (m MemCached) GetRecent(action activity.UserTrackAction) ([]activity.UserTrackAction, error) {
	if action.Action == activity.SkipAction {
		return m.GetRecentSkip(action.UserId)
	}
	return m.GetRecentActivity(action.UserId)
}

func (m MemCached) SetRecent(actions []activity.UserTrackAction) error {
	if actions[0].Action == activity.SkipAction {
		return m.SetTarget(actions, actions[0].UserId, Skip, New)
	}
	return m.SetTarget(actions, actions[0].UserId, Activity, New)
}

func (m MemCached) GetTarget(userId string, typ string, time string) ([]activity.UserTrackAction, error) {
	resultBin, err := m.mcClient.Get(userId + typ + time)
	if errors.Is(err, memcache.ErrCacheMiss) {
		return make([]activity.UserTrackAction, 0), nil
	}
	if err != nil {
		return make([]activity.UserTrackAction, 0), err
	}

	result := make([]activity.UserTrackAction, 0)
	err = json.Unmarshal(resultBin.Value, &result)
	if err != nil {
		m.logger.Errorln("Get Target ", err)
		return make([]activity.UserTrackAction, 0), err
	}

	return result, nil
}

func (m MemCached) SetTarget(actions []activity.UserTrackAction, userId string, typ string, time string) error {
	data, err := json.Marshal(actions)
	if err != nil {
		m.logger.Errorln("Marshalling data error ", err)
		return err
	}

	err = m.mcClient.Set(&memcache.Item{Key: userId + typ + time, Value: data, Expiration: 3600})
	if err != nil {
		m.logger.Errorln("Set Target error ", err)
		return err
	}

	return nil
}

func (m MemCached) SetAndCheck(action activity.UserTrackAction, count uint8) (bool, error) {
	m.logger.Infoln("Set All Recent Activity")

	result, err := m.GetRecent(action)
	result = append(result, action)

	err = m.SetRecent(result)
	if err != nil {
		m.logger.Errorln("Marshalling data  error ", err)
		return false, err
	}

	return len(result) == int(count), nil
}

const MinSizeInteresting = 3

func (m MemCached) GetAllActivity(userId string) ([]activity.UserTrackAction, []activity.UserTrackAction, error) {
	m.logger.Infoln("Get All Recent Activity")

	activities, _ := m.GetRecentActivity(userId)
	if len(activities) < MinSizeInteresting {
		temp, _ := m.GetOldActivity(userId)
		activities = append(activities, temp...)
	}
	m.logger.Errorln("ACTIVITIES", activities)

	skips, _ := m.GetRecentSkip(userId)
	if len(skips) < MinSizeInteresting {
		temp, _ := m.GetOldSkip(userId)
		skips = append(skips, temp...)
	}
	m.logger.Errorln("SKIP", skips)

	return activities, skips, nil
}

func (m MemCached) CleanRecentSkip(userId string) error {
	return m.CleanTarget(userId, Skip, New)
}

func (m MemCached) CleanRecentActivity(userId string) error {
	return m.CleanTarget(userId, Activity, New)
}

func (m MemCached) CleanTarget(userId string, typ string, time string) error {
	err := m.mcClient.Delete(userId + typ + time)
	if err != nil && !errors.Is(err, memcache.ErrCacheMiss) {
		m.logger.Errorln("Cleaning err", err)
		return err
	}

	return nil
}

func (m MemCached) CleanLastAndMerge(userId string) error {
	m.logger.Infoln("Clean Last And Merge")

	recSkip, _ := m.GetRecentSkip(userId)
	recAct, _ := m.GetRecentActivity(userId)
	m.logger.Errorln("RECSKIP", recSkip)
	m.logger.Errorln("RECACT", recAct)

	if len(recSkip) > len(recAct) {
		_ = m.SetTarget(recSkip, userId, Skip, Old)
		_ = m.CleanRecentSkip(userId)
	} else {
		_ = m.SetTarget(recAct, userId, Activity, Old)
		_ = m.CleanRecentActivity(userId)
	}

	return nil
}
