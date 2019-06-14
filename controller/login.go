package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/shinshin8/myFavorite_backend/dto"
	"github.com/shinshin8/myFavorite_backend/model"
	"github.com/shinshin8/myFavorite_backend/utils"
)

// Login function
func Login(w http.ResponseWriter, r *http.Request) {
	// Set CORS
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	w.Header().Set(utils.Cors, utils.CorsWildCard)
	w.Header().Set(utils.ArrowHeader, utils.ContentType)
	w.Header().Set(utils.Credential, utils.True)

	var loginBody dto.LoginBody

	err := json.NewDecoder(r.Body).Decode(&loginBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// analyze request form
	// get username
	username := loginBody.UserName
	// get password
	password := loginBody.Password

	// hashed password
	hash := sha256.New()
	hash.Write([]byte(password))
	hexPassword := hash.Sum(nil)
	hashedPassword := hex.EncodeToString(hexPassword)

	// get login result
	loginRes := model.LoginUser(username, hashedPassword)

	// Create a new session token.
	sessionToken := uuid.NewV4().String()
	// Set session in the cache.
	// Token will expire in 1200 seconds.
	_, er := utils.Cache.Do(utils.SessionSet, sessionToken, utils.SessionTimeOut, loginRes)

	if er != nil {
		// return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Set client cookie
	http.SetCookie(w, &http.Cookie{
		Name:    utils.CookieName,
		Value:   sessionToken,
		Expires: time.Now().Add(utils.SessionExpire * time.Second),
	})

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
}
