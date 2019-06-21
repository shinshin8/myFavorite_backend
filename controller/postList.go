package controller

import (
	"encoding/json"
	"net/http"

	"github.com/shinshin8/myFavorite_backend/dto"
	"github.com/shinshin8/myFavorite_backend/model"
	"github.com/shinshin8/myFavorite_backend/utils"
)

// PostList is controller file for get all post with JSON format.
func PostList(w http.ResponseWriter, r *http.Request) {
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
		// DB result array
		dbResultArray := model.GetPosts()

		resStruct := dto.PostList{
			Status:    true,
			UserID:    0,
			ErrorCode: utils.SuccessCode,
			Posts:     dbResultArray,
		}

		res, err := json.Marshal(resStruct)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Response JSON
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		// DB result array
		dbResultArray := model.GetPosts()

		resStruct := dto.PostList{
			Status:    true,
			UserID:    userID,
			ErrorCode: utils.SuccessCode,
			Posts:     dbResultArray,
		}

		res, err := json.Marshal(resStruct)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Response JSON
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
