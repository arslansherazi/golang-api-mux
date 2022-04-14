package router

import (
	apis_v1 "find_competitor/apis/v1"

	"github.com/gorilla/mux"
)

func RouterV1() *mux.Router {
	baseRouter := mux.NewRouter()
	router := baseRouter.PathPrefix("/fnd/comp/v1").Subrouter()

	router.Methods("POST").Path("/login").HandlerFunc(apis_v1.Login)

	return router
}
