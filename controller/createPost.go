package controller

import (
	"encoding/json"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/shinshin8/myFavorite_backend/dto"
	"github.com/shinshin8/myFavorite_backend/model"
	"github.com/shinshin8/myFavorite_backend/utils"
)

// CreatePost creates a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	// Set CORS
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	w.Header().Set(utils.Cors, utils.CorsWildCard)
	w.Header().Set(utils.ArrowHeader, utils.ContentType)
	w.Header().Set(utils.ArrowMethods, utils.Methods)
	w.Header().Set(utils.Credential, utils.True)

	// Get jwt from header.
	reqToken := r.Header.Get(utils.Authorization)
	// Check if jwt is verified.
	token, _ := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("boobar"), nil
	})
	if token == nil {
		// set values in structs
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
	// Get user id from jwt.
	claims := token.Claims.(jwt.MapClaims)
	userIDkey := "user_id"
	userIDFloat64, _ := claims[userIDkey]
	userID := int(userIDFloat64.(float64))

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
			ErrorCode: utils.InvalidCreateTitle,
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
		// Set values into the struct
		resStruct := dto.NewPost{
			Status:    false,
			ErrorCode: utils.InvalidCreateContent,
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
		w.WriteHeader(http.StatusOK)
		// Response JSON
		w.Write(res)
		return
	}
	// Execute insert data to DB.
	result := model.CreateNewPost(userID, title, content)
	// convert struct to JSON
	res, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Response JSON
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
