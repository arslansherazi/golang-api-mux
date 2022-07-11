package middlewares

import (
	"find_competitor/common"
	"fmt"
	"net/http"
)

func BasicAuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, _ := r.BasicAuth()
		if user != common.BASIC_AUTH_USERNAME || pass != common.BASIC_AUTH_PASSWORD {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			common.ErrorResponse(r.URL.Path, http.StatusUnauthorized, common.UNAUTHORIZED_ACCESS_ERROR_MESSAGE, w)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func JwtTokenMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("JWT Token Middleware")
		handler.ServeHTTP(w, r)
	})
}

type Middleware func(http.HandlerFunc) http.HandlerFunc
type Chain []Middleware

func New(middlewares ...Middleware) Chain {
	var slice Chain
	return append(slice, middlewares...)
}

func (c Chain) Then(originalHandler http.HandlerFunc) http.HandlerFunc {
	for i := range c {
		// Equivalent to middleware1(middleware2(middleware3(originalHandler)))
		originalHandler = c[len(c)-1-i](originalHandler)
	}
	return originalHandler
}
