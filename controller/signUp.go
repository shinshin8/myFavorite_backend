/*
	signUp.go is controller for sign-up manipulation.
*/
package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"../dto"
	"../model"
	"../utils"
)

// SignUp returns the sign up results in JSON.
func SignUp(w http.ResponseWriter, r *http.Request) {
	// Input form name
	var (
		username        = "username"
		emailAddress    = "emailAddress"
		password        = "password"
		confirmPassword = "confirmPassword"
	)

	if r.Method == utils.Post {

		// Username value
		username := r.PostFormValue(username)
		// Email address value
		emailAddress := r.PostFormValue(emailAddress)
		// Password value
		password := r.PostFormValue(password)
		// Confirm password value
		confirmPassword := r.PostFormValue(confirmPassword)

		// Validation check for username.
		if !utils.IsName(username) {
			// Invalid username
			invalidUsername := 3
			// Set values into the struct
			resStruct := dto.SignUpResult(http.StatusOK, invalidUsername, username, emailAddress)
			// convert struct to JSON
			res, err := json.Marshal(resStruct)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Set HTTP header and defined MIME type
			w.Header().Set(utils.ContentType, utils.ApplicationJSON)
			// Response JSON
			w.Write(res)
			return

		}

		// Validation check for email address
		if !utils.IsEmailAddress(emailAddress) {
			// Invalid emailAddress
			invalidMailAddress := 4
			// Set values into the struct
			resStruct := dto.SignUpResult(http.StatusOK, invalidMailAddress, username, emailAddress)
			// convert struct to JSON
			res, err := json.Marshal(resStruct)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Set HTTP header and defined MIME type
			w.Header().Set(utils.ContentType, utils.ApplicationJSON)
			// Response JSON
			w.Write(res)
			return
		}

		// Validation check for password
		if !utils.IsPassword(password) {
			// Invalid password
			invalidPassword := 5
			// Set values into the struct
			resStruct := dto.SignUpResult(http.StatusOK, invalidPassword, username, emailAddress)

			// convert struct to JSON
			res, err := json.Marshal(resStruct)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Set HTTP header and defined MIME type
			w.Header().Set(utils.ContentType, utils.ApplicationJSON)
			// Response JSON
			w.Write(res)
			return
		}

		// Check whether or not the both values: password and confrim password are equal.
		if password != confirmPassword {
			// Password and confirm password don't match.
			notMatchPasswords := 6
			// Set values into the struct
			resStruct := dto.SignUpResult(http.StatusOK, notMatchPasswords, username, emailAddress)

			// convert struct to JSON
			res, err := json.Marshal(resStruct)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Set HTTP header and defined MIME type
			w.Header().Set(utils.ContentType, utils.ApplicationJSON)
			// Response JSON
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

		// In the Model, the function returns JSON in other way.
		// So in this part, just response result.

		// convert struct to JSON
		res, err := json.Marshal(signUpRes)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Set HTTP header and defined MIME type
		w.Header().Set(utils.ContentType, utils.ApplicationJSON)
		// Response JSON
		w.Write(res)
	} else {
		// Wrong HTTP request method
		wrongHTTPMethod := 7
		// Set values into the struct
		resStruct := dto.SignUpResult(http.StatusOK, wrongHTTPMethod, username, emailAddress)

		// convert struct to JSON
		res, err := json.Marshal(resStruct)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Set HTTP header and defined MIME type
		w.Header().Set(utils.ContentType, utils.ApplicationJSON)
		// Response JSON
		w.Write(res)
	}
}
