package middlewares

import (
	"log"
	"net/http"
)

// SetMiddlewareLogger returns a logger that outputs & logs the method, host, request URI and protocol of the hit endpoint
// e.g. 2020/07/12 22:16:40 \n GET localhost:9000/users HTTP/1.1
// e.g. 2020/07/12 22:18:24 \n POST localhost:9000/users HTTP/1.1
func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		next(w, r)
	}
}

// SetMiddlewareJSON sets Content-Type header to "application/json"
// for example, execute POST request to create user, check Content-Type header
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// SetMiddlewareAuthentication extracts the token from the request body and checks if it is valid
// If no error,
// func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		err := auth.TokenService.ValidateToken(r)

// 		if err != nil {
// 			responses.ERROR(w, http.StatusUnauthorized, err)
// 			return
// 		}

// 		next(w, r)
// 	}
// }
