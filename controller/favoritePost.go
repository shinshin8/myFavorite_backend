package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../dto"
	"../utils"
)

// FavoritePost insert new favorite post.
func FavoritePost(w http.ResponseWriter, r *http.Request) {
	// Get user id from URL query parameter and convert its type string to int.
	usrID := "user_id"
	userIDStr := r.URL.Query().Get(usrID)
	userID, _ := strconv.Atoi(userIDStr)
	// Get article id from URL query parameter and convert its type string to int.
	atcID := "article_id"
	articleIDStr := r.URL.Query().Get(atcID)
	articleID, _ := strconv.Atoi(articleIDStr)

	// Execute register liked post
	res := model.favoritePost(userID, articleID)

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
		// set header and defined response type for json
		w.Header().Set(utils.ContentType, utils.ApplicationJSON)
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

		// set header and defined response type for json
		w.Header().Set(utils.ContentType, utils.ApplicationJSON)
		w.Write(res)
	}
}
