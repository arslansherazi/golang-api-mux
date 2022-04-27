package signup_api

import (
	"bytes"
	"encoding/json"
	"find_competitor/common"
	"find_competitor/configs"
	"find_competitor/models"
	"image"
	"image/png"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/google/uuid"
	"github.com/sunshineplan/imgconv"
	"golang.org/x/crypto/bcrypt"
)

/* process request */
func Signup(w http.ResponseWriter, r *http.Request) {
	// response header
	w.Header().Set("Content-Type", "application/json")

	logger := common.GetLogger("signup_api")
	user, profileImage := processRequestParams(logger, r)
	user.ProfileImageUrl = generateProfileImageUrl(profileImage)
	user.Password = createHashOfPassword(user.Password)
	insertUserIntoDB(user)
	response := generateSuccessResponse(r.URL.Path)

	json.NewEncoder(w).Encode(&response)
}

/* request processing utils */
func processRequestParams(logger *log.Logger, r *http.Request) (models.User, multipart.File) {
	var user models.User
	error := r.ParseMultipartForm(5 * 1024 * 1024)
	if error != nil {
		logger.Println(error)
	}
	user.Name = r.PostForm.Get("name")
	user.PhoneNumber = r.PostForm.Get("phone_number")
	user.Password = r.PostForm.Get("password")
	profileImage, _, _ := r.FormFile("profile_image")
	defer profileImage.Close()
	return user, profileImage
}

func generateProfileImageUrl(profileImage multipart.File) string {
	var profileImageUrl string
	if profileImage != nil {
		s3Uploader := common.GetS3Uploader()
		fileName := uuid.New().String() + common.IMAGES_EXTENSION
		convertedProfileImage := handleProfileImage(profileImage)
		common.UploadFileIntoS3(s3Uploader, common.AWS_FND_COMP_BUCKET, fileName, convertedProfileImage)
	} else {
		profileImageUrl = ""
	}
	return profileImageUrl
}

func handleProfileImage(profileImage multipart.File) *bytes.Reader {
	// convert multipart file into image file
	image, _, _ := image.Decode(profileImage)
	// resize image
	convertedImage := imgconv.Resize(
		image, imgconv.ResizeOption{Width: common.PROFILE_IMAGE_WIDTH, Height: common.PROFILE_IMAGE_HEIGHT},
	)
	// convert format of image to ong
	imgconv.Write(io.Discard, convertedImage, imgconv.FormatOption{Format: imgconv.PNG})
	// convert image into bytes data
	buffer := new(bytes.Buffer)
	_ = png.Encode(buffer, convertedImage)
	imageData := bytes.NewReader(buffer.Bytes())

	return imageData
}

func createHashOfPassword(password string) string {
	hashedPasswordData, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
	hashedPassword := string(hashedPasswordData[:])
	return hashedPassword
}

func insertUserIntoDB(user models.User) {
	db := configs.GetDbInstance()
	models.InsertUserIntoDB(db, user)
}

func generateSuccessResponse(requestUrl string) SuccessResponse {
	successResponse := SuccessResponse{
		IsSignedUp:   true,
		BaseResponse: common.BaseResponse{StatusCode: 200, Success: true, Cmd: requestUrl},
	}
	return successResponse
}
