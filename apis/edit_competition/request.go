package edit_competition_api

import (
	"find_competitor/common"
	"find_competitor/configs"
	"log"
	"net/http"
	"strings"
)

func EditCompetition(w http.ResponseWriter, r *http.Request) {
	// response header
	w.Header().Set("Content-Type", "application/json")

	// request url
	requestUrl := r.URL.Path

	// get logger
	logger, err := common.GetLogger("edit_competition_api")
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
			competitionData, err, isValidationError, deletedImages, addedImages := processRequestParams(r)
			if err != nil {
				if isValidationError {
					validationMessage := common.ParseValidationError(err)
					common.ErrorResponse(requestUrl, http.StatusUnprocessableEntity, validationMessage, w)
				} else {
					common.LogError(logger, err)
					common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
				}
			} else {
				competitionImagesData, err := getCompetitionImagesData(db, competitionData.ID)
				if err != nil {
					common.LogError(logger, err)
					common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
				} else {
					competitionImagesURLs := strings.Split(competitionImagesData, ",")
					if err != nil {
						common.LogError(logger, err)
						common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
					} else {
						competitionURLsCurrentLength := len(competitionImagesURLs)

						// handle newly added images
						competitionImagesURLs, err := handleNewlyAddedImages(addedImages, competitionImagesURLs)
						if err != nil {
							common.LogError(logger, err)
							common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
						} else {
							// handle deleted images
							competitionImagesURLs, err = handleDeletedImages(competitionImagesURLs, deletedImages)
							if err != nil {
								common.LogError(logger, err)
								common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
							} else {
								err := editCompetition(db, competitionData, competitionImagesURLs, competitionURLsCurrentLength)
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
		}
	}
}
