package login_api

import (
	"log"
	"net/http"

	"github.com/go-playground/validator"
)

func processRequestParams(logger *log.Logger, r *http.Request) (Validator, error) {
	var requestData Validator

	requestData.PhoneNumber = r.PostForm.Get("phone_number")
	requestData.Password = r.PostForm.Get("password")

	// validate the request data
	validate := validator.New()
	err := validate.Struct(requestData)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return requestData, validationErrors
	}

	return requestData, nil
}
