package router

import (
	"github/quattad/mini-shopee/products-service/src/api/router/routes"

	"github.com/gorilla/mux"
)

// Registers path and returns a new router instance
func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutesWithMiddlewares(r)
}
