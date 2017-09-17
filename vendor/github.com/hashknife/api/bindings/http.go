package bindings

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashknife/api/config"
	"github.com/hashknife/api/endpoints"
	"github.com/hashknife/common/middleware"
)

const (
	component = "hashknife-api"
	prefix    = "/" + component
	version   = "/v1"
)

// HTTPListenerParams holds parameters for building routers
type HTTPListenerParams struct {
	Logger  kitlog.Logger
	Root    context.Context
	ErrChan chan error
	Config  *config.Config
}

// StartApplicationHTTPListener creates a goroutine that has an HTTP listener for the application endpoints
func StartApplicationHTTPListener(hlp *HTTPListenerParams) {
	go func() {
		ctx, cancel := context.WithCancel(hlp.Root)
		defer cancel()
		hlp.Logger.Log("HTTPAddress", *hlp.Config.HTTPAddress, "transport", "HTTP/JSON")
		e := &endpoints.Endpoints{
			Package:  endpoints.NewPackageService(hlp.Config),
			Profile:  endpoints.NewProfileService(hlp.Config),
			User:     endpoints.NewUserService(hlp.Config),
			Frontend: endpoints.NewFrontend(hlp.Config),
		}
		router := createApplicationRouter(ctx, hlp.Logger, hlp.Config, e)
		//listenerMetrics := metrics.NewSimpleMetricsMiddleware(component, "app_http_listener", *hlp.Config.StatsdReportingIntervalSeconds, *hlp.Config.StatsdAddress)
		//hlp.ErrChan <- http.ListenAndServe(*hlp.Config.HTTPAddress, handlers.RecoveryHandler()(handlers.CombinedLoggingHandler(kitlog.NewStdlibAdapter(hlp.Logger), listenerMetrics.Annotate(router))))
		hlp.ErrChan <- http.ListenAndServe(*hlp.Config.HTTPAddress, handlers.RecoveryHandler()(handlers.CombinedLoggingHandler(kitlog.NewStdlibAdapter(hlp.Logger), router)))
	}()
}

// createApplicationRouter sets up the router that will handle all of the application routes
func createApplicationRouter(ctx context.Context, l kitlog.Logger, conf *config.Config, e *endpoints.Endpoints) *mux.Router {
	router := mux.NewRouter()
	router.Handle("/frontend", e.Frontend)
	apiRouter := router.PathPrefix(prefix + version).Subrouter()
	auth := middleware.NewHashknifeRequestAuthenticator(conf.HashknifeAuthToken)
	apiRouter.Handle(
		"/package/status/{account_id}/{package_id}",
		kithttp.NewServer(
			endpoint.Chain(auth.EndpointAuthenticate())(e.Package.Accept),
			decodePackageStatusHTTPRequest,
			encodeResponse,
			kithttp.ServerBefore(middleware.KitServerBefore),
		)).Methods(http.MethodGet)
	apiRouter.Handle(
		"/package/deliver",
		kithttp.NewServer(
			endpoint.Chain(auth.EndpointAuthenticate())(e.Package.Deliver),
			decodePackageDeliverHTTPRequest,
			encodeResponse,
			kithttp.ServerBefore(middleware.KitServerBefore),
		)).Methods(http.MethodPost)
	apiRouter.Handle(
		"/package/accept",
		kithttp.NewServer(
			endpoint.Chain(auth.EndpointAuthenticate())(e.Package.Accept),
			decodePackageAcceptHTTPRequest,
			encodeResponse,
			kithttp.ServerBefore(middleware.KitServerBefore),
		)).Methods(http.MethodPost)

	apiRouter.Handle(
		"/user/{account_id}/{user_id}",
		kithttp.NewServer(
			endpoint.Chain(auth.EndpointAuthenticate())(e.User.Retrieve),
			decodeRetrieveUserHTTPRequest,
			encodeResponse,
			kithttp.ServerBefore(middleware.KitServerBefore),
		)).Methods(http.MethodGet)
	apiRouter.Handle(
		"/user/{account_id}/{user_id}",
		kithttp.NewServer(
			endpoint.Chain(auth.EndpointAuthenticate())(e.User.Update),
			decodeUpdateUserHTTPRequest,
			encodeResponse,
			kithttp.ServerBefore(middleware.KitServerBefore),
		)).Methods(http.MethodPut)
	apiRouter.Handle(
		"/user/{account_id}/{user_id}",
		kithttp.NewServer(
			endpoint.Chain(auth.EndpointAuthenticate())(e.User.Disable),
			decodeCreateDisableHTTPRequest,
			encodeResponse,
			kithttp.ServerBefore(middleware.KitServerBefore),
		)).Methods(http.MethodPost)
	return router
}

// StartHealthCheckHTTPListener creates a goroutine that has an HTTP listener for the healthcheck endpoint
func StartHealthCheckHTTPListener(p *HTTPListenerParams, gs string) {
	go func() {
		ctx, cancel := context.WithCancel(p.Root)
		defer cancel()
		p.Logger.Log("HealthCheckAddress", *p.Config.HealthCheckAddress, "transport", "HTTP/JSON")
		router := createHealthCheckRouter(ctx, p.Logger, endpoints.NewHealthCheckEndpoint(gs))
		p.ErrChan <- http.ListenAndServe(*p.Config.HealthCheckAddress, handlers.RecoveryHandler()(handlers.CombinedLoggingHandler(kitlog.NewStdlibAdapter(p.Logger), router)))
	}()
}

// createHealthCheckRouter setups up the router that provides the health checking functionality
func createHealthCheckRouter(ctx context.Context, l kitlog.Logger, h endpoints.HealthCheckServicer) *mux.Router {
	router := mux.NewRouter().PathPrefix(prefix).Subrouter()
	router.Handle(
		"/healthcheck",
		kithttp.NewServer(
			h.Run,
			noOpDecodeRequest,
			encodeHealthCheckHTTPResponse,
		)).Methods(http.MethodGet)
	return router
}

// addPoweredHeaders adds a header indicating the
func addPoweredHeaders(w http.ResponseWriter) {
	w.Header().Add("X-Powered-By", "hashknife"+prefix)
}
