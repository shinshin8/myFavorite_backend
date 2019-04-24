/*
	This file handling login function.
*/
package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"../model"
)

// Login result json
type resultJSON struct {
	Status    int `json:"status"`
	ErrorCode int `json:"error_code"`
}

// Login function
func Login(w http.ResponseWriter, r *http.Request) {

	// input form name
	var (
		username = "username"
		password = "password"
	)

	// HTTP method
	var post = "POST"

	// header information
	var (
		contentType     = "Content-Type"
		applicationJSON = "application/json"
	)

	// judge http method
	if r.Method == post {
		// analyze request form
		// get username
		username := r.PostFormValue(username)
		// get password
		password := r.PostFormValue(password)

		// hashed password
		hash := sha256.New()
		hash.Write([]byte(password))
		hexPassword := hash.Sum(nil)
		hashedPassword := hex.EncodeToString(hexPassword)

		// get login result
		loginBooleanResult := model.LoginUser(username, hashedPassword)

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
