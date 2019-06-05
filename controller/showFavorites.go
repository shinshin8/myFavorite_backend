package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/gomodule/redigo/redis"
	"github.com/shinshin8/myFavorite/dto"
	"github.com/shinshin8/myFavorite/model"
	"github.com/shinshin8/myFavorite/utils"
)

// ShowFavoritePosts returns the list of favorite posts
func ShowFavoritePosts(w http.ResponseWriter, r *http.Request) {
	// listening port
	var localHostConfig dto.IPAddressConfig
	// decoding toml
	_, err := toml.DecodeFile(utils.ConfigFile, &localHostConfig)
	if err != nil {
		fmt.Println(err)
	}
	ipAddress := localHostConfig.IPAddress
	// Set CORS
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	w.Header().Set(utils.Cors, ipAddress)
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
	// Get article id from the URL query parameter in string type and conver it to int type.
	atlID := "article_id"
	articleIDStr := r.URL.Query().Get(atlID)
	articleID, _ := strconv.Atoi(articleIDStr)

	// Execute get all favorite posts list and response it in JSON.
	successfulCode := 0
	favoritePosts := model.ShowFavoritePosts(userID, articleID)

	resStruct := dto.PostList{
		Status:    true,
		ErrorCode: successfulCode,
		Posts:     favoritePosts,
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
