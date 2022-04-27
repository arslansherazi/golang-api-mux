package common

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func GetS3Uploader() *s3manager.Uploader {
	session, _ := session.NewSession(&aws.Config{Region: aws.String(AWS_REGION)})
	s3Uploader := s3manager.NewUploader(session)
	return s3Uploader
}

func UploadFileIntoS3(s3Uploader *s3manager.Uploader, bucket string, fileName string, fileData *bytes.Reader) {
	result, error := s3Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(AWS_FND_COMP_BUCKET),
		Key:    aws.String(fileName),
		Body:   fileData,
		ACL:    aws.String("public-read"),
	})
	fmt.Println(result)
	fmt.Println(error)
}
