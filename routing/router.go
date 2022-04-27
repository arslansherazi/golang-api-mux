package router

import (
	signup_api "find_competitor/apis/signup"

	"github.com/gorilla/mux"
)

func RouterV1() *mux.Router {
	baseRouter := mux.NewRouter()
	router := baseRouter.PathPrefix("/fnd/comp").Subrouter()

	router.Methods("POST").Path("/signup").HandlerFunc(signup_api.Signup)

	return router
}
