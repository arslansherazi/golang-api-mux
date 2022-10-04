package participation_api

import (
	"find_competitor/common"
	"find_competitor/configs"
	"fmt"
	"log"
	"net/http"
)

func Participation(w http.ResponseWriter, r *http.Request) {
	// response header
	w.Header().Set("Content-Type", "application/json")

	// request url
	requestUrl := r.URL.Path

	// get logger
	logger, err := common.GetLogger("participation_api")
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
			participationData, err, isValidationError := processRequestParams(r)
			if err != nil {
				if isValidationError {
					validationMessage := common.ParseValidationError(err)
					common.ErrorResponse(requestUrl, http.StatusUnprocessableEntity, validationMessage, w)
				} else {
					common.LogError(logger, err)
					common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
				}
			} else {
				// models.VerifyParticipant(db, participationData)
				fmt.Println(participationData)
			}
		}
	}
}
