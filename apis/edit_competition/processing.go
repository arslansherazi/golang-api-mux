package edit_competition_api

import (
	"encoding/json"
	"find_competitor/models"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func processRequestParams(r *http.Request) (Validator, error, bool) {
	var requestData Validator

	parseMultipartFormError := r.ParseMultipartForm(10 * 1024 * 1024)
	if parseMultipartFormError != nil {
		return requestData, parseMultipartFormError, false, nil
	}

	var err error

	// handle title
	requestData.Title = r.PostForm.Get("title")

	// handle description
	requestData.Description = r.PostForm.Get("description")

	// handle latitude
	latitude := r.PostForm.Get("latitude")
	if latitude != "" {
		requestData.Latitude, err = strconv.ParseFloat(latitude, 64)
		if err != nil {
			return Validator{}, err, false
		}
	}

	// handle longitude
	longitude := r.PostForm.Get("longitude")
	if longitude != "" {
		requestData.Longitude, err = strconv.ParseFloat(longitude, 64)
		if err != nil {
			return Validator{}, err, false
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

	// handle deleted images
	var deletedImages []string
	deletedImagesJson := r.PostForm.Get("deleted_images")

	if deletedImagesJson != "" {
		var deletedImagesData []string
		json.Unmarshal([]byte(deletedImagesJson), &deletedImagesData)
		for _, deletedImage := range deletedImagesData {
			deletedImages = append(deletedImages, deletedImage)
		}
		requestData.DeletedImages = deletedImages
	}

	// handle added images
	var images []multipart.File
	var imageFile multipart.File

	imagesData := r.MultipartForm.File["images"]

	for _, image := range imagesData {
		imageFile, err = image.Open()
		if err != nil {
			return Validator{}, err, true
		}
		images = append(images, imageFile)
	}
	requestData.AddedImages = images

	// validate the request data
	validate := validator.New()
	err = validate.Struct(requestData)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return requestData, validationErrors, true
	}

	return requestData, nil, false
}

func editCompetition(db *gorm.DB, competition models.Competition) error {
	err := models.InsertCompetitionIntoDB(db, competition)
	return err
}
