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
	"main/internal/pkg/track"
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
	// if err := json.NewDecoder(r.Body).Decode(&base); err != nil {
	// 	return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	// }
	base.AuthorId = userId

	result, err := handler.playlistUseCase.Create(base)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	if err = response.RenderJSON(w, result); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	return nil
}

func (handler *Handler) Get(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("PlaylistGet Handler entered")

	sessionId, err := response.GetCookie(r)
	if err != nil {
		handler.logger.Errorln("error get from cookie", sessionId, err)
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

	if err = response.RenderJSON(w, result); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	return nil
}

func (handler *Handler) AddTrack(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("PlaylistAddTrack Handler entered")

	playlistId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed playlistId from Vars")

	var ids track.Id
	if err := json.NewDecoder(r.Body).Decode(&ids); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Got playlist and track ids")

	if err = handler.playlistUseCase.AddTrack(uint64(playlistId), ids.Id); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (handler *Handler) RemoveTrack(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("PlaylistRemoveTrack Handler entered")

	playlistId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed playlistId from Vars")

	var ids track.Id
	if err := json.NewDecoder(r.Body).Decode(&ids); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Got playlist and track ids")

	if err = handler.playlistUseCase.RemoveTrack(uint64(playlistId), ids.Id); err != nil {
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

	if err = handler.playlistUseCase.UpdatePreview(uint64(playlistId), src, hdr.Size); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (handler *Handler) Like(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("LikePlaylist Handler entered")

	playlistId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed playlistId from Vars")

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

	if err = handler.playlistUseCase.Like(userId, uint64(playlistId)); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("playlist like created successfully")

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (handler *Handler) IsLike(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("LikePlaylist Handler entered")

	playlistId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed playlistId from Vars")

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

	isLiked, err := handler.playlistUseCase.IsLike(userId, uint64(playlistId))
	if err != nil {
		return common_handler.StatusError{Code: http.StatusNotFound, Err: err}
	}
	handler.logger.Infoln("User like checked")

	if err = response.RenderJSON(w, response.IsLiked{IsLiked: isLiked}); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("response  formed")

	return nil
}

func (handler *Handler) Unlike(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("LikePlaylist Handler entered")

	playlistId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed playlistId from Vars")

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

	if err = handler.playlistUseCase.Unlike(userId, uint64(playlistId)); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("playlist like deleted successfully")

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

	if err = handler.playlistUseCase.DeleteById(uint64(playlistId)); err != nil {
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

	if err = handler.playlistUseCase.MakePublic(uint64(playlistId)); err != nil {
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

	if err = handler.playlistUseCase.MakePrivate(uint64(playlistId)); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
