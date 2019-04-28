package controller

import (
	"fmt"
	"net/http"
	"encoding/json"
	"../utils"
	"../dto"
	"../model"
)

// signUpResultJSON struct
type signUpResultJSON struct {
	Status    int `json:"status"`
	ErrorCode int `json:"error_code"`
	Username string `json:username`
	EmailAddress string `json:email_address`
	Password string `json:password`
	ConfirmPassword string `json:confrim_password`
}

// SignUP returns the sign up results in JSON.
func SignUp(w http.ResponseWriter, r *http.Request){
	// Input form name
	var (
		username = "username"
		emailAddress = "emailAddress"
		password = "password"
		confirmPassword = "confirmPassword"
	)

	// HTTP method 
	var post = "POST"

	// HTTP header information
	var (
		contentType     = "Content-Type"
		applicationJSON = "application/json"
	)

	// Username value
	username := r.PostFormValue(username)
	// Email address value
	emailAddress := r.PostFormValue(emailAddress)
	// Password value
	password := r.PostFormValue(password)
	// Confirm password value
	confirmPassword := r.PostFormValue(confirmPassword)

	r.Method == post{

		// Validation check for username.
		if !utils.IsName(username) {
			// Invalid username
			invalidUsername := 3
			// Set values into the struct
			resStruct := signUpResult(http.StatusOK, invalidUsername, username, emailAddress, password, confirmPassword)
			// convert struct to JSON
			res, err := json.Marshal(resStruct)
			
			if err != nil{
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Set HTTP header and defined MIME type
			w.Header().Set(contentType, applicationJSON)
			// Response JSON
			w.Write(res)
			
		}

		// Validation check for email address
		if !utils.IsEmailAddress(emailAddress) {
			// Invalid emailAddress 
			invalidMailAddress := 4
			// Set values into the struct
			resStruct := signUpResult(http.StatusOK, invalidMailAddress, username, emailAddress, password, confirmPassword)
			// convert struct to JSON
			res, err := json.Marshal(resStruct)

			if err != nil{
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Set HTTP header and defined MIME type
			w.Header().Set(contentType, applicationJSON)
			// Response JSON
			w.Write(res)
		}
		
		// Validation check for password
		if !utils.IsPassword(password) {
			// Invalid password
			invalidPassword := 5
			// Set values into the struct
			resStruct := signUpResult(http.StatusOK, invalidPassword, username, emailAddress, password, confirmPassword)

			// convert struct to JSON
			res, err := json.Marshal(resStruct)
			
			if err != nil{
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Set HTTP header and defined MIME type
			w.Header().Set(contentType, applicationJSON)
			// Response JSON
			w.Write(res)
		}
		
		// Check whether or not the both values: password and confrim password are equal.
		if password != confirmPassword {
			// Password and confirm password don't match.
			notMatchPasswords := 6
			// Set values into the struct
			resStruct := signUpResult(http.StatusOK, notMatchPasswords, username, emailAddress, password, confirmPassword)

			// convert struct to JSON
			res, err := json.Marshal(resStruct)
			
			if err != nil{
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Set HTTP header and defined MIME type
			w.Header().Set(contentType, applicationJSON)
			// Response JSON
			w.Write(res)
		}
		// Create Sign-up DTO
		signupDTO := dto.SignUpInfo(username, emailAddress, password)

		// TODO:Insert sign up information in database
		// In this time, method returns only int; error_code.
		signUpRes := model.SignUp(signupDTO)
		// Set values into the struct
		resStruct := signUpResult(http.StatusOK, signUpRes, username, emailAddress, password, confirmPassword)

		// convert struct to JSON
		res, err := json.Marshal(resStruct)
		
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Set HTTP header and defined MIME type
		w.Header().Set(contentType, applicationJSON)
		// Response JSON
		w.Write(res)
	}else{
		// Wrong HTTP request method
		wrongHTTPMethod := 7
		// Set values into the struct
		resStruct := signUpResult(http.StatusOK, wrongHTTPMethod, username, emailAddress, password, confirmPassword)

		// convert struct to JSON
		res, err := json.Marshal(resStruct)
		
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Set HTTP header and defined MIME type
		w.Header().Set(contentType, applicationJSON)
		// Response JSON
		w.Write(res)
	}
}

// signUpResult returns the pointer of struct; signUpResultJSON.
// Status is in the first paramter.
// ErrorCode is in the second parameter.
// Username is in third parameter.
// EmailAddress is in the forth parameter.
// Password is in the fifth parameter.
// ConfirmPassword is in the sixth parameter.
func signUpResult(status string, errorCode int, username string, emailAddress string, password string, confirmPassword string) *signUpResultJSON {
	res := new(signUpResultJSON)
	res.Status = status
	res.ErrorCode = errorCode
	res.Username = username
	res.EmailAddress = emailAddress
	res.Password = password
	res.ConfirmPassword = confirmPassword
	return res
}