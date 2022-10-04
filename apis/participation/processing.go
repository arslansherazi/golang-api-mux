package participation_api

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
)

func processRequestParams(r *http.Request) (Validator, error, bool) {
	var requestData Validator

	var err error

	// user id
	userID := r.PostForm.Get("user_id")
	if userID != "" {
		requestData.UserID, err = strconv.ParseUint(userID, 10, 64)
		if err != nil {
			return Validator{}, err, false
		}
	}

	// competition id
	competitionID := r.PostForm.Get("competition_id")
	if competitionID != "" {
		requestData.CompetitionID, err = strconv.ParseUint(competitionID, 10, 64)
		if err != nil {
			return Validator{}, err, false
		}
	}

	// validate the request data
	validate := validator.New()
	err = validate.Struct(requestData)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return requestData, validationErrors, true
	}

	return requestData, nil, false
}

// func editCompetition(db *gorm.DB, competition models.Competition, competitionURLs []string, competitionURLsCurrentLength int) error {
// 	if len(competitionURLs) != competitionURLsCurrentLength {
// 		if len(competitionURLs) > 1 {
// 			competition.Images = common.JoinString(competitionURLs, ",")
// 		} else {
// 			competition.Images = competitionURLs[0]
// 		}
// 	}

// 	err := models.EditCompetition(db, competition)
// 	if err != nil {
// 		return nil
// 	}
// 	return nil
// }
