package activity_repository_consumer

import (
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type Default struct {
	logger *logrus.Logger
	queue  sarama.Consumer
}

func NewDefault(q sarama.Consumer, l *logrus.Logger) Default {
	return Default{
		logger: l,
		queue:  q,
	}
}

func (d *Default) PullLikeTrack() {

}

func (d *Default) PullLikeAlbum() {

}

func (d *Default) PullLikeArtist() {

}

func (d *Default) PullLikeGenre() {

}

func (d *Default) PullSkipTrack() {

}

func (d *Default) PullListenTrack() {

}
