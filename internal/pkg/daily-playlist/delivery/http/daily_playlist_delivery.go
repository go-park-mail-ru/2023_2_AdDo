package daily_playlist_delivery

import (
	"github.com/mailru/easyjson"
	"github.com/sirupsen/logrus"
	common_handler "main/internal/common/handler"
	"main/internal/common/response"
	"main/internal/common/utils"
	daily_playlist "main/internal/pkg/daily-playlist"
	"main/internal/pkg/session"
	"net/http"
)

type Handler struct {
	logger               *logrus.Logger
	sessionUseCase       session.UseCase
	dailyPlaylistUseCase daily_playlist.ServiceUseCase
}

func NewHandler(dpu daily_playlist.ServiceUseCase, su session.UseCase, l *logrus.Logger) Handler {
	return Handler{
		logger:               l,
		sessionUseCase:       su,
		dailyPlaylistUseCase: dpu,
	}
}

func (h *Handler) GetDailyPlaylist(w http.ResponseWriter, r *http.Request) error {
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

	result, err := h.dailyPlaylistUseCase.GetUserDailyPlaylist(userId)
	if err != nil {
		h.logger.Errorln("error while getting daily", err)
		return common_handler.StatusError{Code: http.StatusNotFound, Err: err}
	}
	h.logger.Infoln("Got daily playlist Id")

	_, _, err = easyjson.MarshalToHTTPResponseWriter(result, w)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	h.logger.Infoln("Response formed")

	return nil
}
