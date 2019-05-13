package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../dto"
	"../model"
	"../utils"
)

// DeleteFavoritePost deletes specific favorite post resource and returns the result in JSON
func DeleteFavoritePost(w http.ResponseWriter, r *http.Request) {
	// Get user id from URL query paramter with string type and convert it to int.
	usrID := "user_id"
	usrIDStr := r.URL.Query().Get(usrID)
	userID, _ := strconv.Atoi(usrIDStr)
	// Get article id from URL query parameter with string type and convert it to int.
	atlID := "article_id"
	atlIDStr := r.URL.Query().Get(atlID)
	articleID, _ := strconv.Atoi(atlIDStr)

	// Execute delete resouce.
	res := model.DeleteFavoritePost(userID, articleID)

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
		failedCode := 13
		// set values in structs
		resultjson := dto.SimpleResutlJSON{
			Status:    http.StatusOK,
			ErrorCode: failedCode,
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
