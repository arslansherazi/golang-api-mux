package router

import (
	signup_api "find_competitor/apis/signup"
	middlewares "find_competitor/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func RouterV1() *mux.Router {
	baseRouter := mux.NewRouter()
	router := baseRouter.PathPrefix("/fnd/comp").Subrouter()

	// signup api
	signupHandlerFunc := http.HandlerFunc(signup_api.Signup)
	signupMiddlewaresChain := middlewares.New(middlewares.BasicAuthMiddleware).Then(signupHandlerFunc)
	router.Methods("POST").Path("/signup").HandlerFunc(signupMiddlewaresChain)

	return router
}
