package playlist_delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	common_handler "main/internal/common/handler"
	"main/internal/common/response"
	"main/internal/common/utils"
	"main/internal/pkg/playlist"
	"main/internal/pkg/session"
	"net/http"
	"strconv"
)

type Handler struct {
	playlistUseCase playlist.UseCase
	sessionUseCase  session.UseCase
	logger          *logrus.Logger
}

func NewHandler(pu playlist.UseCase, su session.UseCase, logger *logrus.Logger) Handler {
	return Handler{
		playlistUseCase: pu,
		sessionUseCase:  su,
		logger:          logger,
	}
}

func (handler *Handler) Create(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("PlaylistCreate Handler entered")

	sessionId, err := response.GetCookie(r)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	handler.logger.Infoln("Got Cookie")

	userId, err := handler.sessionUseCase.GetUserId(sessionId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	handler.logger.Infoln("Got user id")

	var base playlist.Base
	if err := json.NewDecoder(r.Body).Decode(&base); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	base.AuthorId = userId
	handler.logger.Infoln("Got base playlist formed")

	err = handler.playlistUseCase.Create(base)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (handler *Handler) Get(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("PlaylistGet Handler entered")

	sessionId, err := response.GetCookie(r)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	handler.logger.Infoln("Got Cookie")

	userId, err := handler.sessionUseCase.GetUserId(sessionId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	handler.logger.Infoln("Got user id")

	playlistId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed playlistId from Vars")

	result, err := handler.playlistUseCase.Get(userId, uint64(playlistId))
	if err != nil {
		return common_handler.StatusError{Code: http.StatusNotFound, Err: err}
	}

	err = response.RenderJSON(w, result)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	return nil
}

func (handler *Handler) AddTrack(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("PlaylistAddTrack Handler entered")

	var ids playlist.ToTrackId
	if err := json.NewDecoder(r.Body).Decode(&ids); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Got playlist and track ids")

	err := handler.playlistUseCase.AddTrack(ids.PlaylistId, ids.TrackId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (handler *Handler) RemoveTrack(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("PlaylistRemoveTrack Handler entered")

	var ids playlist.ToTrackId
	if err := json.NewDecoder(r.Body).Decode(&ids); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Got playlist and track ids")

	err := handler.playlistUseCase.RemoveTrack(ids.PlaylistId, ids.TrackId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (handler *Handler) UpdatePreview(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("PlaylistAddTrack Handler entered")

	playlistId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed playlistId from Vars")

	src, hdr, err := r.FormFile("Preview")
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	defer src.Close()
	handler.logger.Infoln("formed file")

	err = handler.playlistUseCase.UpdatePreview(uint64(playlistId), src, hdr.Size)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (handler *Handler) Like(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("LikePlaylist Handler entered")

	var id playlist.Id
	if err := json.NewDecoder(r.Body).Decode(&id); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("got track id from body")

	sessionId, err := response.GetCookie(r)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	handler.logger.Infoln("got user cookie")

	userId, err := handler.sessionUseCase.GetUserId(sessionId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	handler.logger.Infoln("got user id by cookie")

	err = handler.playlistUseCase.Like(userId, id.Id)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("like created successfully")

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (handler *Handler) Delete(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("PlaylistDelete Handler entered")

	playlistId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed playlistId from Vars")

	err = handler.playlistUseCase.DeleteById(uint64(playlistId))
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (handler *Handler) MakePublic(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("MakePublic Handler entered")

	playlistId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed playlistId from Vars")

	err = handler.playlistUseCase.MakePublic(uint64(playlistId))
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (handler *Handler) MakePrivate(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("MakePrivate Handler entered")

	playlistId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed playlistId from Vars")

	err = handler.playlistUseCase.MakePrivate(uint64(playlistId))
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
