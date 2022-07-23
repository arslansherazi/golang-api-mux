package common

import (
	"bytes"
	"image"
	"image/png"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/efimovalex/stackerr"
	"github.com/google/uuid"
	"github.com/sunshineplan/imgconv"
)

func GetS3Uploader() (*s3manager.Uploader, error) {
	session, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})
	if err != nil {
		return nil, err
	}
	s3Uploader := s3manager.NewUploader(session)
	return s3Uploader, nil
}

func UploadFileIntoS3(s3Uploader *s3manager.Uploader, bucket string, folderName string, fileName string, fileData *bytes.Reader) error {
	_, err := s3Uploader.Upload(&s3manager.UploadInput{
		Bucket: &bucket,
		Key:    aws.String(path.Join(folderName, fileName)),
		Body:   fileData,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return err
	}
	return nil
}

func GetLogger(fileName string) (*log.Logger, error) {
	absPath, err := filepath.Abs("./logs")
	if err != nil {
		return nil, err
	}
	_ = os.Mkdir(absPath, os.ModePerm) // create logs folder if not exists
	loggerFile, err := os.OpenFile(absPath+"/"+fileName+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return log.New(loggerFile, "", log.Ldate|log.Ltime|log.Lshortfile), nil
}

func ParseValidationError(err error) string {
	errorMessage := strings.Split(err.Error(), "\n")[0]
	errors := strings.Split(errorMessage, "'")
	errorField := ToSnakeCase(errors[3])
	validationError := errors[5]
	if validationError == "required" {
		validationError = REQUIRED_FIELD_VALIDATION_MESSAGE
	}
	return errorField + ": " + validationError
}

func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func LogError(logger *log.Logger, err error) {
	errorStack := stackerr.NewFromError(err).StackWithContext(err.Error()).Sprint()
	logger.Print(errorStack + "\n")
}

func HandleImage(imageFile multipart.File, fileType string) (*bytes.Reader, error) {
	// convert multipart file into image file
	image, _, err := image.Decode(imageFile)
	if err != nil {
		return nil, err
	}
	// resize image
	var imageWidth, imageHeight int
	if fileType == PROFILE_IMAGE_TYPE {
		imageWidth = PROFILE_IMAGE_WIDTH
		imageHeight = PROFILE_IMAGE_HEIGHT
	} else if fileType == COMPETITION_IMAGE_TYPE {
		imageWidth = COMPETITION_IMAGE_WIDTH
		imageHeight = COMPETITION_IMAGE_HEIGHT
	}

	convertedImage := imgconv.Resize(
		image, imgconv.ResizeOption{Width: imageWidth, Height: imageHeight},
	)
	// convert format of image to png
	imgconv.Write(io.Discard, convertedImage, imgconv.FormatOption{Format: imgconv.PNG})

	encodedImage, err := EncodeFile(fileType, convertedImage, nil)
	if err != nil {
		return nil, err
	}

	return encodedImage, nil
}

func EncodeFile(fileType string, imageFile image.Image, file multipart.File) (*bytes.Reader, error) {
	// convert image into bytes data
	buffer := new(bytes.Buffer)
	if fileType == PROFILE_IMAGE_TYPE || fileType == COMPETITION_IMAGE_TYPE {
		err := png.Encode(buffer, imageFile)
		if err != nil {
			return nil, err
		}
	} else {
		// handle files other than images
	}
	imageData := bytes.NewReader(buffer.Bytes())

	return imageData, nil
}

func UploadFile(file multipart.File, fileType string) (string, error) {
	var fileURL string
	if file != nil {
		s3Uploader, err := GetS3Uploader()
		if err != nil {
			return "", err
		}

		var fileExtension string
		var compressedFile *bytes.Reader
		if fileType == PROFILE_IMAGE_TYPE || fileType == COMPETITION_IMAGE_TYPE {
			fileExtension = IMAGES_EXTENSION
			compressedFile, err = HandleImage(file, fileType)
			if err != nil {
				return "", err
			}
		} else {
			// handle files other than images
		}
		fileName := uuid.New().String() + fileExtension

		// get s3 bucket folder name
		var folderName string
		if fileType == PROFILE_IMAGE_TYPE {
			folderName = os.Getenv("AWS_FND_COMP_BUCKET_PROFILE_IMAGES_FOLDER")
		} else if fileType == COMPETITION_IMAGE_TYPE {
			folderName = os.Getenv("AWS_FND_COMP_BUCKET_COMPETITION_IMAGES_FOLDER")
		}

		// upload file into s3 and generate file url
		err = UploadFileIntoS3(s3Uploader, os.Getenv("AWS_FND_COMP_BUCKET"), folderName, fileName, compressedFile)
		if err != nil {
			return "", err
		} else {
			fileURL = os.Getenv("BUCKET_BASE_URL") + folderName + "/" + fileName
		}
	} else {
		fileURL = ""
	}

	return fileURL, nil
}
