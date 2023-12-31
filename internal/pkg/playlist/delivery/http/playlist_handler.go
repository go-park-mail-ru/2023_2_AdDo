package playlist_delivery

import (
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
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
	trackUseCase    track.UseCase
	logger          *logrus.Logger
}

func NewHandler(pu playlist.UseCase, tu track.UseCase, su session.UseCase, logger *logrus.Logger) Handler {
	return Handler{
		playlistUseCase: pu,
		sessionUseCase:  su,
		logger:          logger,
		trackUseCase:    tu,
	}
}

// Create
//
//	@Summary		Create
//	@Description	Create playlist
//	@Tags			playlist
//	@Security		cookieAuth
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Produce		json
//	@Success		200	{object}	playlist.Response
//	@Failure		401	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/playlist [post]
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
	base.AuthorId = userId

	result, err := handler.playlistUseCase.Create(base)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	if _, _, err = easyjson.MarshalToHTTPResponseWriter(result, w); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	return nil
}

// Get
//
//	@Summary		Get
//	@Description	Return playlist info and tracks
//	@Tags			playlist
//	@Produce		json
//	@Param			id	path		integer	true	"playlist id"
//	@Success		200	{object}	playlist.Response
//	@Failure		400	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		404	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/playlist/{id} [get]
func (handler *Handler) Get(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("PlaylistGet Handler entered")

	sessionId, _ := response.GetCookie(r)
	userId, _ := handler.sessionUseCase.GetUserId(sessionId)

	playlistId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed playlistId from Vars")

	result, err := handler.playlistUseCase.Get(uint64(playlistId))
	if err != nil {
		return common_handler.StatusError{Code: http.StatusNotFound, Err: err}
	}

	if userId != "" {
		labeledTracks, err := handler.trackUseCase.LabelIsLikedTracks(userId, result.Tracks)
		if err != nil {
			handler.logger.Errorln("Error Labeling Tracks with IsLiked")
		}
		result.Tracks = labeledTracks
	}

	if _, _, err = easyjson.MarshalToHTTPResponseWriter(result, w); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	if userId == "" {
		w.WriteHeader(http.StatusUnauthorized)
	}

	return nil
}

// IsCreator
//
//	@Summary		IsCreator
//	@Description	Check if user is creator of playlist
//	@Tags			playlist
//	@Security		cookieAuth
//	@Produce		json
//	@Param			id	path		integer	true	"playlist id"
//	@Success		200	{object}	playlist.IsCreator
//	@Failure		400	{string}	errMsg
//	@Failure		401	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		404	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/playlist/{id}/is_creator [get]
func (handler *Handler) IsCreator(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("IsCreator Handler entered")

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

	isCreator, err := handler.playlistUseCase.IsCreator(userId, uint64(playlistId))
	if err != nil {
		return common_handler.StatusError{Code: http.StatusNotFound, Err: err}
	}
	handler.logger.Infoln("Checked whether the user is creator of playlist")

	if _, _, err = easyjson.MarshalToHTTPResponseWriter(playlist.IsCreator{IsCreator: isCreator}, w); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("response  formed")

	return nil
}

// AddTrack
//
//	@Summary		AddTrack
//	@Description	Add track to playlist
//	@Tags			playlist
//	@Security		cookieAuth
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Accept			json
//	@Param			id		path	integer		true	"playlist id"
//	@Param			trackId	body	track.Id	true	"track id"
//	@Success		204
//	@Failure		400	{string}	errMsg
//	@Failure		401	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/playlist/{id}/add_track [post]
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
	if err = easyjson.UnmarshalFromReader(r.Body, &ids); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Got playlist and track ids")

	if err = handler.playlistUseCase.AddTrack(uint64(playlistId), ids.Id); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

// RemoveTrack
//
//	@Summary		RemoveTrack
//	@Description	Remove track from playlist
//	@Tags			playlist
//	@Security		cookieAuth
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Accept			json
//	@Param			id		path	integer		true	"playlist id"
//	@Param			trackId	body	track.Id	true	"track id"
//	@Success		204
//	@Failure		400	{string}	errMsg
//	@Failure		401	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/playlist/{id}/remove_track [delete]
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
	if err = easyjson.UnmarshalFromReader(r.Body, &ids); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Got playlist and track ids")

	if err = handler.playlistUseCase.RemoveTrack(uint64(playlistId), ids.Id); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

// UpdatePreview
//
//	@Summary		UpdatePreview
//	@Description	Update playlist preview
//	@Tags			playlist
//	@Security		cookieAuth
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Param			id		path		integer	true	"playlist id"
//	@Param			image	formData	file	true	"preview image"
//	@Success		204
//	@Failure		400	{string}	errMsg
//	@Failure		401	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/playlist/{id}/update_preview [post]
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

// UpdateName
//
//	@Summary		UpdateName
//	@Description	Update playlist name
//	@Tags			playlist
//	@Security		cookieAuth
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Param			id		path	integer			true	"playlist id"
//	@Param			name	body	playlist.Name	true	"playlist name"
//	@Success		204
//	@Failure		400	{string}	errMsg
//	@Failure		401	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/playlist/{id}/update_name [post]
func (handler *Handler) UpdateName(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("Playlist Update Name Handler entered")

	playlistId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed playlistId from Vars")

	var title playlist.Name
	if err = easyjson.UnmarshalFromReader(r.Body, &title); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Got playlist and new title")

	if err := title.Validate(); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("title validated")

	if err = handler.playlistUseCase.UpdateName(uint64(playlistId), title.Name); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

// Like
//
//	@Summary		Like
//	@Description	Like playlist
//	@Tags			playlist
//	@Security		cookieAuth
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Param			id	path	integer	true	"playlist id"
//	@Success		204
//	@Failure		400	{string}	errMsg
//	@Failure		401	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/playlist/{id}/like [post]
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

// IsLike
//
//	@Summary		IsLike
//	@Description	Check if playlist is liked
//	@Tags			playlist
//	@Security		cookieAuth
//	@Produce		json
//	@Param			id	path		integer	true	"playlist id"
//	@Success		200	{object}	response.IsLiked
//	@Failure		400	{string}	errMsg
//	@Failure		401	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		404	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/playlist/{id}/is_like [get]
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

	if _, _, err = easyjson.MarshalToHTTPResponseWriter(response.IsLiked{IsLiked: isLiked}, w); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("response  formed")

	return nil
}

// Unlike
//
//	@Summary		Unlike
//	@Description	Unlike playlist
//	@Tags			playlist
//	@Security		cookieAuth
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Param			id	path	integer	true	"playlist id"
//	@Success		204
//	@Failure		400	{string}	errMsg
//	@Failure		401	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/playlist/{id}/unlike [delete]
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

// Delete
//
//	@Summary		Delete
//	@Description	Delete playlist
//	@Tags			playlist
//	@Security		cookieAuth
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Param			id	path	integer	true	"playlist id"
//	@Success		204
//	@Failure		400	{string}	errMsg
//	@Failure		401	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/playlist/{id} [delete]
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

// MakePublic
//
//	@Summary		MakePublic
//	@Description	Make playlist public
//	@Tags			playlist
//	@Security		cookieAuth
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Param			id	path	integer	true	"playlist id"
//	@Success		204
//	@Failure		400	{string}	errMsg
//	@Failure		401	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/playlist/{id}/make_public [put]
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

// MakePrivate
//
//	@Summary		MakePrivate
//	@Description	Make playlist private
//	@Tags			playlist
//	@Security		cookieAuth
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Param			id	path	integer	true	"playlist id"
//	@Success		204
//	@Failure		400	{string}	errMsg
//	@Failure		401	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/playlist/{id}/make_private [put]
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

// CollectionPlaylist
//
//	@Summary		CollectionPlaylist
//	@Description	Return user's playlist collection
//	@Tags			playlist
//	@Produce		json
//	@Security		cookieAuth
//	@Success		200	{array}		playlist.Base
//	@Failure		401	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/collection/playlists [get]
func (handler *Handler) CollectionPlaylist(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("Collection Playlists Handler entered")

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

	result, err := handler.playlistUseCase.CollectionPlaylists(userId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	if _, _, err = easyjson.MarshalToHTTPResponseWriter(playlist.Playlists{Playlists: result}, w); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("response  formed")

	return nil
}

func (handler *Handler) UserPlaylists(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("User Playlists Handler entered")

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

	result, err := handler.playlistUseCase.GetUserPlaylists(userId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	if _, _, err = easyjson.MarshalToHTTPResponseWriter(playlist.Playlists{Playlists: result}, w); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("response  formed")

	return nil
}
