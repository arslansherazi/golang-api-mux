package middlewares

import (
	"fmt"
	"net/http"
)

func BasicAuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("422 - Basic Auth Failed"))
		return
		// handler.ServeHTTP(w, r)
	})
}

func JwtTokenMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("JWT Token Middleware")
		handler.ServeHTTP(w, r)
	})
}

type Middleware func(http.Handler) http.Handler
type Chain []Middleware

func New(middlewares ...Middleware) Chain {
	var slice Chain
	return append(slice, middlewares...)
}

func (c Chain) Then(originalHandler http.Handler) http.Handler {
	if originalHandler == nil {
		originalHandler = http.DefaultServeMux
	}

	for i := range c {
		// Equivalent to middleware1(middleware2(middleware3(originalHandler)))
		originalHandler = c[len(c)-1-i](originalHandler)
	}
	return originalHandler
}
