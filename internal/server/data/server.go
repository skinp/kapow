package data

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type routeSpec struct {
	route  string
	method string
	rh     resourceHandler
}

func configRouter(rs []routeSpec) (r *mux.Router) {
	r = mux.NewRouter()
	for _, s := range rs {
		r.HandleFunc(s.route, checkHandler(s.rh)).Methods(s.method)
	}
	r.HandleFunc(
		"/handlers/{handlerID}/{resource:.*}",
		func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusBadRequest) })
	return r
}

func Run(bindAddr string) {
	rs := []routeSpec{
		// request
		{"/handlers/{handlerID}/request/method", "GET", getRequestMethod},
		{"/handlers/{handlerID}/request/host", "GET", getRequestHost},
		{"/handlers/{handlerID}/request/path", "GET", getRequestPath},
		{"/handlers/{handlerID}/request/matches/{name}", "GET", getRequestMatches},
		{"/handlers/{handlerID}/request/params/{name}", "GET", getRequestParams},
		{"/handlers/{handlerID}/request/headers/{name}", "GET", getRequestHeaders},
		{"/handlers/{handlerID}/request/cookies/{name}", "GET", getRequestCookies},
		{"/handlers/{handlerID}/request/form/{name}", "GET", getRequestForm},
		{"/handlers/{handlerID}/request/files/{name}/filename", "GET", getRequestFileName},
		{"/handlers/{handlerID}/request/files/{name}/content", "GET", getRequestFileContent},
		{"/handlers/{handlerID}/request/body", "GET", getRequestBody},

		// response
		{"/handlers/{handlerID}/response/status", "PUT", lockResponseWriter(setResponseStatus)},
		{"/handlers/{handlerID}/response/headers/{name}", "PUT", lockResponseWriter(setResponseHeaders)},
		{"/handlers/{handlerID}/response/cookies/{name}", "PUT", lockResponseWriter(setResponseCookies)},
		{"/handlers/{handlerID}/response/body", "PUT", lockResponseWriter(setResponseBody)},
		{"/handlers/{handlerID}/response/stream", "PUT", lockResponseWriter(setResponseBody)},
	}
	log.Fatal(http.ListenAndServe(bindAddr, configRouter(rs)))
}
