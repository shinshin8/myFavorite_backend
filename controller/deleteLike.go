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

// DeleteLikedPost deals with deleting resouce and return its result in JSON format.
func DeleteLikedPost(w http.ResponseWriter, r *http.Request) {
	// Set CORS
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	w.Header().Set(utils.Cors, utils.CorsWildCard)
	w.Header().Set(utils.ArrowHeader, utils.ContentType)
	w.Header().Set(utils.ArrowMethods, utils.Methods)
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
	// Get article id from URL query parameter with string type and convert it to int.
	atlID := "article_id"
	atlIDStr := r.URL.Query().Get(atlID)
	articleID, _ := strconv.Atoi(atlIDStr)

	// Execute delete resouce.
	res := model.DeleteLikedPost(userID, articleID)

	if res {
		successfulLoginCode := 0
		// set values in structs
		resultjson := dto.SimpleResutlJSON{
			Status:    true,
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
		failedCode := 11
		// set values in structs
		resultjson := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: failedCode,
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
