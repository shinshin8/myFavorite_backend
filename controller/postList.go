package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gomodule/redigo/redis"
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
	w.Header().Set(utils.Credential, utils.True)

	// Session
	c, err := r.Cookie(utils.CookieName)

	// In case user is not logged in.
	if c == nil {
		successfulCode := 0
		// DB result array
		dbResultArray := model.GetPosts()

		resStruct := dto.PostList{
			Status:    true,
			UserID:    0,
			ErrorCode: successfulCode,
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
		// In case user is logged in.
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

		successfulCode := 0
		// DB result array
		dbResultArray := model.GetPosts()

		resStruct := dto.PostList{
			Status:    true,
			UserID:    userID,
			ErrorCode: successfulCode,
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
