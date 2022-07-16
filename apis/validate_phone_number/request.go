package validate_phone_number_api

import (
	"find_competitor/common"
	"find_competitor/configs"
	"log"
	"net/http"
)

func ValidatePhoneNumber(w http.ResponseWriter, r *http.Request) {
	// response header
	w.Header().Set("Content-Type", "application/json")

	// request url
	requestUrl := r.URL.Path

	// get logger
	logger, err := common.GetLogger("validate_phone_number_api")
	if err != nil {
		log.Println(common.LOGGER_ERROR_MESSAGE)
		common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
	} else {
		// get db instance
		isScript := true
		db, err := configs.GetDbInstance(isScript)

		if err != nil {
			logger.Println(err)
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
				isValidated, err := validatePhoneNumber(db, requestData.PhoneNumber)
				if err != nil {
					logger.Println(err)
					common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
				} else {
					generateSuccessResponse(requestUrl, isValidated, w)
				}
			}
		}
	}
}
