package websocket_wave

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"main/internal/common/response"
	"main/internal/pkg/session"
	"main/internal/pkg/track"
	"main/internal/pkg/wave"
	"net/http"
)

type Handler struct {
	upgrader       websocket.Upgrader
	logger         *logrus.Logger
	sessionUseCase session.UseCase
	waveUseCase    wave.UseCase
}

func NewHandler(u websocket.Upgrader, su session.UseCase, ru wave.UseCase, l *logrus.Logger) Handler {
	return Handler{
		upgrader:       u,
		logger:         l,
		sessionUseCase: su,
		waveUseCase:    ru,
	}
}

const MyWaveTrackBatch = 5

func (h *Handler) MyWave(w http.ResponseWriter, r *http.Request) {
	h.logger.Infoln("My Wave handler entered")

	sessionId, err := response.GetCookie(r)
	if err != nil {
		w.WriteHeader(401)
		h.logger.Errorln("error get from cookie", sessionId, err)
		return
	}
	h.logger.Infoln("Got Cookie")

	userId, err := h.sessionUseCase.GetUserId(sessionId)
	if err != nil {
		w.WriteHeader(401)
		h.logger.Errorln("error got user id", err)
		return
	}
	h.logger.Infoln("Got user id")

	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.logger.Errorln("error while upgrading connection to ws", err)
		w.WriteHeader(400)
		return
	}
	defer conn.Close()
	h.logger.Infoln("Websocket connection opened")

	uniqTracks := make(map[uint64]bool)
	for {
		dummy := 0
		err := conn.ReadJSON(&dummy)
		if err != nil {
			h.logger.Errorln("error while reading message", err)
			w.WriteHeader(500)
		}
		h.logger.Infoln("Got signal for next batch")

		if dummy == 0 {
			break
		}

		tracks, err := h.getNextBatch(userId, uniqTracks)
		if err != nil {
			h.logger.Errorln("error getting new batch", err)
			w.WriteHeader(500)
			return
		}
		h.logger.Errorln("uniq tracks map", uniqTracks)

		err = conn.WriteJSON(tracks)
		if err != nil {
			h.logger.Errorln("error sending a message", err)
			w.WriteHeader(500)
			return
		}
		h.logger.Infoln("Track Batch sent")
	}

	w.WriteHeader(204)
}

func deleteDuplicatesFromWave(uniqTracks map[uint64]bool, vec []track.Response) []track.Response {
	result := make([]track.Response, 0)
	for _, t := range vec {
		if _, ok := uniqTracks[t.Id]; ok {
			continue
		}
		uniqTracks[t.Id] = true
		result = append(result, t)
	}

	return result
}

func (h *Handler) getNextBatch(userId string, uniqTracks map[uint64]bool) ([]track.Response, error) {
	result := make([]track.Response, 0)
	mult := 1
	for len(result) < 5 {
		tracks, err := h.waveUseCase.GetMyWaveMusic(userId, MyWaveTrackBatch)
		if err != nil {
			h.logger.Errorln("error getting new track batch from wave micros", err)
			return nil, err
		}

		temp := deleteDuplicatesFromWave(uniqTracks, tracks)
		result = append(result, temp...)
		h.logger.Infoln("Result Is ", result, "Mult is ", mult)
		mult++
	}

	return result, nil
}
