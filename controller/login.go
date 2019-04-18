/*
	This file handling login function.
*/
package controller

import (
	"crypto/sha256"
	"encoding/json"
	"net/http"
)

// HTTP method
var post = "POST"

// input form name
var (
	username = "username"
	password = "password"
)

// header
var (
	contentType     = "Content-Type"
	applicationJSON = "application/json"
)

// Login result json
type resultJSON struct {
	Status    int `json:"status"`
	ErrorCode int `json:"error_code"`
}

// Login function
func Login(w http.ResponseWriter, r *http.Request) {
	// judge http method
	if r.Method == post {
		// analyze request form
		r.ParseForm()
		// get username
		username := r.PostFormValue(username)
		// get password
		password := r.PostFormValue(password)
		// hashed password
		hashedPassword := sha256.Sum256([]byte(password))
		// get login result
		loginBooleanResult := LoginUser(username, hashed_password)

		// response json
		if loginBooleanResult == true {
			successfulLoginCode := 0
			// set values in structs
			resultjson := resultJSON{
				Status:    http.StatusOK,
				ErrorCode: successfulLoginCode,
			}
			// convert structs to json
			res, err := json.Marshal(resultjson)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// set header and defined response type for json
			w.Header().Set(contentType, applicationJSON)
			w.Write(res)
		} else {
			failedLoginCode := 1
			// set values in structs
			resultjson := resultJSON{
				Status:    http.StatusOK,
				ErrorCode: failedLoginCode,
			}
			// convert structs to json
			res, err := json.Marshal(resultjson)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// set header and defined response type for json
			w.Header().Set(contentType, applicationJSON)
			w.Write(res)
		}
	} else {
		wrongHTTPMethodCode := 2
		// set values in structs
		resultjson := resultJSON{
			Status:    http.StatusNotFound,
			ErrorCode: wrongHTTPMethodCode,
		}
		// convert structs to json
		res, err := json.Marshal(resultjson)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// set header and defined response type for json
		w.Header().Set(contentType, applicationJSON)
		w.Write(res)
	}
}
