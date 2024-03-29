package signup_api

import (
	"find_competitor/common"
	"find_competitor/configs"
	"log"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	// response header
	w.Header().Set("Content-Type", "application/json")

	// request url
	requestUrl := r.URL.Path

	logger, err := common.GetLogger("signup_api")
	if err != nil {
		log.Println(common.LOGGER_ERROR_MESSAGE)
		common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
	} else {

		// get db instance
		isScript := true
		db, err := configs.GetDbInstance(isScript)

		if err != nil {
			common.LogError(logger, err)
			common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
		} else {
			user, profileImage, err, isValidationError := processRequestParams(logger, r)
			if err != nil {
				if isValidationError {
					validationMessage := common.ParseValidationError(err)
					common.ErrorResponse(requestUrl, http.StatusUnprocessableEntity, validationMessage, w)
				} else {
					common.LogError(logger, err)
					common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
				}
			} else {
				user.ProfileImageUrl, err = common.UploadFile(profileImage, common.PROFILE_IMAGE_TYPE)
				if err != nil {
					common.LogError(logger, err)
					common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
				} else {
					user.Password, err = createHashOfPassword(user.Password)
					if err != nil {
						common.LogError(logger, err)
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
