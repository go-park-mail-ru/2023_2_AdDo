package websocket_wave

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"main/internal/common/response"
	"main/internal/common/utils"
	"main/internal/pkg/session"
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

var upgrader = websocket.Upgrader{} // use default options
func (h *Handler) MyWave(w http.ResponseWriter, r *http.Request) {
	h.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("My Wave Handler entered")

	sessionId, err := response.GetCookie(r)
	if err != nil {
		h.logger.Errorln("error get from cookie", sessionId, err)
		w.WriteHeader(401)
	}
	h.logger.Infoln("Got Cookie")

	userId, err := h.sessionUseCase.GetUserId(sessionId)
	if err != nil {
		w.WriteHeader(401)
		//return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	h.logger.Infoln("Got user id")

	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.logger.Errorln("error while upgrading connection to ws", err)
		w.WriteHeader(400)
	}
	defer conn.Close()
	h.logger.Infoln("Websocket connection opened")

	// нужна какая-то мапа, чтобы не давать дубликаты в рамках одной волны
	for {
		dummy := 0
		err := conn.ReadJSON(&dummy)
		if err != nil {
			h.logger.Errorln("error while reading message", err)
			w.WriteHeader(500)
			//return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
		}
		h.logger.Infoln("Got signal for next batch")

		if dummy == 0 {
			break
		}

		tracks, err := h.waveUseCase.GetMyWaveMusic(userId, MyWaveTrackBatch)
		if err != nil {
			h.logger.Errorln("error getting new track batch from wave micros", err)
			w.WriteHeader(500)
			//return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
		}
		h.logger.Infoln("Track Batch get from service")

		err = conn.WriteJSON(tracks)
		if err != nil {
			h.logger.Errorln("error sending a message", err)
			w.WriteHeader(500)
			//return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
		}
		h.logger.Infoln("Track Batch sent")
	}

	w.WriteHeader(204)
}
