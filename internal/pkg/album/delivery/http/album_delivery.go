package album_delivery

import (
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
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

type AlbumHandler struct {
	trackUseCase   track.UseCase
	albumUseCase   album.UseCase
	sessionUseCase session.UseCase
	logger         *logrus.Logger
}

func NewHandler(trackUseCase track.UseCase, albumUseCase album.UseCase, session session.UseCase, logger *logrus.Logger) AlbumHandler {
	return AlbumHandler{
		trackUseCase:   trackUseCase,
		albumUseCase:   albumUseCase,
		sessionUseCase: session,
		logger:         logger,
	}
}

// Feed
//
//	@Summary		Feed
//	@Description	Return all albums
//	@Tags			album
//	@Produce		json
//	@Success		200	{array}		album.Response
//	@Failure		500	{string}	errMsg
//	@Router			/feed [get]

func (handler *AlbumHandler) Feed(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("Feed Handler entered")

	albums, err := handler.albumUseCase.GetRandom()
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("got random albums")

	return handler.handleQuery(albums, w, r)
}

// New
//
//	@Summary		New
//	@Description	Return new albums
//	@Tags			album
//	@Produce		json
//	@Success		200	{array}		album.Response
//	@Failure		500	{string}	errMsg
//	@Router			/new [get]
func (handler *AlbumHandler) New(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("New Handler entered")

	albums, err := handler.albumUseCase.GetNew()
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("got new albums")

	return handler.handleQuery(albums, w, r)
}

// MostLiked
//
//	@Summary		MostLiked
//	@Description	Return albums sorted by likes count
//	@Tags			album
//	@Produce		json
//	@Success		200	{array}		album.Response
//	@Failure		500	{string}	errMsg
//	@Router			/most_liked [get]
func (handler *AlbumHandler) MostLiked(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("MostLiked Handler entered")

	albums, err := handler.albumUseCase.GetMostLiked()
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("got most liked albums")

	return handler.handleQuery(albums, w, r)
}

// Popular
//
//	@Summary		Popular
//	@Description	Return albums sorted by listen count
//	@Tags			album
//	@Produce		json
//	@Success		200	{array}		album.Response
//	@Failure		500	{string}	errMsg
//	@Router			/popular [get]
func (handler *AlbumHandler) Popular(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("Popular Handler entered")

	albums, err := handler.albumUseCase.GetPopular()
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("got popular albums")

	return handler.handleQuery(albums, w, r)
}

// AlbumTracks
//
//	@Summary		AlbumTracks
//	@Description	Return album info with all tracks
//	@Tags			album
//	@Produce		json
//	@Param			id	path		integer	true	"album id"
//	@Success		200	{object}	album.Response
//	@Failure		400	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/album/{id} [get]
func (handler *AlbumHandler) AlbumTracks(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("Album Handler entered")

	albumId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("got id from query var")

	result, err := handler.albumUseCase.GetAlbum(uint64(albumId))
	if err != nil {
		return common_handler.StatusError{Code: http.StatusNotFound, Err: err}
	}
	handler.logger.Infoln("got album by id")

	sessionId, _ := response.GetCookie(r)
	userId, err := handler.sessionUseCase.GetUserId(sessionId)
	if err != nil {
		userId = ""
	}

	if userId != "" {
		labeledTracks, err := handler.trackUseCase.LabelIsLikedTracks(userId, result.Tracks)
		if err != nil {
			handler.logger.Errorln("Error Labeling Tracks with IsLiked", err)
		}
		result.Tracks = labeledTracks
	}

	if _, _, err = easyjson.MarshalToHTTPResponseWriter(result, w); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("formed response")

	if userId == "" {
		w.WriteHeader(http.StatusUnauthorized)
	}

	return nil
}

// AlbumWithRequiredTrack
//
//	@Summary		AlbumWithRequiredTrack
//	@Description	Return album that contain required track
//	@Tags			album
//	@Produce		json
//	@Param			id	path		integer	true	"track id"
//	@Success		200	{object}	album.Response
//	@Failure		400	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/track/{id} [get]
func (handler *AlbumHandler) AlbumWithRequiredTrack(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("AlbumWithRequiredTrack Handler entered")

	trackId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("got id from query var")

	result, err := handler.albumUseCase.GetAlbumByTrack(uint64(trackId))
	if err != nil {
		return common_handler.StatusError{Code: http.StatusNotFound, Err: err}
	}
	handler.logger.Infoln("got album with required track by track id")

	sessionId, _ := response.GetCookie(r)
	userId, _ := handler.sessionUseCase.GetUserId(sessionId)

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
	handler.logger.Infoln("formed response")

	if userId == "" {
		w.WriteHeader(http.StatusUnauthorized)
	}

	return nil
}

func (handler *AlbumHandler) handleQuery(albums []album.Response, w http.ResponseWriter, _ *http.Request) error {
	handler.logger.Infoln("handle query entered")

	if _, _, err := easyjson.MarshalToHTTPResponseWriter(album.Albums{Albums: albums}, w); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("formed response")

	return nil
}

// Like
//
//	@Summary		Like
//	@Description	Like album
//	@Tags			album
//	@Security		cookieAuth
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Param			id	path	integer	true	"album id"
//	@Success		204
//	@Failure		400	{string}	errMsg
//	@Failure		401	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/album/{id}/like [post]
func (handler *AlbumHandler) Like(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("Like Handler entered")

	albumId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed albumId from Vars")

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

	if err = handler.albumUseCase.Like(userId, uint64(albumId)); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("like created successfully")

	w.WriteHeader(http.StatusNoContent)
	return nil
}

// IsLike
//
//	@Summary		IsLike
//	@Description	Check if album is liked
//	@Tags			album
//	@Security		cookieAuth
//	@Produce		json
//	@Param			id	path		integer	true	"album id"
//	@Success		200	{object}	response.IsLiked
//	@Failure		400	{string}	errMsg
//	@Failure		401	{string}	errMsg
//	@Failure		404	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/album/{id}/is_like [get]
func (handler *AlbumHandler) IsLike(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("Like Handler entered")

	albumId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed albumId from Vars")

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

	isLiked, err := handler.albumUseCase.IsLike(userId, uint64(albumId))
	if err != nil {
		return common_handler.StatusError{Code: http.StatusNotFound, Err: err}
	}
	handler.logger.Infoln("artist like checked")

	if _, _, err = easyjson.MarshalToHTTPResponseWriter(response.IsLiked{IsLiked: isLiked}, w); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("response  formed")

	return nil
}

// Unlike
//
//	@Summary		Unlike
//	@Description	Unlike album
//	@Tags			album
//	@Security		cookieAuth
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Param			id	path	integer	true	"album id"
//	@Success		204
//	@Failure		400	{string}	errMsg
//	@Failure		401	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/album/{id}/unlike [delete]
func (handler *AlbumHandler) Unlike(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("Like Handler entered")

	albumId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed albumId from Vars")

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

	if err = handler.albumUseCase.Unlike(userId, uint64(albumId)); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("like deleted successfully")

	w.WriteHeader(http.StatusNoContent)
	return nil
}

// CollectionAlbum
//
//	@Summary		CollectionAlbum
//	@Description	Return user's album collection
//	@Tags			album
//	@Produce		json
//	@Security		cookieAuth
//	@Success		200	{object}	album.LikedAlbums
//	@Failure		401	{string}	errMsg
//	@Failure		404	{string}	errMsg
//	@Router			/collection/albums [get]
func (handler *AlbumHandler) CollectionAlbum(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("Artist delivery Collection Artist entered")

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

	result, err := handler.albumUseCase.GetUserAlbums(userId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}

	if _, _, err = easyjson.MarshalToHTTPResponseWriter(result, w); err != nil {
		return common_handler.StatusError{Code: http.StatusNotFound, Err: err}
	}
	return nil
}
