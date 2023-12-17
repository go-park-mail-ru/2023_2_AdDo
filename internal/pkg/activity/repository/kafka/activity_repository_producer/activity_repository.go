package activity_repository_producer

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
	"main/internal/pkg/activity"
)

type Default struct {
	logger *logrus.Logger
	queue  sarama.SyncProducer
}

func NewDefault(q sarama.SyncProducer, l *logrus.Logger) Default {
	return Default{
		logger: l,
		queue:  q,
	}
}

func (d *Default) PushListenTrack(userId string, trackId uint64, dur uint32) error {
	return d.pushMessage("ListenTrack", activity.UserListenTrack{UserId: userId, TrackId: trackId, DurationS: dur})
}

func (d *Default) PushSkipTrack(userId string, trackId uint64, dur uint32) error {
	return d.pushMessage("SkipTrack", activity.UserSkipTrack{UserId: userId, TrackId: trackId, DurationS: dur})
}

func (d *Default) PushLikeTrack(userId string, trackId uint64) error {
	return d.pushMessage("LikeTrack", activity.UserLikeTrack{UserId: userId, TrackId: trackId})
}

func (d *Default) PushLikeAlbum(userId string, albumId uint64) error {
	return d.pushMessage("LikeAlbum", activity.UserLikeAlbum{UserId: userId, AlbumId: albumId})
}

func (d *Default) PushLikeArtist(userId string, artistId uint64) error {
	return d.pushMessage("LikeArtist", activity.UserLikeArtist{ArtistId: artistId, UserId: userId})
}

func (d *Default) PushLikeGenre(userId string, genreId uint64) error {
	return d.pushMessage("LikeGenre", activity.UserLikeGenre{GenreId: genreId, UserId: userId})
}

func (d *Default) pushMessage(topic string, value any) error {
	d.logger.Infoln("push message kafka producer entered")
	valueBytes, err := json.Marshal(value)
	if err != nil {
		d.logger.Errorln("marshaling data error", err)
		return err
	}

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(valueBytes),
	}

	_, _, err = d.queue.SendMessage(message)
	if err != nil {
		d.logger.Errorln("messages sending error", err)
		return err
	}

	return nil
}
