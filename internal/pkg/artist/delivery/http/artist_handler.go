package artist_delivery

import (
	"github.com/sirupsen/logrus"
	"main/internal/pkg/artist"
	common_handler "main/internal/pkg/common/handler"
	"main/internal/pkg/common/response"
	"main/internal/pkg/common/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ArtistHandler struct {
	ArtistUseCase artist.UseCase
	logger        *logrus.Logger
}

func NewHandler(artistUseCase artist.UseCase, logger *logrus.Logger) ArtistHandler {
	return ArtistHandler{
		ArtistUseCase: artistUseCase,
		logger:        logger,
	}
}

// ArtistInfo
//
//	@Description	return artist info
//	@Tags			artist
//	@Produce		json
//	@Param			id	path		integer	true	"artist id"
//	@Success		200	{object}	artist.Response
//	@Failure		400	{string}	errMsg
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

	err = response.RenderJSON(w, artistInfo)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("response  formed")

	return nil
}
