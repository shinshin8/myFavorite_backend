package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gomodule/redigo/redis"
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
	// Get article id from URL query parameter and convert its type string to int.
	atcID := "article_id"
	articleIDStr := r.URL.Query().Get(atcID)
	articleID, _ := strconv.Atoi(articleIDStr)

	// Session
	c, err := r.Cookie(utils.CookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sessionToken := c.Value

	// Get user id from cache.
	userIDCache, err := utils.Cache.Do(utils.SessionGet, sessionToken)
	userID, _ := redis.Int(userIDCache, err)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if userIDCache == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

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
	// Check userID
	if !utils.IsID(userID) {
		// Invalid user id
		invalidUserID := 17
		// Set values into the struct
		resStruct := dto.NewPost{
			Status:    false,
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
		w.WriteHeader(http.StatusOK)
		// Response JSON
		w.Write(res)
		return
	}

	// Check title
	if !utils.IsTitle(title) {
		// Invalid title
		invalidTitle := 18
		// Set values into the struct
		resStruct := dto.NewPost{
			Status:    false,
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
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	}

	// Check content
	if !utils.IsContent(content) {
		// Invalid content
		invalidContent := 19
		// Set values into the struct
		resStruct := dto.NewPost{
			Status:    false,
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
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	}

	// Execute update data to DB.
	result := model.EditPost(userID, articleID, title, content)

	// In the Model, the function returns JSON in other way.
	// So in this part, just response result.

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
