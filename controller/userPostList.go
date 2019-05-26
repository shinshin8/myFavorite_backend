package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../dto"
	"../model"
	"../utils"
)

// UserPostsList shows specific user's posts list in JSON.
func UserPostsList(w http.ResponseWriter, r *http.Request) {
	// listening port
	port := portConfig.Port.Port
	// Set CORS
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	w.Header().Set(utils.Cors, "http://localhost"+port)
	w.Header().Set(utils.ArrowHeader, utils.ContentType)
	w.Header().Set(utils.Credential, utils.True)
	// Get user id from URL query parameter and convert its type string to int.
	usrID := "user_id"
	userIDStr := r.URL.Query().Get(usrID)
	userID, _ := strconv.Atoi(userIDStr)

	// Execute get user's posts list
	postList := model.UserPostsList(userID)
	successfulCode := 0
	resStruct := dto.PostList{
		Status:    http.StatusOK,
		ErrorCode: successfulCode,
		Posts:     postList,
	}

	res, err := json.Marshal(resStruct)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Response JSON
	w.Write(res)
}
