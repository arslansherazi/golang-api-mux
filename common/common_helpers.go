package common

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/efimovalex/stackerr"
)

func GetS3Uploader() (*s3manager.Uploader, error) {
	session, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})
	if err != nil {
		return nil, err
	}
	s3Uploader := s3manager.NewUploader(session)
	return s3Uploader, nil
}

func UploadFileIntoS3(s3Uploader *s3manager.Uploader, bucket string, fileName string, fileData *bytes.Reader) error {
	_, err := s3Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("AWS_FND_COMP_BUCKET")),
		Key:    aws.String(fileName),
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
