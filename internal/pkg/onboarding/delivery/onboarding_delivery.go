package onboarding_delivery

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	common_handler "main/internal/common/handler"
	"main/internal/common/response"
	"main/internal/pkg/onboarding"
	"main/internal/pkg/session"
	"net/http"
)

type Handler struct {
	onboardingUseCase onboarding.UseCase
	sessionUseCase    session.UseCase
	logger            *logrus.Logger
}

func NewHandler(su session.UseCase, ou onboarding.UseCase, l *logrus.Logger) Handler {
	return Handler{
		onboardingUseCase: ou,
		sessionUseCase:    su,
		logger:            l,
	}
}

func (h *Handler) GetGenres(w http.ResponseWriter, r *http.Request) error {
	h.logger.Infoln("Onboarding delivery Get Genres entered")

	result, err := h.onboardingUseCase.GetGenres()
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	h.logger.Infoln("Got Genres")

	if err = response.RenderJSON(w, result); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	h.logger.Infoln("Genres formed")

	return nil
}

func (h *Handler) GetArtists(w http.ResponseWriter, _ *http.Request) error {
	h.logger.Infoln("Onboarding delivery Get Artists entered")

	result, err := h.onboardingUseCase.GetArtists()
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	h.logger.Infoln("Got Artists")

	if err = response.RenderJSON(w, result); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	h.logger.Infoln("Artists formed")

	return nil
}

func (h *Handler) SaveArtists(w http.ResponseWriter, r *http.Request) error {
	h.logger.Infoln("Onboarding delivery Save Artists entered")

	sessionId, err := response.GetCookie(r)
	if err != nil {
		h.logger.Errorln("error get from cookie", sessionId, err)
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	h.logger.Infoln("Got Cookie")

	userId, err := h.sessionUseCase.GetUserId(sessionId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	h.logger.Infoln("Got user id")

	var artists onboarding.ArtistOnboarding
	if err := json.NewDecoder(r.Body).Decode(&artists); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	h.logger.Infoln("artists decoded")

	artists.UserId = userId
	if err = h.onboardingUseCase.SetUserArtists(artists); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	h.logger.Infoln("artists saved for user")

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (h *Handler) SaveGenres(w http.ResponseWriter, r *http.Request) error {
	h.logger.Infoln("Onboarding delivery Save Genres entered")

	sessionId, err := response.GetCookie(r)
	if err != nil {
		h.logger.Errorln("error get from cookie", sessionId, err)
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	h.logger.Infoln("Got Cookie")

	userId, err := h.sessionUseCase.GetUserId(sessionId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	h.logger.Infoln("Got user id")

	var genres onboarding.GenreOnboarding
	if err := json.NewDecoder(r.Body).Decode(&genres); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	h.logger.Infoln("artists decoded")

	genres.UserId = userId
	if err = h.onboardingUseCase.SetUserGenres(genres); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	h.logger.Infoln("artists saved for user")

	w.WriteHeader(http.StatusNoContent)
	return nil
}
