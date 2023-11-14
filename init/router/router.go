package router_init

import (
	_ "main/api/openapi"
	common_handler "main/internal/common/handler"
	"main/internal/common/metrics"
	"main/internal/common/middleware/common"
	prom "main/internal/common/middleware/prometheus"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

//	@title			MusicOn API
//	@version		1.0
//	@description	Music web app

//	@termsOfService	http://swagger.io/terms/

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@SecurityDefinitions.apikey	cookieAuth
//	@in							cookie
//	@name						JSESSIONID

//	@SecurityDefinitions.apikey	csrfToken
//	@in							header
//	@name						X-Csrf-Token

//	@SecurityDefinitions.apikey	cookieCsrfToken
//	@in							cookie
//	@name						X-Csrf-Token

// @host		musicon.space
// @BasePath	/api/v1
const MethodPost = "POST"
const MethodGet = "GET"
const MethodPut = "PUT"
const MethodDelete = "DELETE"

type Route struct {
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request) error
	Path    string
}

func NewRoute(p string, h func(w http.ResponseWriter, r *http.Request) error, m string) Route {
	return Route{
		Method:  m,
		Handler: h,
		Path:    p,
	}
}

type Config struct {
	Routes           []Route
	Prefix           string
	Middlewares      []mux.MiddlewareFunc
	SubRouterConfigs []Config
}

func New(config Config, logger *logrus.Logger) http.Handler {
	router := mux.NewRouter()
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	reg := prometheus.NewRegistry()
	metrics := metrics.New(reg)
	handlersMap := prom.NewHandlersMap()
	// router.PathPrefix("/metrics").Handler(promhttp.Handler())
	router.PathPrefix("/metrics").Handler(promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))

	for _, route := range config.Routes {
		router.Handle(config.Prefix+route.Path, common_handler.Handler{H: route.Handler}).Methods(route.Method)
	}

	for _, subConfig := range config.SubRouterConfigs {
		subRouter := router.PathPrefix("").Subrouter()
		for _, route := range subConfig.Routes {
			subRouter.Handle(config.Prefix+subConfig.Prefix+route.Path, common_handler.Handler{H: route.Handler}).Methods(route.Method)
		}

		subRouter.Use(subConfig.Middlewares...)
	}

	router.Use(config.Middlewares...)

	routerWithMiddleware := common.Logging(router, logger)
	routerWithMiddleware = prom.CollectMetrics(routerWithMiddleware, metrics, handlersMap)
	routerWithMiddleware = common.PanicRecovery(routerWithMiddleware, logger)

	return routerWithMiddleware
}

/// TODO написать ручки на все коллекции: треки, артисты, плейлисты и альбомы, все коллекции возвращают имя владельца и его фото, тесты на них
/// TODO полнотекстовый поиск по артистам, трекам и альбомам среди имен, поиск работает в реальном времени
/// TODO easyjson для сериализации и десериализации
