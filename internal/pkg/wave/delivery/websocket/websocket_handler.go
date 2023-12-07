package websocket_wave

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	common_handler "main/internal/common/handler"
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

func (h *Handler) MyWave(w http.ResponseWriter, r *http.Request) error {
	h.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("My Wave Handler entered")

	sessionId, err := response.GetCookie(r)
	if err != nil {
		h.logger.Errorln("error get from cookie", sessionId, err)
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	h.logger.Infoln("Got Cookie")

	userId, err := h.sessionUseCase.GetUserId(sessionId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	h.logger.Infoln("Got user id")

	conn, err := h.upgrader.Upgrade(w, r, nil)
	defer conn.Close()
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	h.logger.Infoln("Websocket connection opened")

	// нужна какая-то мапа, чтобы не давать дубликаты в рамках одной волны
	// нужно понимать, когда пользователь прослушал некоторую часть треков, чтобы отправить ему новые
	tracks, err := h.waveUseCase.GetMyWaveMusic(userId, MyWaveTrackBatch)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	h.logger.Infoln("Track Batch get from service")

	err = conn.WriteJSON(tracks)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	h.logger.Infoln("Track Batch sent")

	return nil
}
