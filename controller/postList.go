package controller

import (
	"encoding/json"
	"net/http"

	"../dto"
	"../model"
	"../utils"
)

// PostList is controller file for get all post with JSON format.
func PostList(w http.ResponseWriter, r *http.Request) {
	// listening port
	port := portConfig.Port.Port
	// Set CORS
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	w.Header().Set(utils.Cors, "http://localhost"+port)
	w.Header().Set(utils.ArrowHeader, utils.ContentType)
	w.Header().Set(utils.Credential, utils.True)

	successfulCode := 0
	// DB result array
	dbResultArray := model.GetPosts()

	resStruct := dto.PostList{
		Status:    http.StatusOK,
		ErrorCode: successfulCode,
		Posts:     dbResultArray,
	}

	res, err := json.Marshal(resStruct)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Response JSON
	w.Write(res)

}
