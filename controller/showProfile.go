package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../dto"
	"../model"
	"../utils"
)

// ShowProfile show a user's profile.
func ShowProfile(w http.ResponseWriter, r *http.Request) {
	// Get user id from the URL query paramter in string type and conver it to int type.
	usrID := "user_id"
	userIDStr := r.URL.Query().Get(usrID)
	userID, _ := strconv.Atoi(userIDStr)
	// Get profile.
	userProfile := model.ShowProfile(userID)
	successfulCode := 0

	resStruct := dto.ProfileResult{
		Status:    http.StatusOK,
		ErrorCode: successfulCode,
		Profile:   userProfile,
	}

	res, err := json.Marshal(resStruct)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Set HTTP header and defined MIME type
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	// Response JSON
	w.Write(res)
}
