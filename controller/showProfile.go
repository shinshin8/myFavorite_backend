package controller

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/shinshin8/myFavorite_backend/dto"
	"github.com/shinshin8/myFavorite_backend/model"
	"github.com/shinshin8/myFavorite_backend/utils"
)

// ShowProfile show a user's profile.
func ShowProfile(w http.ResponseWriter, r *http.Request) {
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
	// Get icon.
	userIcon := model.GetIcon(userID)
	// Get profile.
	userProfile, err := model.ShowProfile(userID)
	if err != nil {
		resultjson := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: utils.FailedGetProfile,
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
	// Set values into the struct
	profile := dto.Profile{
		UserID:      userID,
		IconURL:     os.Getenv("S3_URL") + userIcon,
		UserName:    userProfile.UserName,
		Birthday:    userProfile.Birthday,
		MailAddress: userProfile.MailAddress,
		Comment:     userProfile.Comment,
	}
	resStruct := dto.ProfileResult{
		Status:    true,
		ErrorCode: utils.SuccessCode,
		Profile:   profile,
	}

	res, err := json.Marshal(resStruct)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Response JSON
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	return
}
