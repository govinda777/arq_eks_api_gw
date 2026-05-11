package router

import (
	"github.com/marcosouzatech/items-api/api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate returns a router com as routes configured
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
