package routes

import "net/http"

type Route struct {
	Uri          string
	Method       string
	Handler      func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

func Load() []Route {
	routes := productsRoutes
	return routes
}
