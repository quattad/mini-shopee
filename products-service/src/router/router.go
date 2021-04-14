package router

import (
	"github.com/gorilla/mux"
)

// Registers path and returns a new router instance
func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return r // testing
	// return routes.SetupRoutesWithMiddlewares(r)
}
