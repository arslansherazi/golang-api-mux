package signup_api

import (
	"find_competitor/common"
	"find_competitor/configs"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	// response header
	w.Header().Set("Content-Type", "application/json")

	err := godotenv.Load()
	if err != nil {
		log.Fatal(common.ENVIRONMENT_VARIBALES_ERROR_MESSAGE)
		common.ErrorResponse(r.URL.Path, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
	} else {
		logger, err := common.GetLogger("signup_api")
		if err != nil {
			log.Println(common.LOGGER_ERROR_MESSAGE)
			common.ErrorResponse(r.URL.Path, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
		} else {
			// request url
			requestUrl := r.URL.Path

			// get db instance
			isScript := true
			db, err := configs.GetDbInstance(isScript)

			if err != nil {
				log.Print(err)
				common.ErrorResponse(r.URL.Path, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
			} else {
				user, profileImage, err, isValidationError := processRequestParams(logger, r)
				if err != nil {
					if isValidationError {
						validationMessage := common.ParseValidationError(err)
						common.ErrorResponse(requestUrl, http.StatusUnprocessableEntity, validationMessage, w)
					} else {
						logger.Println(err)
						common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
					}
				} else {
					user.ProfileImageUrl, err = generateProfileImageUrl(profileImage)
					if err != nil {
						common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
					} else {
						user.Password, err = createHashOfPassword(user.Password)
						if err != nil {
							common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
						} else {
							insertUserIntoDB(db, user)
							generateSuccessResponse(requestUrl, w)
						}
					}
				}
			}
		}
	}
}
