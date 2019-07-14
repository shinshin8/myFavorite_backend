package controller

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/shinshin8/myFavorite_backend/dto"
	"github.com/shinshin8/myFavorite_backend/model"
	"github.com/shinshin8/myFavorite_backend/utils"
)

// ChangeIcon changes a picture of user's profile
func ChangeIcon(w http.ResponseWriter, r *http.Request) {
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
	// create an AWS session which can be
	// reused if we're uploading many files
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("ID"),
			os.Getenv("KEY"),
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
	// Get user's icon url.
	var ProfileIcon dto.ProfileIcon
	// TODO: delete baseURL
	iconURL := ProfileIcon.Icon
	// TODO: delete user's icon
	deleteIcon := utils.DeleteIcon(s, iconURL)
	if !deleteIcon {
		resultjson := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: utils.FailedDeleteIconFromS3,
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
	// Field name of profile icon.
	fieldName := "profile_icon"
	file, fileHeader, err := r.FormFile(fieldName)
	// Uploading icon to AWS S3.
	iconPath, uploadError := utils.UploadIcon(s, file, fileHeader)
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
	newIconURL := os.Getenv("S3_URL") + iconPath
	// Delete icon url from DB.
	deleteIconFromDB := model.UpdateIcon(newIconURL, userID)
	if !deleteIconFromDB {
		resultjson := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: utils.FailedUpdateIcon,
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
