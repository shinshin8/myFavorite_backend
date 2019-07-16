package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/shinshin8/myFavorite_backend/dto"
	"github.com/shinshin8/myFavorite_backend/model"
	"github.com/shinshin8/myFavorite_backend/utils"
)

// DeleteImages delete user posts' images.
func DeleteImages(w http.ResponseWriter, r *http.Request) {
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
		// set values in structs
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

	var deleteImagePostBody dto.DeleteImage

	er := json.NewDecoder(r.Body).Decode(&deleteImagePostBody)

	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get URL
	imageURL := deleteImagePostBody.ImageURL
	// Validate URL
	if !utils.IsImageURL(imageURL) {
		// set values in structs
		resultjson := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: utils.InvalidDeleteImageURL,
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

	// Create an AWS session
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
	// Array for data for delete from DB.
	deleteArray := []dto.ImageStruct{}
	for _, str := range imageURL {
		// Get base url part.
		baseURL := os.Getenv("S3_URL")
		// Trim base url from url.
		targetBucket := strings.Trim(str, baseURL)
		// Delete bucket from S3.
		deleteIcon := utils.DeleteBucket(session, targetBucket)
		if !deleteIcon {
			resultjson := dto.SimpleResutlJSON{
				Status:    false,
				ErrorCode: utils.FailedDeleteImages,
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
			ImageURL:  targetBucket,
			UserID:    userID,
			ArticleID: articleID,
		}

		deleteArray = append(deleteArray, imageData)
	}
	fmt.Println(deleteArray)
	// Delete data from DB.
	deleteImagesFromDB := model.DeleteImage(deleteArray)

	if !deleteImagesFromDB {
		resultjson := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: utils.FailedDeleteImageFromDB,
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
