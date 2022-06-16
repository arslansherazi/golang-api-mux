package common

import (
	"bytes"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func GetS3Uploader() (*s3manager.Uploader, error) {
	session, err := session.NewSession(&aws.Config{Region: aws.String(AWS_REGION)})
	if err != nil {
		return nil, err
	}
	s3Uploader := s3manager.NewUploader(session)
	return s3Uploader, nil
}

func UploadFileIntoS3(s3Uploader *s3manager.Uploader, bucket string, fileName string, fileData *bytes.Reader) error {
	_, err := s3Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(AWS_FND_COMP_BUCKET),
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
	err = os.Mkdir(absPath, os.ModePerm) // create logs folder if not exists
	if err != nil {
		return nil, err
	}
	loggerFile, err := os.OpenFile(absPath+"/"+fileName+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return log.New(loggerFile, "", log.Ldate|log.Ltime|log.Lshortfile), nil
}

func ParseValidationError(err error) string {
	return ""
}
