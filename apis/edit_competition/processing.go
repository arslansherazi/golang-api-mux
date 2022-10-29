package edit_competition_api

import (
	"find_competitor/common"
	"find_competitor/models"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func processRequestParams(r *http.Request) (models.Competition, error, bool, []string, []multipart.File) {
	var requestData models.Competition

	parseMultipartFormError := r.ParseMultipartForm(10 * 1024 * 1024)
	if parseMultipartFormError != nil {
		return requestData, parseMultipartFormError, false, nil, nil
	}

	var err error

	// user id
	userID := r.PostForm.Get("user_id")
	if userID != "" {
		requestData.UserID, err = strconv.ParseUint(userID, 10, 64)
		if err != nil {
			return models.Competition{}, err, false, nil, nil
		}
	}

	// competition id
	competitionID := r.PostForm.Get("competition_id")
	if competitionID != "" {
		requestData.ID, err = strconv.ParseUint(competitionID, 10, 64)
		if err != nil {
			return models.Competition{}, err, false, nil, nil
		}
	}

	// handle title
	title := r.PostForm.Get("title")
	if title != "" {
		requestData.Title = title
	}

	// handle description
	description := r.PostForm.Get("description")
	if description != "" {
		requestData.Description = description
	}

	// handle latitude
	latitude := r.PostForm.Get("latitude")
	if latitude != "" {
		requestData.Latitude, err = strconv.ParseFloat(latitude, 64)
		if err != nil {
			return models.Competition{}, err, false, nil, nil
		}
	}

	// handle longitude
	longitude := r.PostForm.Get("longitude")
	if longitude != "" {
		requestData.Longitude, err = strconv.ParseFloat(longitude, 64)
		if err != nil {
			return models.Competition{}, err, false, nil, nil
		}
	}

	// handle address
	address := r.PostForm.Get("address")
	if address != "" {
		requestData.Address = address
	}

	// handle starting date
	startingDate := r.PostForm.Get("starting_date")
	if startingDate != "" {
		requestData.StartingDate = startingDate
	}

	// handle starting time
	startingTime := r.PostForm.Get("starting_time")
	if startingTime != "" {
		requestData.StartingTime = startingTime
	}

	// handle ending time
	endingTime := r.PostForm.Get("ending_time")
	if endingTime != "" {
		requestData.EndingTime = endingTime
	}

	// handle deleted images
	var deletedImages []string
	deletedImagesData := r.PostForm.Get("deleted_images")

	if deletedImagesData != "" {
		deletedImages = strings.Split(deletedImagesData, ",")
	}

	// handle added images
	var images []multipart.File
	var imageFile multipart.File

	imagesData := r.MultipartForm.File["images"]

	for _, image := range imagesData {
		imageFile, err = image.Open()
		if err != nil {
			return models.Competition{}, err, true, nil, nil
		}
		images = append(images, imageFile)
	}

	// validate the request data
	validate := validator.New()
	err = validate.Struct(requestData)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return requestData, validationErrors, true, nil, nil
	}

	return requestData, nil, false, deletedImages, images
}

func editCompetition(db *gorm.DB, competition models.Competition, competitionURLs []string, competitionURLsCurrentLength int) error {
	if len(competitionURLs) != competitionURLsCurrentLength {
		competition.Images = competitionURLs
	}

	err := models.EditCompetition(db, competition)
	if err != nil {
		return nil
	}
	return nil
}

func getCompetitionImagesURLs(db *gorm.DB, competitionID uint64) ([]string, error) {
	competitionImagesURLs, err := models.GetCompetitionImagesURLs(db, competitionID)
	if err != nil {
		return nil, err
	}
	return competitionImagesURLs, nil
}

func handleDeletedImages(competitionImagesURLs []string, deletedImages []string) ([]string, error) {
	for _, deletedImageURL := range deletedImages {
		err := common.DeleteFile(deletedImageURL, common.COMPETITION_IMAGE_TYPE)
		if err != nil {
			return nil, err
		} else {
			var updatedImagesURLs []string
			for _, imageURL := range competitionImagesURLs {
				if imageURL != deletedImageURL {
					updatedImagesURLs = append(updatedImagesURLs, imageURL)
				}
			}
			competitionImagesURLs = updatedImagesURLs
		}
	}
	return competitionImagesURLs, nil
}

func handleNewlyAddedImages(addedImages []multipart.File, competitionImagesURLs []string) ([]string, error) {
	for _, image := range addedImages {
		imageURL, err := common.UploadFile(image, common.COMPETITION_IMAGE_TYPE)
		if err != nil {
			return nil, err
		} else {
			competitionImagesURLs = append(competitionImagesURLs, imageURL)
		}
	}
	return competitionImagesURLs, nil
}
