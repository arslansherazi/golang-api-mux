package login_api

import (
	"find_competitor/common"
	"find_competitor/configs"
	"find_competitor/models"
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
	} else {
		// request url
		requestUrl := r.URL.Path

		logger, err := common.GetLogger("login_api")
		if err != nil {
			log.Fatal(common.LOGGER_ERROR_MESSAGE)
			common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
		} else {
			requestData, err, isValidationError := processRequestParams(r)
			if err != nil {
				if isValidationError {
					validationMessage := common.ParseValidationError(err)
					common.ErrorResponse(requestUrl, http.StatusUnprocessableEntity, validationMessage, w)
				} else {
					logger.Println(err)
					common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
				}
			} else {
				phoneNumber := requestData.PhoneNumber
				password := requestData.Password

				// get db instance
				isScript := true
				db, err := configs.GetDbInstance(isScript)

				if err != nil {
					logger.Println(err)
					common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
				} else {
					userData := validateUser(db, phoneNumber)
					if (models.User{}) == userData {
						common.ErrorResponse(requestUrl, http.StatusUnprocessableEntity, common.USER_NOT_EXIST_ERROR_MESSAGE, w)
					} else {
						isPasswordVerified := verifyPassword(password, userData.Password)
						if !isPasswordVerified {
							common.ErrorResponse(requestUrl, http.StatusUnprocessableEntity, common.INCORRECT_PASSWORD_ERROR_MESSAGE, w)
						} else {
							token, err := generateToken(phoneNumber)
							if err != nil {
								logger.Println(err)
								common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
							} else {
								generateSuccessResponse(requestUrl, int(userData.ID), userData.Name, userData.ProfileImageUrl, token, w)
							}
						}
					}
				}
			}
		}
	}
}
