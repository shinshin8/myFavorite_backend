package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/shinshin8/myFavorite_backend/dto"
	"github.com/shinshin8/myFavorite_backend/model"
	"github.com/shinshin8/myFavorite_backend/utils"
)

// EditPost edits a existing post.
func EditPost(w http.ResponseWriter, r *http.Request) {
	// Set CORS
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	w.Header().Set(utils.Cors, utils.CorsWildCard)
	w.Header().Set(utils.ArrowHeader, utils.ContentType)
	w.Header().Set(utils.ArrowMethods, utils.Methods)
	w.Header().Set(utils.Credential, utils.True)
	// Get jwt from header.
	reqToken := r.Header.Get(utils.Authorization)
	// Check if jwt is verified.
	userID := utils.VerifyToken(reqToken)
	if userID == 0 {
		resultjson := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: utils.InvalidToken,
		}
		// convert structs to json
		res, err := json.Marshal(resultjson)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
	// Get article id from URL query parameter and convert its type string to int.
	atcID := "article_id"
	articleIDStr := r.URL.Query().Get(atcID)
	articleID, _ := strconv.Atoi(articleIDStr)

	var editPostBody dto.EditPostBody

	er := json.NewDecoder(r.Body).Decode(&editPostBody)

	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get title
	title := editPostBody.Title
	//Get content
	content := editPostBody.Content

	// Check title
	if !utils.IsTitle(title) {
		// Set values into the struct
		resStruct := dto.NewPost{
			Status:    false,
			ErrorCode: utils.InvalidEditTitle,
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
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	}

	// Check content
	if !utils.IsContent(content) {
		// Set values into the struct
		resStruct := dto.NewPost{
			Status:    false,
			ErrorCode: utils.InvalidEditContent,
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
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	}

	// Execute update data to DB.
	result := model.EditPost(userID, articleID, title, content)

	if result {
		// set values in structs
		resultjson := dto.SimpleResutlJSON{
			Status:    true,
			ErrorCode: utils.SuccessCode,
		}
		// convert structs to json
		res, err := json.Marshal(resultjson)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		// set values in structs
		resultjson := dto.SimpleResutlJSON{
			Status:    true,
			ErrorCode: utils.FailedEditPost,
		}
		// convert structs to json
		res, err := json.Marshal(resultjson)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
