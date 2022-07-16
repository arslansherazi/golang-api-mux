package signup_api

import (
	"find_competitor/common"
	"log"
	"net/http"
	"sync"

	"github.com/joho/godotenv"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	// response header
	w.Header().Set("Content-Type", "application/json")

	wg := new(sync.WaitGroup)

	err := godotenv.Load()
	if err != nil {
		log.Fatal(common.ENVIRONMENT_VARIBALES_ERROR_MESSAGE)
		common.ErrorResponse(r.URL.Path, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
	}

	logger, err := common.GetLogger("signup_api")
	if err != nil {
		log.Fatal(common.LOGGER_ERROR_MESSAGE)
		common.ErrorResponse(r.URL.Path, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
	} else {
		user, profileImage, err, isValidationError := processRequestParams(logger, r)
		if err != nil {
			if isValidationError {
				validationMessage := common.ParseValidationError(err)
				common.ErrorResponse(r.URL.Path, http.StatusUnprocessableEntity, validationMessage, w)
			} else {
				logger.Println(err)
				common.ErrorResponse(r.URL.Path, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
			}
		} else {
			// wg.Add(1)

			profileImageUrl, fileName := generateProfileImageUrl()
			user.ProfileImageUrl = profileImageUrl
			go uploadImage(profileImage, fileName, wg)
			if err != nil {
				common.ErrorResponse(r.URL.Path, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
			} else {
				user.Password, err = createHashOfPassword(user.Password)
				if err != nil {
					common.ErrorResponse(r.URL.Path, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
				} else {
					err = insertUserIntoDB(user)
					if err != nil {
						common.ErrorResponse(r.URL.Path, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
					} else {
						generateSuccessResponse(r.URL.Path, w)
					}
				}
			}
		}
	}

	wg.Wait()
}
