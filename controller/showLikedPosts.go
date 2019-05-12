package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../dto"
	"../model"
	"../utils"
)

// ShowLikedPosts returns resutls in JSON format.
func ShowLikedPosts(w http.ResponseWriter, r *http.Request) {
	usrID := "user_id"
	userIDStr := r.URL.Query().Get(usrID)
	userID, _ := strconv.Atoi(userIDStr)
	successfulCode := 0
	// Execute showLikedPosts model and returns json
	likedPosts := model.ShowLikedPosts(userID)

	resStruct := dto.PostList{
		Status:    http.StatusOK,
		ErrorCode: successfulCode,
		Posts:     likedPosts,
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
