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

// SignUp returns the sign up results in JSON.
func SignUp(w http.ResponseWriter, r *http.Request) {
	// Set CORS
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	w.Header().Set(utils.Cors, utils.CorsWildCard)
	w.Header().Set(utils.ArrowHeader, utils.ContentType)
	w.Header().Set(utils.Credential, utils.True)

	var signUpBody dto.SingUpBody

	err := json.NewDecoder(r.Body).Decode(&signUpBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Username value
	username := signUpBody.UserName
	// Email address value
	emailAddress := signUpBody.Email
	// Password value
	password := signUpBody.Password
	// Confirm password value
	confirmPassword := signUpBody.ConfirmPassword

	// Validation check for username.
	if !utils.IsName(username) {
		// Invalid username
		invalidUsername := 3
		// Set values into the struct
		resStruct := dto.SignUpResult(false, invalidUsername, username, emailAddress)
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

	// Validation check for email address
	if !utils.IsEmailAddress(emailAddress) {
		// Invalid emailAddress
		invalidMailAddress := 4
		// Set values into the struct
		resStruct := dto.SignUpResult(false, invalidMailAddress, username, emailAddress)
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

	// Validation check for password
	if !utils.IsPassword(password) {
		// Invalid password
		invalidPassword := 5
		// Set values into the struct
		resStruct := dto.SignUpResult(false, invalidPassword, username, emailAddress)

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

	// Check whether or not the both values: password and confrim password are equal.
	if password != confirmPassword {
		// Password and confirm password don't match.
		notMatchPasswords := 6
		// Set values into the struct
		resStruct := dto.SignUpResult(false, notMatchPasswords, username, emailAddress)

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

	// Hashing password
	hash := sha256.New()
	hash.Write([]byte(password))
	hexPassword := hash.Sum(nil)
	hashedPassword := hex.EncodeToString(hexPassword)

	// In this time, method returns only int; error_code.
	signUpRes := model.SignUp(username, emailAddress, hashedPassword)

	// Create a new session token.
	sessionToken := uuid.NewV4().String()
	// Set session in the cache.
	// Token will expire in 1200 seconds.
	_, er := utils.Cache.Do(utils.SessionSet, sessionToken, utils.SessionTimeOut, signUpRes)

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

	// convert struct to JSON
	res, err := json.Marshal(resultjson)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Response JSON
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
