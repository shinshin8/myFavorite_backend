package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../dto"
	"../model"
	"../utils"
)

// IndividualPost returns an individual post in JSON 
func IndividualPost(w http.ResponseWriter, r *http.Request){
	// Get user id from the URL query paramter in string type and conver it to int type.
	usrID := "user_id"
	userIDStr := r.URL.Query().Get(usrID)
	userID, _ := strconv.Atoi(userIDStr)
	// Get article id from the URL query parameter in string type and conver it to int type.
	atlID := "article_id"
	articleIDStr := r.URL.Query().Get(atlID)
	articleID, _ := strconv.Atoi(articleIDStr)

	// Get a result if this post is liked by a user.
	// this method returns the result in boolean.
	// If the result is true, it means this post is liked.
	// Otherwise, it means this post is not liked.
	likedResult := model.LikedOrNot(userID, articleID)
	// Get a result if this post is favorited by a user.
	// this method returns the result in boolean.
	// If the result is true, it means this post is favorited.
	// Otherwise, it means this post is not favorited.
	favoriteResult := model.FavoriteOrNot(userID, articleID)
	// Post result
	singlePost := model.IndividualPost(userID, articleID)
	successfulCode := 0

	resStruct := dto.SiglePost{
		Status       http.StatusOK,
		ErrorCode    successfulCode,
		LikedFlg     likedResult,
		FavoriteFlg  favoriteResult,
		ArticleID    singlePost.ArticleID,
		UserName     singlePost.UserName,
		Title        singlePost.Title,
		Content      singlePost.Content,
		CreatedTime  singlePost.CreatedTime,
		ModifiedTime singlePost.ModifiedTime,
	}

	res, err := json.Marshal(resStruct)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Set HTTP header and defined MIME type
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	// Response JSON
	w.Write(res)

}