package controller

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/shinshin8/myFavorite_backend/dto"
	"github.com/shinshin8/myFavorite_backend/model"
	"github.com/shinshin8/myFavorite_backend/utils"
)

// Timeline is controller file for get all post with JSON format.
func Timeline(w http.ResponseWriter, r *http.Request) {
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
	// Get article data
	articleDataArray := model.Timeline()
	// Get imgage Data
	imageDataArray := model.GetAllImages()
	// Struct for Posts data
	var postsData []dto.Posts
	// Looping article data
	for _, article := range articleDataArray {
		for _, imageData := range imageDataArray {
			if article.ArticleID == imageData.ArticleID {
				var imageArray []string
				imageArray = append(imageArray, os.Getenv("S3_URL")+imageData.ImageURL)
				var image []string
				firstImage := imageArray[0]
				image = append(image, firstImage)
				post := dto.Posts{
					ArticleID:    article.ArticleID,
					LikedSum:     article.LikedSum,
					ImageURL:     image,
					UserName:     article.UserName,
					Title:        article.Title,
					Content:      article.Content,
					CreatedTime:  article.CreatedTime,
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
