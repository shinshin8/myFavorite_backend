package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/shinshin8/myFavorite/dto"
	"github.com/shinshin8/myFavorite/model"
	"github.com/shinshin8/myFavorite/utils"
)

// LikePost likes posts when user is login.
func LikePost(w http.ResponseWriter, r *http.Request) {
	// Set CORS
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	w.Header().Set(utils.Cors, utils.CorsWildCard)
	w.Header().Set(utils.ArrowHeader, utils.ContentType)
	w.Header().Set(utils.Credential, utils.True)

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

	atcID := "article_id"
	articleIDStr := r.URL.Query().Get(atcID)
	articleID, _ := strconv.Atoi(articleIDStr)

	// Execute register liked post
	res := model.LikePost(userID, articleID)

	if res {
		successfulLoginCode := 0
		// set values in structs
		resultjson := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: successfulLoginCode,
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
		failedLoginCode := 10
		// set values in structs
		resultjson := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: failedLoginCode,
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
