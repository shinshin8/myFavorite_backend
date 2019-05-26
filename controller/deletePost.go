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
	// listening port
	port := portConfig.Port.Port
	// Set CORS
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	w.Header().Set(utils.Cors, "http://localhost"+port)
	w.Header().Set(utils.ArrowHeader, utils.ContentType)
	w.Header().Set(utils.Credential, utils.True)
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
	// Response JSON
	w.Write(res)
}
