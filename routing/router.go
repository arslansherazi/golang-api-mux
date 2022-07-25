package router

import (
	edit_competition_api "find_competitor/apis/edit_competition"
	host_competition_api "find_competitor/apis/host_competition"
	login_api "find_competitor/apis/login"
	signup_api "find_competitor/apis/signup"
	validate_phone_number_api "find_competitor/apis/validate_phone_number"
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
	router := baseRouter.PathPrefix("/fnd/comp/").Subrouter()

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

	// validate phone number api
	validatePhoneNumberHandlerFunc := http.HandlerFunc(validate_phone_number_api.ValidatePhoneNumber)
	validatePhoneNumberMiddlewaresChain := middlewares.New(middlewares.BasicAuthMiddleware).Then(validatePhoneNumberHandlerFunc)
	router.Methods("POST").Path("/validate/phone/number").HandlerFunc(validatePhoneNumberMiddlewaresChain)

	// host competition api
	hostCompetitionHandlerFunc := http.HandlerFunc(host_competition_api.HostCompetition)
	hostCompetitionMiddlewaresChain := middlewares.New(middlewares.JwtTokenMiddleware).Then(hostCompetitionHandlerFunc)
	router.Methods("POST").Path("/host/competition").HandlerFunc(hostCompetitionMiddlewaresChain)

	// host competition api
	editCompetitionHandlerFunc := http.HandlerFunc(edit_competition_api.EditCompetition)
	editCompetitionMiddlewaresChain := middlewares.New(middlewares.JwtTokenMiddleware).Then(editCompetitionHandlerFunc)
	router.Methods("POST").Path("/edit/competition").HandlerFunc(editCompetitionMiddlewaresChain)

	return router
}
