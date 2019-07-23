package controller

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/shinshin8/myFavorite_backend/dto"
	"github.com/shinshin8/myFavorite_backend/model"
	"github.com/shinshin8/myFavorite_backend/utils"
)

// ShowFavoritePosts returns the list of favorite posts
func ShowFavoritePosts(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(res)
		return
	}

	favoritePosts := model.ShowFavoritePosts(userID)
	// Get images
	favoriteImages := model.GetFavoriteListImage(userID)
	// Get Icon
	iconDataArray := model.GetAllIcon()
	// Struct for Posts data
	var postsData []dto.Posts
	// Looping article data
	for _, article := range favoritePosts {
		var image string
		for _, imageData := range favoriteImages {
			if len(imageData.ImageURL) >= 1 {
				image = os.Getenv("S3_URL") + imageData.ImageURL[0]
			} else {
				image = ""
			}
		}
		var icon string
		for _, eachIcon := range iconDataArray {
			if eachIcon.UserID == article.UserID {
				icon = os.Getenv("S3_URL") + eachIcon.ImageURL
			} else {
				icon = ""
			}
		}
		post := dto.Posts{
			ArticleID:   article.ArticleID,
			LikedSum:    article.LikedSum,
			ImageURL:    image,
			IconURL:     icon,
			UserName:    article.UserName,
			Title:       article.Title,
			CreatedTime: article.CreatedTime,
		}
		postsData = append(postsData, post)
	}

	resStruct := dto.PostList{
		Status:    true,
		ErrorCode: utils.SuccessCode,
		UserID:    userID,
		Posts:     postsData,
	}

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
