package router_init

import (
	"github.com/gorilla/mux"
	common_handler "main/internal/pkg/common/handler"
	track_delivery "main/internal/pkg/track/delivery/http"
	user_delivery "main/internal/pkg/user/delivery/http"
)

//	@title			MusicOn API
//	@version		1.0
//	@description	Music web app

//	@termsOfService	http://swagger.io/terms/

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@SecurityDefinitions.apikey	cookieAuth 
//	@in							header
//	@name						JSESSIONID

//	@SecurityDefinitions.apikey	csrfToken 
//	@in							header
//	@name						X-CSRFTOKEN

//	@host		musicon.space
//	@BasePath	/api/v1
func New(router *mux.Router, userHandler user_delivery.UserHandler, trackHandler track_delivery.TrackHandler) *mux.Router {
	router.Handle("/api/v1/sign_up", common_handler.Handler{H: userHandler.SignUp}).Methods("POST")
	router.Handle("/api/v1/login", common_handler.Handler{H: userHandler.Login}).Methods("POST")
	router.Handle("/api/v1/auth", common_handler.Handler{H: userHandler.Auth}).Methods("GET")
	router.Handle("/api/v1/logout", common_handler.Handler{H: userHandler.LogOut}).Methods("POST")

	router.Handle("/api/v1/music", common_handler.Handler{H: trackHandler.Music}).Methods("GET")

	return router
}
