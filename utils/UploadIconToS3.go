package utils

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/globalsign/mgo/bson"
)

// UploadIconToS3 saves a file to aws bucket and returns the url to the file and an error if there's any
func UploadIconToS3(s *session.Session, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	// get the file size and read
	// the file content into a buffer
	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)
	// create a unique file name for the file
	tempFileName := "icon/" + bson.NewObjectId().Hex() + filepath.Ext(fileHeader.Filename)
	// config settings: this is where you choose the bucket,
	// filename, content-type and storage class of the file
	// you're uploading
	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(os.Getenv("BUCKET_NAME")),
		Key:                  aws.String(tempFileName),
		ACL:                  aws.String(os.Getenv("ACL_SETTING")), // could be private if you want it to be access by only authorized users
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String(os.Getenv("ATTACHEMENT")),
		ServerSideEncryption: aws.String(os.Getenv("ENCRYPTION")),
		StorageClass:         aws.String(os.Getenv("STORAGE")),
	})
	if err != nil {
		return "", err
	}
	return tempFileName, err
}
