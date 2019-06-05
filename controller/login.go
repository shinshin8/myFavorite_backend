package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/shinshin8/myFavorite/dto"
	"github.com/shinshin8/myFavorite/model"
	"github.com/shinshin8/myFavorite/utils"
)

// Login function
func Login(w http.ResponseWriter, r *http.Request) {
	// listening port
	port := portConfig.Port.Port
	// Set CORS
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	w.Header().Set(utils.Cors, "http://localhost"+port)
	w.Header().Set(utils.ArrowHeader, utils.ContentType)
	w.Header().Set(utils.Credential, utils.True)

	// input form name
	var (
		usernm = "username"
		pwd    = "password"
	)

	// analyze request form
	// get username
	username := r.PostFormValue(usernm)
	// get password
	password := r.PostFormValue(pwd)

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
	// Token will expire in 300 seconds.
	_, err := utils.Cache.Do("SETEX", sessionToken, "300", loginRes.UserID)

	if err != nil {
		// return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set client cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(300 * time.Second),
	})

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
