package track_delivery

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"main/internal/common/handler"
	"main/internal/common/response"
	"main/internal/common/utils"
	"main/internal/pkg/album"
	"main/internal/pkg/session"
	"main/internal/pkg/track"
	"net/http"
)

type TrackHandler struct {
	trackUseCase   track.UseCase
	albumUseCase   album.UseCase
	sessionUseCase session.UseCase
	logger         *logrus.Logger
}

func NewHandler(track track.UseCase, session session.UseCase, logger *logrus.Logger) TrackHandler {
	return TrackHandler{
		trackUseCase:   track,
		sessionUseCase: session,
		logger:         logger,
	}
}

// Listen
//
//	@Description	listen track
//	@Tags			track
//	@Accept			json
//	@Param			trackId	body	integer	true	"track id"
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Success		204
//	@Failure		400	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/listen [post]
func (handler *TrackHandler) Listen(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("Listen Handler entered")

	var trackId track.Id
	if err := json.NewDecoder(r.Body).Decode(&trackId); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("track id decoded")

	err := handler.trackUseCase.Listen(trackId.Id)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("play count for track ", trackId, "incremented")

	w.WriteHeader(http.StatusNoContent)
	return nil
}

// Like
//
//	@Description	like track
//	@Tags			track
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Accept			json
//	@Param			trackId	body		integer	true	"track id"
//	@Success		200		{object}	track.Response
//	@Failure		400		{string}	errMsg
//	@Failure		401		{string}	errMsg
//	@Failure		403		{string}	errMsg
//	@Failure		500		{string}	errMsg
//	@Router			/like [post]
func (handler *TrackHandler) Like(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("Like TrackHandler entered")

	var trackId track.Id
	if err := json.NewDecoder(r.Body).Decode(&trackId); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("track id decoded")

	sessionId, err := response.GetCookie(r)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	handler.logger.Infoln("got cookie")

	userId, err := handler.sessionUseCase.GetUserId(sessionId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	handler.logger.Infoln("got user id by session id")

	err = handler.trackUseCase.Like(userId, trackId.Id)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("like for track ", trackId, "created")

	w.WriteHeader(http.StatusNoContent)
	return nil
}
