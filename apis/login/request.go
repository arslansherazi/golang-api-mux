package login_api

import (
	"find_competitor/common"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// response header
	w.Header().Set("Content-Type", "application/json")

	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal(common.ENVIRONMENT_VARIBALES_ERROR_MESSAGE)
		common.ErrorResponse(r.URL.Path, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
	}

	logger, err := common.GetLogger("login_api")
	if err != nil {
		log.Fatal(common.LOGGER_ERROR_MESSAGE)
		common.ErrorResponse(r.URL.Path, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
	} else {

	}
}
