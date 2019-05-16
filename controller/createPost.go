package controller

import (
	"net/http"
	"strconv"

	"../utils"
)

// CreatePost creates a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
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

	}

	// Check title
	if !utils.IsTitle(title) {

	}

	// Check content
	if !utils.IsContent(content) {

	}
}
