package host_competition_api

import (
	"encoding/json"
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
				var imagesURLs []string

				for _, image := range images {
					imageURL, err := common.UploadFile(image, common.COMPETITION_IMAGE_TYPE)
					if err != nil {
						common.LogError(logger, err)
						common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
					} else {
						imagesURLs = append(imagesURLs, imageURL)
					}
				}

				// convert image urls into json
				imagesJsonData, err := json.MarshalIndent(imagesURLs, "", "\t")
				if err != nil {
					common.LogError(logger, err)
					common.ErrorResponse(requestUrl, http.StatusInternalServerError, common.INTERNAL_SERVER_ERROR_MESSAGE, w)
				} else {
					if len(imagesJsonData) > 4 {
						competitionData.Images = imagesJsonData
					}
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
