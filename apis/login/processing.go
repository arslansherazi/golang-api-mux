package login_api

import (
	"find_competitor/common"
	"find_competitor/models"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func processRequestParams(r *http.Request) (Validator, error, bool) {
	var requestData Validator

	parseMultipartFormError := r.ParseMultipartForm(5 * 1024 * 1024)
	if parseMultipartFormError != nil {
		return requestData, parseMultipartFormError, false
	}

	requestData.PhoneNumber = r.PostForm.Get("phone_number")
	requestData.Password = r.PostForm.Get("password")

	// validate the request data
	validate := validator.New()
	err := validate.Struct(requestData)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return requestData, validationErrors, true
	}

	return requestData, nil, false
}

func validateUser(db *gorm.DB, phoneNumber string) models.User {
	userData := models.GetUserData(db, phoneNumber)
	return userData
}

func verifyPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func generateToken(phoneNumber string) (string, error) {
	var secretKey = []byte(common.JWT_SECRET_KEY)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"phoneNumber": phoneNumber,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secretKey)
	return tokenString, err
}
