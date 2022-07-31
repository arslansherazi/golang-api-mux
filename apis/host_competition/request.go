package host_competition_api

import (
	"find_competitor/common"
	"find_competitor/configs"
	"log"
	"net/http"
)

func HostCompetition(w http.ResponseWriter, r *http.Request) {
	// response header
	w.Header().Set("Content-Type", "application/json")

	// request url
	requestUrl := r.URL.Path

	// get logger
	logger, err := common.GetLogger("find_competition_api")
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
			competitionData, err, isValidationError, images := processRequestParams(r)
			if err != nil {
				if isValidationError {
					validationMessage := common.ParseValidationError(err)
					common.ErrorResponse(requestUrl, http.StatusUnprocessableEntity, validationMessage, w)
				} else {
					common.LogError(logger, err)
					common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
				}
			} else {
				// handle images

				imagesURLs, err := uploadImages(images)
				if err != nil {
					common.LogError(logger, err)
					common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
				} else {
					competitionData.Images = getCompetitionImagesData(imagesURLs)
					err := insertCompetitionIntoDB(db, competitionData)
					if err != nil {
						common.LogError(logger, err)
						common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
					} else {
						generateSuccessResponse(requestUrl, w)
					}
				}

			}
		}
	}
}
