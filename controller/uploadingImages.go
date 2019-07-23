package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/shinshin8/myFavorite_backend/dto"
	"github.com/shinshin8/myFavorite_backend/model"
	"github.com/shinshin8/myFavorite_backend/utils"
)

// UploadingImages uploads multiple photos to AWS S3.
func UploadingImages(w http.ResponseWriter, r *http.Request) {
	// Set CORS
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	w.Header().Set(utils.Cors, utils.CorsWildCard)
	w.Header().Set(utils.ArrowHeader, utils.ContentType)
	w.Header().Set(utils.ArrowMethods, utils.Methods)
	w.Header().Set(utils.Credential, utils.True)
	// Get jwt from header.
	reqToken := r.Header.Get(utils.Authorization)
	// Check if jwt is verified.
	userID := utils.VerifyToken(reqToken)
	if userID == 0 {
		resultjson := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: utils.InvalidToken,
		}
		// convert structs to json
		res, err := json.Marshal(resultjson)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(res)
		return
	}
	// Get article id from URL query parameter with string type and convert it to int.
	atlID := "article_id"
	atlIDStr := r.URL.Query().Get(atlID)
	articleID, _ := strconv.Atoi(atlIDStr)
	var maxSize int64 = 200000
	err := r.ParseMultipartForm(maxSize)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	formdata := r.MultipartForm

	var fieldName = "multiplefiles"
	files := formdata.File[fieldName]

	session, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("ID"),
			os.Getenv("KEY"),
			""),
	})
	if err != nil {
		resultjson := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: utils.FailedGenerateAWSSession,
		}
		// convert structs to json
		res, err := json.Marshal(resultjson)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}

	// Array for image path
	urlArray := []dto.ImageStruct{}

	for i := range files {
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		fileHead := files[i]
		// Get file name
		fileName := fileHead.Filename
		// Get file extension.
		fileExtension := filepath.Ext(fileName)
		// Acceptable extensions
		var (
			jpg  = ".jpg"
			jpeg = ".jpeg"
			png  = ".png"
		)
		// Check extensions
		if fileExtension != jpeg && fileExtension != jpg && fileExtension != png {
			break
		}
		// Uploading icon to AWS S3.
		imagePath, uploadError := utils.UploadingToS3(session, file, fileHead)
		if uploadError != nil {
			resultjson := dto.SimpleResutlJSON{
				Status:    false,
				ErrorCode: utils.FailedUploadImages,
			}
			// convert structs to json
			res, err := json.Marshal(resultjson)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)
			return
		}

		imageData := dto.ImageStruct{
			ImageURL:  imagePath,
			UserID:    userID,
			ArticleID: articleID,
		}

		// Push generated path to slice
		urlArray = append(urlArray, imageData)
	}
	// Insert DB
	RegisterImages := model.UploadImage(urlArray)

	if !RegisterImages {
		resultjson := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: utils.FailedUploadImages,
		}
		// convert structs to json
		res, err := json.Marshal(resultjson)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}

	resultjson := dto.SimpleResutlJSON{
		Status:    true,
		ErrorCode: utils.SuccessCode,
	}
	// convert structs to json
	res, err := json.Marshal(resultjson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	return
}
