package activity_repository_consumer

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
	"main/internal/pkg/activity"
	"main/internal/pkg/track"
)

type Default struct {
	logger    *logrus.Logger
	queue     sarama.Consumer
	trackRepo track.Repository
}

func NewDefault(cr track.Repository, q sarama.Consumer, l *logrus.Logger) Default {
	return Default{
		logger:    l,
		queue:     q,
		trackRepo: cr,
	}
}

func (d *Default) PopLikeTrack(out chan<- activity.UserTrackAction) {
	partitionConsumer, err := d.queue.ConsumePartition("LikeTrack", 0, sarama.OffsetOldest)
	if err != nil {
		d.logger.Errorln("error while creating partition LikeTrack", err)
	}
	defer partitionConsumer.Close()

	for {
		msg := <-partitionConsumer.Messages()
		var like activity.UserLikeTrack
		err := json.Unmarshal(msg.Value, &like)
		if err != nil {
			d.logger.Errorln("decoding error", err)
		}

		out <- activity.UserTrackAction{
			UserId:  like.UserId,
			TrackId: like.TrackId,
			Action:  activity.LikeAction,
		}
		d.logger.Infoln("Got messages", string(msg.Key), string(msg.Value))
	}
}

func (d *Default) PopLikeAlbum(out chan<- activity.UserTrackAction) {
	partitionConsumer, err := d.queue.ConsumePartition("LikeAlbum", 0, sarama.OffsetOldest)
	if err != nil {
		d.logger.Errorln("error while creating partition LikeAlbum", err)
	}
	defer partitionConsumer.Close()

	for {
		msg := <-partitionConsumer.Messages()
		var like activity.UserLikeAlbum
		err := json.Unmarshal(msg.Value, &like)
		if err != nil {
			d.logger.Errorln("decoding error", err)
		}

		id, err := d.trackRepo.GetRotationTrackForAlbum(like.AlbumId)
		if err != nil {
			d.logger.Errorln("Error getting nearest track for album", err)
		}

		out <- activity.UserTrackAction{
			UserId:  like.UserId,
			TrackId: id,
			Action:  activity.RotationAction,
		}
		d.logger.Infoln("Got messages", string(msg.Value))
	}
}

func (d *Default) PopLikeArtist(out chan<- activity.UserTrackAction) {
	partitionConsumer, err := d.queue.ConsumePartition("LikeArtist", 0, sarama.OffsetOldest)
	if err != nil {
		d.logger.Errorln("error while creating partition LikeArtist", err)
	}
	defer partitionConsumer.Close()

	for {
		msg := <-partitionConsumer.Messages()
		var like activity.UserLikeArtist
		err := json.Unmarshal(msg.Value, &like)
		if err != nil {
			d.logger.Errorln("decoding error", err)
		}

		id, err := d.trackRepo.GetRotationTrackForArtist(like.ArtistId)
		if err != nil {
			d.logger.Errorln("Error getting nearest track for artist", err)
		}

		out <- activity.UserTrackAction{
			UserId:  like.UserId,
			TrackId: id,
			Action:  activity.RotationAction,
		}
		d.logger.Infoln("Got messages", string(msg.Value))
	}
}

func (d *Default) PopLikeGenre(out chan<- activity.UserTrackAction) {
	partitionConsumer, err := d.queue.ConsumePartition("LikeGenre", 0, sarama.OffsetOldest)
	if err != nil {
		d.logger.Errorln("error while creating partition LikeGenre", err)
	}
	defer partitionConsumer.Close()

	for {
		msg := <-partitionConsumer.Messages()
		var like activity.UserLikeGenre
		err := json.Unmarshal(msg.Value, &like)
		if err != nil {
			d.logger.Errorln("decoding error", err)
		}

		id, err := d.trackRepo.GetRotationTrackForGenre(like.GenreId)
		if err != nil {
			d.logger.Errorln("Error getting nearest track for album", err)
		}

		out <- activity.UserTrackAction{
			UserId:  like.UserId,
			TrackId: id,
			Action:  activity.RotationAction,
		}
		d.logger.Infoln("Got messages", string(msg.Value))
	}
}

func (d *Default) PopSkipTrack(out chan<- activity.UserTrackAction) {
	partitionConsumer, err := d.queue.ConsumePartition("SkipTrack", 0, sarama.OffsetOldest)
	if err != nil {
		d.logger.Errorln("error while creating partition SkipTrack", err)
	}
	defer partitionConsumer.Close()

	for {
		msg := <-partitionConsumer.Messages()
		var like activity.UserSkipTrack
		err := json.Unmarshal(msg.Value, &like)
		if err != nil {
			d.logger.Errorln("decoding error", err)
		}

		d.logger.Infoln("Got messages", string(msg.Key), string(msg.Value))

		out <- activity.UserTrackAction{
			UserId:  like.UserId,
			TrackId: like.TrackId,
			Action:  activity.SkipAction,
		}
		d.logger.Infoln("Got messages", string(msg.Key), string(msg.Value))
	}
}

func (d *Default) PopListenTrack(out chan<- activity.UserTrackAction) {
	partitionConsumer, err := d.queue.ConsumePartition("ListenTrack", 0, sarama.OffsetOldest)
	if err != nil {
		d.logger.Errorln("error while creating partition ListenTrack", err)
	}
	defer partitionConsumer.Close()

	for {
		msg := <-partitionConsumer.Messages()
		var like activity.UserListenTrack
		err := json.Unmarshal(msg.Value, &like)
		if err != nil {
			d.logger.Errorln("decoding error", err)
		}

		out <- activity.UserTrackAction{
			UserId:  like.UserId,
			TrackId: like.TrackId,
			Action:  activity.ListenAction,
		}
		d.logger.Infoln("Got messages", string(msg.Key), string(msg.Value))
	}
}
