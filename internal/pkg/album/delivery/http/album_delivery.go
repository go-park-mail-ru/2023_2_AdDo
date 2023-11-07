package album_delivery

import (
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
//	@Description	return all albums
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
//	@Description	return new albums
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
//	@Description	return albums sorted by likes count
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
//	@Description	return albums sorted by listen count
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
//	@Description	return album info with all tracks
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
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("got album by id")

	if err = response.RenderJSON(w, result); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("formed response")

	return nil
}

func (handler *AlbumHandler) handleQuery(albums []album.Response, w http.ResponseWriter, r *http.Request) error {
	handler.logger.Infoln("handle query entered")

	if err := response.RenderJSON(w, albums); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("formed response")

	return nil
}

// Like
//
//	@Description	like album
//	@Tags			track
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Accept			json
//	@Param			trackId	body		integer	true	"track id"
//	@Success		204
//	@Failure		400		{string}	errMsg
//	@Failure		401		{string}	errMsg
//	@Failure		403		{string}	errMsg
//	@Failure		500		{string}	errMsg
//	@Router			/like [post]

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

	if err = response.RenderJSON(w, response.IsLiked{IsLiked: isLiked}); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("response  formed")

	return nil
}

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
