package validate_phone_number_api

import (
	"find_competitor/models"
	"net/http"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func processRequestParams(r *http.Request) (Validator, error, bool) {
	var requestData Validator

	parseMultipartFormError := r.ParseMultipartForm(5 * 1024 * 1024)
	if parseMultipartFormError != nil {
		return requestData, parseMultipartFormError, false
	}

	requestData.PhoneNumber = r.PostForm.Get("phone_number")

	// validate the request data
	validate := validator.New()
	err := validate.Struct(requestData)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return requestData, validationErrors, true
	}

	return requestData, nil, false
}

func validatePhoneNumber(db *gorm.DB, phoneNumber string) (bool, error) {
	isValidated, err := models.ValidatePhoneNumber(db, phoneNumber)
	return isValidated, err
}
