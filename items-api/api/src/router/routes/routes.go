package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all routes da API
type Route struct {
	URI                string
	Method             string
	Function             func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := routesItems

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
