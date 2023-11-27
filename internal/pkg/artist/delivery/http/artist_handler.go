package artist_delivery

import (
	"github.com/sirupsen/logrus"
	"main/internal/common/handler"
	"main/internal/common/response"
	"main/internal/common/utils"
	"main/internal/pkg/artist"
	"main/internal/pkg/session"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ArtistHandler struct {
	ArtistUseCase  artist.UseCase
	SessionUseCase session.UseCase
	logger         *logrus.Logger
}

func NewHandler(su session.UseCase, artistUseCase artist.UseCase, logger *logrus.Logger) ArtistHandler {
	return ArtistHandler{
		ArtistUseCase:  artistUseCase,
		SessionUseCase: su,
		logger:         logger,
	}
}

// ArtistInfo
//
//	@Summary		ArtistInfo
//	@Description	Return artist info
//	@Tags			artist
//	@Produce		json
//	@Param			id	path		integer	true	"artist id"
//	@Success		200	{object}	artist.Response
//	@Failure		400	{string}	errMsg
//	@Failure		404	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/artist/{id} [get]
func (handler *ArtistHandler) ArtistInfo(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("ArtistInfo Handler entered")

	artistId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed artistId from Vars")

	artistInfo, err := handler.ArtistUseCase.GetArtistInfo(uint64(artistId))
	if err != nil {
		return common_handler.StatusError{Code: http.StatusNotFound, Err: err}
	}
	handler.logger.Infoln("Got artist from use case")

	if err = response.RenderJSON(w, artistInfo); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("response  formed")

	return nil
}

// Like
//
//	@Summary		Like
//	@Description	Like artist
//	@Tags			artist
//	@Security		cookieAuth
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Param			id	path	integer	true	"artist id"
//	@Success		204
//	@Failure		400	{string}	errMsg
//	@Failure		401	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		404	{string}	errMsg
//	@Router			/artist/{id}/like [post]
func (handler *ArtistHandler) Like(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("Like Artist Handler entered")

	artistId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed artistId from Vars")

	sessionId, err := response.GetCookie(r)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	handler.logger.Infoln("got cookie")

	userId, err := handler.SessionUseCase.GetUserId(sessionId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	handler.logger.Infoln("got user id by session id")

	if err = handler.ArtistUseCase.Like(userId, uint64(artistId)); err != nil {
		return common_handler.StatusError{Code: http.StatusNotFound, Err: err}
	}
	handler.logger.Infoln("Like created")

	w.WriteHeader(http.StatusNoContent)
	return nil
}

// IsLike
//
//	@Summary		IsLike
//	@Description	Check if artist is liked
//	@Tags			artist
//	@Security		cookieAuth
//	@Produce		json
//	@Param			id	path		integer	true	"artist id"
//	@Success		200	{object}	response.IsLiked
//	@Failure		400	{string}	errMsg
//	@Failure		401	{string}	errMsg
//	@Failure		404	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/artist/{id}/is_like [get]
func (handler *ArtistHandler) IsLike(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("Like Artist Handler entered")

	artistId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed artistId from Vars")

	sessionId, err := response.GetCookie(r)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	handler.logger.Infoln("got cookie")

	userId, err := handler.SessionUseCase.GetUserId(sessionId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	handler.logger.Infoln("got user id by session id")

	isLiked, err := handler.ArtistUseCase.IsLike(userId, uint64(artistId))
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

// Unlike
//
//	@Summary		Unlike
//	@Description	Unlike artist
//	@Tags			artist
//	@Security		cookieAuth
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Param			id	path	integer	true	"artist id"
//	@Success		204
//	@Failure		400	{string}	errMsg
//	@Failure		401	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		404	{string}	errMsg
//	@Router			/artist/{id}/unlike [delete]
func (handler *ArtistHandler) Unlike(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("Like Artist Handler entered")

	artistId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed artistId from Vars")

	sessionId, err := response.GetCookie(r)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	handler.logger.Infoln("got cookie")

	userId, err := handler.SessionUseCase.GetUserId(sessionId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	handler.logger.Infoln("got user id by session id")

	if err = handler.ArtistUseCase.Unlike(userId, uint64(artistId)); err != nil {
		return common_handler.StatusError{Code: http.StatusNotFound, Err: err}
	}
	handler.logger.Infoln("Like deleted")

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (handler *ArtistHandler) FullSearch(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("Artist delivery Full Search entered")

	params := r.URL.Query()
	query := params.Get("query")
	handler.logger.Infoln("Got query from uri", query)

	result, err := handler.ArtistUseCase.FullSearch(query)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusNotFound, Err: err}
	}
	handler.logger.Infoln("got response from useCase")

	if err = response.RenderJSON(w, result); err != nil {
		return common_handler.StatusError{Code: http.StatusNotFound, Err: err}
	}

	return nil
}
