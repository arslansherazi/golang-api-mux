package host_competition_api

import (
	"find_competitor/models"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func processRequestParams(r *http.Request) (models.Competition, error, bool, []multipart.File) {
	var requestData models.Competition

	parseMultipartFormError := r.ParseMultipartForm(10 * 1024 * 1024)
	if parseMultipartFormError != nil {
		return requestData, parseMultipartFormError, false, nil
	}

	var err error

	// handle user id
	userID := r.PostForm.Get("user_id")
	if userID != "" {
		requestData.UserID, err = strconv.Atoi(userID)
		if err != nil {
			return models.Competition{}, err, false, []multipart.File{}
		}
	}

	// handle title
	requestData.Title = r.PostForm.Get("title")

	// handle description
	requestData.Description = r.PostForm.Get("description")

	// handle latitude
	latitude := r.PostForm.Get("latitude")
	if latitude != "" {
		requestData.Latitude, err = strconv.ParseFloat(latitude, 64)
		if err != nil {
			return models.Competition{}, err, false, []multipart.File{}
		}
	}

	// handle longitude
	longitude := r.PostForm.Get("longitude")
	if longitude != "" {
		requestData.Longitude, err = strconv.ParseFloat(longitude, 64)
		if err != nil {
			return models.Competition{}, err, false, []multipart.File{}
		}
	}

	// handle address
	requestData.Address = r.PostForm.Get("address")

	// handle starting date
	requestData.StartingDate = r.PostForm.Get("starting_date")

	// handle starting time
	requestData.StartingTime = r.PostForm.Get("starting_time")

	// handle ending time
	requestData.EndingTime = r.PostForm.Get("ending_time")

	// handle images
	var images []multipart.File
	var imageFile multipart.File

	imagesData := r.MultipartForm.File["images"]

	for _, image := range imagesData {
		imageFile, err = image.Open()
		if err != nil {
			return models.Competition{}, err, true, []multipart.File{}
		}
		images = append(images, imageFile)
	}

	// validate the request data
	validate := validator.New()
	err = validate.Struct(requestData)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return requestData, validationErrors, true, []multipart.File{}
	}

	return requestData, nil, false, images
}

func insertCompetitionIntoDB(db *gorm.DB, competition models.Competition) error {
	err := models.InsertCompetitionIntoDB(db, competition)
	return err
}
