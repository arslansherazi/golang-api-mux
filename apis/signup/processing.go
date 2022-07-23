package signup_api

import (
	"find_competitor/models"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func processRequestParams(logger *log.Logger, r *http.Request) (models.User, multipart.File, error, bool) {
	var user models.User
	parseMultipartFormError := r.ParseMultipartForm(5 * 1024 * 1024)
	if parseMultipartFormError != nil {
		return user, nil, parseMultipartFormError, false
	}
	user.Name = r.PostForm.Get("name")
	user.PhoneNumber = r.PostForm.Get("phone_number")
	user.Password = r.PostForm.Get("password")
	profileImage, _, profileImageError := r.FormFile("profile_image")
	if profileImageError != nil {
		return user, nil, profileImageError, false
	}
	defer profileImage.Close()

	// validate the request data
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return user, nil, validationErrors, true
	}

	return user, profileImage, nil, false
}

func createHashOfPassword(password string) (string, error) {
	hashedPasswordData, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return "", err
	}
	hashedPassword := string(hashedPasswordData[:])
	return hashedPassword, nil
}

func insertUserIntoDB(db *gorm.DB, user models.User) {
	models.InsertUserIntoDB(db, user)
}
