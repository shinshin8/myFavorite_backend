package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/shinshin8/myFavorite_backend/dto"
	"github.com/shinshin8/myFavorite_backend/model"
	"github.com/shinshin8/myFavorite_backend/utils"
)

// UploadingIcon is a function that uploading images to AWS S3.
func UploadingIcon(w http.ResponseWriter, r *http.Request) {
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
	atcID := "article_id"
	articleIDStr := r.URL.Query().Get(atcID)
	articleID, _ := strconv.Atoi(articleIDStr)

	// allow only 1MB of file size
	maxSize := int64(1024000)
	err := r.ParseMultipartForm(maxSize)
	if err != nil {
		// set values in structs
		resultjson := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: utils.OverSizeIcon,
		}
		// convert structs to json
		res, err := json.Marshal(resultjson)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write(res)
		return
	}
	// Field name of profile icon.
	fieldName := "profile_icon"
	file, fileHeader, err := r.FormFile(fieldName)
	if err != nil {
		resultjson := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: utils.NoIconSelected,
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
	defer file.Close()
	// create an AWS session which can be
	// reused if we're uploading many files
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(utils.Region),
		Credentials: credentials.NewStaticCredentials(
			utils.ID,
			utils.Key,
			""),
	})
	if err != nil {
		resultjson := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: utils.NoIconSelected,
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
	// Uploading icon to AWS S3.
	iconPath, uploadError := utils.UploadIconToS3(s, file, fileHeader)
	if uploadError != nil {
		resultjson := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: utils.NoIconSelected,
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
	// Icon URL
	iconURL := "https://findmyfavorite.s3-ap-northeast-1.amazonaws.com/" + iconPath
	//Insert DB
	res := model.RegisterIcon(iconURL, articleID)

	if res {
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
	} else {
		resultjson := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: utils.FailedRegisterIcon,
		}
		// convert structs to json
		res, err := json.Marshal(resultjson)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
	}
}