package server

import (
	"github.com/arpit006/go_txns/internal/constants"
	"github.com/arpit006/go_txns/internal/services"
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter will register all the http-api routes to our http-router
func NewRouter() http.Handler  {
	// creating a new instance of the Mux Router
	r := mux.NewRouter()
	// creating a sub-router with prefix "/txns"
	// implicitly all the paths that we register with this router will have a prefix /txns in them
	router := r.PathPrefix(constants.TXN_PREFIX).Subrouter()

	// register Health Router
	router.HandleFunc(constants.HEALTH, services.HealthCheckHandler)

	return http.HandlerFunc(router.ServeHTTP)
}
