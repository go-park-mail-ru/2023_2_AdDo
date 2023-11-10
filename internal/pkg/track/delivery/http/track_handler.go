package track_delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"main/internal/common/handler"
	"main/internal/common/response"
	"main/internal/common/utils"
	"main/internal/pkg/album"
	"main/internal/pkg/session"
	"main/internal/pkg/track"
	"net/http"
	"strconv"
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
//	@Summary		Listen
//	@Description	Listen track
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

	if err := handler.trackUseCase.Listen(trackId.Id); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("play count for track ", trackId, "incremented")

	w.WriteHeader(http.StatusNoContent)
	return nil
}

// Like
//
//	@Summary		Like
//	@Description	Like track
//	@Tags			track
//	@Security		cookieAuth
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Param			id	path	integer	true	"track id"
//	@Success		204
//	@Failure		400	{string}	errMsg
//	@Failure		401	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/track/{id}/like [post]
func (handler *TrackHandler) Like(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("Like TrackHandler entered")

	trackId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed trackId from Vars")

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

	if err = handler.trackUseCase.Like(userId, uint64(trackId)); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("like for track ", trackId, "created")

	w.WriteHeader(http.StatusNoContent)
	return nil
}

// IsLike
//
//	@Summary		IsLike
//	@Description	Check if track is liked
//	@Tags			track
//	@Security		cookieAuth
//	@Produce		json
//	@Param			id	path		integer	true	"track id"
//	@Success		200	{object}	response.IsLiked
//	@Failure		400	{string}	errMsg
//	@Failure		401	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/track/{id}/is_like [get]
func (handler *TrackHandler) IsLike(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("Like TrackHandler entered")

	trackId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed trackId from Vars")

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

	isLiked, err := handler.trackUseCase.IsLike(userId, uint64(trackId))
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("like for track ", trackId, "checked")

	if err = response.RenderJSON(w, response.IsLiked{IsLiked: isLiked}); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("response  formed")

	return nil
}

// Unlike
//
//	@Summary		Unlike
//	@Description	Unlike track
//	@Tags			track
//	@Security		cookieAuth
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Param			id	path	integer	true	"track id"
//	@Success		204
//	@Failure		400	{string}	errMsg
//	@Failure		401	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/track/{id}/unlike [delete]
func (handler *TrackHandler) Unlike(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("Like TrackHandler entered")

	trackId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed trackId from Vars")

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

	if err = handler.trackUseCase.Unlike(userId, uint64(trackId)); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("like for track ", trackId, "deleted")

	w.WriteHeader(http.StatusNoContent)
	return nil
}

// GetUserTracks
//
//	@Summary		GetUserTracks
//	@Description	Get user tracks
//	@Tags			track
//	@Produce		json
//	@Security		cookieAuth
//	@Success		200 {object}	track.LikedTracks
//	@Failure		400	{string}	errMsg
//	@Failure		401	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/collection/tracks [get]
func (handler *TrackHandler) GetUserTracks(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("GetUserTracks TrackHandler entered")

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

	result, err := handler.trackUseCase.GetUserLikedTracks(userId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("liked tracks for user ", userId, "get")

	if err = response.RenderJSON(w, track.LikedTracks{Tracks: result}); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("response  formed")
	return nil
}
