package controller

import (
	"encoding/json"
	"net/http"

	"github.com/shinshin8/myFavorite/dto"
	"github.com/shinshin8/myFavorite/utils"
)

// Logout delete session and let users logout.
func Logout(w http.ResponseWriter, r *http.Request) {
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

	_, err = utils.Cache.Do(utils.SessionDelete, sessionToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	successfulLoginCode := 0
	// set values in structs
	resultjson := dto.SimpleResutlJSON{
		Status:    http.StatusOK,
		ErrorCode: successfulLoginCode,
	}
	// convert structs to json
	res, err := json.Marshal(resultjson)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}
