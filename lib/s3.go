package prombackup

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func (app *App) UploadFile(filename string) error {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(app.S3Region)})
	if err != nil {
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	uploader := s3manager.NewUploader(sess)

	// Upload the file's body to S3 bucket as an object with the key being the
	// same as the filename.
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(app.S3Bucket),
		Key:    aws.String(filename),
		Body:   file,
	})

	return err
}
