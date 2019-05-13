package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../dto"
	"../model"
	"../utils"
)

// ShowFavoritePosts returns the list of favorite posts
func ShowFavoritePosts(w http.ResponseWriter, r *http.Request) {
	// Get user id from the URL query paramter in string type and conver it to int type.
	usrID := "user_id"
	userIDStr := r.URL.Query().Get(usrID)
	userID, _ := strconv.Atoi(userIDStr)
	// Get article id from the URL query parameter in string type and conver it to int type.
	atlID := "article_id"
	articleIDStr := r.URL.Query().Get(atlID)
	articleID, _ := strconv.Atoi(articleIDStr)

	// Execute get all favorite posts list and response it in JSON.
	successfulCode := 0
	favoritePosts := model.ShowFavoritePosts(userID, articleID)

	resStruct := dto.PostList{
		Status:    http.StatusOK,
		ErrorCode: successfulCode,
		Posts:     favoritePosts,
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
