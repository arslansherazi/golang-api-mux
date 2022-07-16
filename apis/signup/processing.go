package signup_api

import (
	"bytes"
	"find_competitor/common"
	"find_competitor/models"
	"image"
	"image/png"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/sunshineplan/imgconv"
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

func generateProfileImageUrl(profileImage multipart.File) (string, error) {
	var profileImageUrl string
	if profileImage != nil {
		s3Uploader, err := common.GetS3Uploader()
		if err != nil {
			return "", err
		}
		fileName := uuid.New().String() + common.IMAGES_EXTENSION
		convertedProfileImage, err := handleProfileImage(profileImage)
		if err != nil {
			return "", err
		}
		err = common.UploadFileIntoS3(s3Uploader, os.Getenv("AWS_FND_COMP_BUCKET"), fileName, convertedProfileImage)
		if err != nil {
			return "", err
		} else {
			profileImageUrl = os.Getenv("BUCKET_BASE_URL") + fileName
		}
	} else {
		profileImageUrl = ""
	}

	return profileImageUrl, nil
}

func handleProfileImage(profileImage multipart.File) (*bytes.Reader, error) {
	// convert multipart file into image file
	image, _, err := image.Decode(profileImage)
	if err != nil {
		return nil, err
	}
	// resize image
	convertedImage := imgconv.Resize(
		image, imgconv.ResizeOption{Width: common.PROFILE_IMAGE_WIDTH, Height: common.PROFILE_IMAGE_HEIGHT},
	)
	// convert format of image to png
	imgconv.Write(io.Discard, convertedImage, imgconv.FormatOption{Format: imgconv.PNG})
	// convert image into bytes data
	buffer := new(bytes.Buffer)
	err = png.Encode(buffer, convertedImage)
	if err != nil {
		return nil, err
	}
	imageData := bytes.NewReader(buffer.Bytes())

	return imageData, nil
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
