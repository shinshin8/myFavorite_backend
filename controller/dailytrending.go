package controller

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/shinshin8/myFavorite_backend/dto"
	"github.com/shinshin8/myFavorite_backend/model"
	"github.com/shinshin8/myFavorite_backend/utils"
)

// Trending returns trending posts in a list.
func Trending(w http.ResponseWriter, r *http.Request) {
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
	// Get trending data
	articleDataArray := model.Trending()
	// Get imgage Data
	imageDataArray := model.GetAllImages()
	// Get Icon
	iconDataArray := model.GetAllIcon()
	// Struct for Posts data
	var postsData []dto.Posts
	// Looping article data
	for _, article := range articleDataArray {
		for _, imageData := range imageDataArray {
			if article.ArticleID == imageData.ArticleID {
				var image []string
				if len(imageData.ImageURL) >= 1 {
					var imageArray []string
					imageArray = append(imageArray, os.Getenv("S3_URL")+imageData.ImageURL)
					firstImage := imageArray[0]
					image = append(image, firstImage)
				}
				var icon []string
				for _, eachIcon := range iconDataArray {
					if eachIcon.UserID == article.UserID {
						icon = append(icon, os.Getenv("S3_URL")+eachIcon.ImageURL)
					}
				}
				post := dto.Posts{
					ArticleID:    article.ArticleID,
					LikedSum:     article.LikedSum,
					ImageURL:     image[0],
					IconURL:      icon[0],
					UserName:     article.UserName,
					Title:        article.Title,
					ModifiedTime: article.ModifiedTime,
				}
				postsData = append(postsData, post)
			}
		}
	}

	if userID == 0 {

		resStruct := dto.PostList{
			Status:    true,
			UserID:    0,
			ErrorCode: utils.SuccessCode,
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

	resStruct := dto.PostList{
		Status:    true,
		UserID:    userID,
		ErrorCode: utils.SuccessCode,
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
