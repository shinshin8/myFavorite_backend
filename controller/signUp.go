package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"

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
	w.Header().Set(utils.ArrowMethods, utils.Methods)
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
	emailAddress := signUpBody.EmailAddress
	// Password value
	password := signUpBody.Password
	// Confirm password value
	confirmPassword := signUpBody.ConfirmPassword
	// Validation check for username.
	if !utils.IsName(username) {
		// Set values into the struct
		resStruct := dto.SignUpResultJSON{
			Status:       false,
			ErrorCode:    utils.InvalidSignUpUsername,
			Username:     username,
			EmailAddress: emailAddress,
		}
		// convert struct to JSON
		res, err := json.Marshal(resStruct)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Response JSON
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}
	// Validation check for email address
	if !utils.IsEmailAddress(emailAddress) {
		// Set values into the struct
		resStruct := dto.SignUpResultJSON{
			Status:       false,
			ErrorCode:    utils.InvalidSignUpMailAddress,
			Username:     username,
			EmailAddress: emailAddress,
		}
		// convert struct to JSON
		res, err := json.Marshal(resStruct)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Response JSON
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}
	// Validation check for password
	if !utils.IsPassword(password) {
		// Set values into the struct
		resStruct := dto.SignUpResultJSON{
			Status:       false,
			ErrorCode:    utils.InvalidSignUpPassword,
			Username:     username,
			EmailAddress: emailAddress,
		}
		// convert struct to JSON
		res, err := json.Marshal(resStruct)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Response JSON
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}
	// Check whether or not the both values: password and confrim password are equal.
	if password != confirmPassword {
		// Set values into the struct
		resStruct := dto.SignUpResultJSON{
			Status:       false,
			ErrorCode:    utils.NotMatchPasswords,
			Username:     username,
			EmailAddress: emailAddress,
		}
		// convert struct to JSON
		res, err := json.Marshal(resStruct)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Response JSON
		w.WriteHeader(http.StatusBadRequest)
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
	// Creating jwt
	token := utils.CreateToken(signUpRes)
	resultjson := dto.LoginResult{
		Status:    true,
		ErrorCode: utils.SuccessCode,
		Token:     token,
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
	return
}
