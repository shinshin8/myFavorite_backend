package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../dto"
	"../model"
	"../utils"
)

// FavoritePost insert new favorite post.
func FavoritePost(w http.ResponseWriter, r *http.Request) {
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
	// Get article id from URL query parameter and convert its type string to int.
	atcID := "article_id"
	articleIDStr := r.URL.Query().Get(atcID)
	articleID, _ := strconv.Atoi(articleIDStr)

	// Execute register liked post
	res := model.FavoritePost(userID, articleID)

	if res {
		successfulLoginCode := 0
		// set values in structs
		resultjson := dto.SimpleResutlJSON{
			Status:    http.StatusOK,
			ErrorCode: successfulLoginCode,
		}
		// convert structs to json
		res, err := json.Marshal(resultjson)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(res)
	} else {
		failedLoginCode := 12
		// set values in structs
		resultjson := dto.SimpleResutlJSON{
			Status:    http.StatusOK,
			ErrorCode: failedLoginCode,
		}
		// convert structs to json
		res, err := json.Marshal(resultjson)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(res)
	}
}
