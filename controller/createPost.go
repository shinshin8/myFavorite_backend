package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../dto"
	"../model"
	"../utils"
)

var portConfig dto.PortConfig

// CreatePost creates a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	// listening port
	port := portConfig.Port.Port
	// Set CORS
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	w.Header().Set(utils.Cors, "http://localhost"+port)
	w.Header().Set(utils.ArrowHeader, utils.ContentType)
	w.Header().Set(utils.Credential, utils.True)
	// input form name
	var (
		usrID = "user_id"
		til   = "title"
		cont  = "content"
	)
	// Get user id
	userIDStr := r.PostFormValue(usrID)
	userID, _ := strconv.Atoi(userIDStr)
	// Get title
	title := r.PostFormValue(til)
	//Get content
	content := r.PostFormValue(cont)

	// Check userID
	if !utils.IsID(userID) {
		// Invalid user id
		invalidUserID := 14
		// Set values into the struct
		resStruct := dto.NewPost{
			Status:    http.StatusOK,
			ErrorCode: invalidUserID,
			UserID:    userID,
			Title:     title,
			Content:   content,
		}
		// convert struct to JSON
		res, err := json.Marshal(resStruct)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Response JSON
		w.Write(res)
		return
	}

	// Check title
	if !utils.IsTitle(title) {
		// Invalid title
		invalidTitle := 15
		// Set values into the struct
		resStruct := dto.NewPost{
			Status:    http.StatusOK,
			ErrorCode: invalidTitle,
			UserID:    userID,
			Title:     title,
			Content:   content,
		}
		// convert struct to JSON
		res, err := json.Marshal(resStruct)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Response JSON
		w.Write(res)
		return
	}

	// Check content
	if !utils.IsContent(content) {
		// Invalid content
		invalidContent := 16
		// Set values into the struct
		resStruct := dto.NewPost{
			Status:    http.StatusOK,
			ErrorCode: invalidContent,
			UserID:    userID,
			Title:     title,
			Content:   content,
		}
		// convert struct to JSON
		res, err := json.Marshal(resStruct)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Response JSON
		w.Write(res)
		return
	}

	// Execute insert data to DB.
	result := model.CreateNewPost(userID, title, content)

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
