package router

import (
	login_api "find_competitor/apis/login"
	signup_api "find_competitor/apis/signup"
	"find_competitor/common"
	middlewares "find_competitor/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	common.ErrorResponse(r.URL.Path, http.StatusNotFound, "404 Not Found", w)
	return
}

func RouterV1() *mux.Router {
	baseRouter := mux.NewRouter()
	router := baseRouter.PathPrefix("/fnd/comp").Subrouter()

	// handle 404
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	// signup api
	signupHandlerFunc := http.HandlerFunc(signup_api.Signup)
	signupMiddlewaresChain := middlewares.New(middlewares.BasicAuthMiddleware).Then(signupHandlerFunc)
	router.Methods("POST").Path("/signup").HandlerFunc(signupMiddlewaresChain)

	// login api
	loginHandlerFunc := http.HandlerFunc(login_api.Login)
	loginMiddlewaresChain := middlewares.New(middlewares.BasicAuthMiddleware).Then(loginHandlerFunc)
	router.Methods("POST").Path("/login").HandlerFunc(loginMiddlewaresChain)

	return router
}
