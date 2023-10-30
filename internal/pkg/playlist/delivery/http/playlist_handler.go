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

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) error {
	h.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("PlaylistCreate Handler entered")

	sessionId, err := response.GetCookie(r)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	h.logger.Infoln("Got Cookie")

	userId, err := h.sessionUseCase.GetUserId(sessionId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	h.logger.Infoln("Got user id")

	var base playlist.Base
	if err := json.NewDecoder(r.Body).Decode(&base); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	base.AuthorId = userId
	h.logger.Infoln("Got base playlist formed")

	err = h.playlistUseCase.Create(base)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) error {
	h.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("PlaylistGet Handler entered")

	playlistId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	h.logger.Infoln("Parsed playlistId from Vars")

	result, err := h.playlistUseCase.Get(uint64(playlistId))
	if err != nil {
		return common_handler.StatusError{Code: http.StatusNotFound, Err: err}
	}

	err = response.RenderJSON(w, result)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	return nil
}

func (h *Handler) AddTrack(w http.ResponseWriter, r *http.Request) error {
	h.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("PlaylistAddTrack Handler entered")

	var ids playlist.ToTrackId
	if err := json.NewDecoder(r.Body).Decode(&ids); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	h.logger.Infoln("Got playlist and track ids")

	err := h.playlistUseCase.AddTrack(ids.PlaylistId, ids.TrackId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (h *Handler) RemoveTrack(w http.ResponseWriter, r *http.Request) error {
	h.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("PlaylistRemoveTrack Handler entered")

	var ids playlist.ToTrackId
	if err := json.NewDecoder(r.Body).Decode(&ids); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	h.logger.Infoln("Got playlist and track ids")

	err := h.playlistUseCase.RemoveTrack(ids.PlaylistId, ids.TrackId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) error {
	h.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("PlaylistDelete Handler entered")

	playlistId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	h.logger.Infoln("Parsed playlistId from Vars")

	err = h.playlistUseCase.DeleteById(uint64(playlistId))
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
