package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../model"
	"../utils"
)

// DeletePost delete specific post resource.
func DeletePost(w http.ResponseWriter, r *http.Request) {
	// Get user id from URL query paramter with string type and convert it to int.
	usrID := "user_id"
	usrIDStr := r.URL.Query().Get(usrID)
	userID, _ := strconv.Atoi(usrIDStr)
	// Get article id from URL query parameter with string type and convert it to int.
	atlID := "article_id"
	atlIDStr := r.URL.Query().Get(atlID)
	articleID, _ := strconv.Atoi(atlIDStr)

	// Execute delete resouce.
	result := model.DeletePost(userID, articleID)

	// In the Model, the function returns JSON in other way.
	// So in this part, just response result.

	// convert struct to JSON
	res, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Set HTTP header and defined MIME type
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	// Response JSON
	w.Write(res)
}
