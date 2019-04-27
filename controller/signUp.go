package controller

import (
	"fmt"
	"net/http"
)

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

	r.Method == post{
		// Get username value
		username := r.PostFormValue(username)
		// Get email address value
		emailAddress := r.PostFormValue(emailAddress)
		// Get password value
		password := r.PostFormValue(password)
		// Get confirm password value
		confirmPassword := r.PostFormValue(confirmPassword)

		// TODO: Create Util package and validation file.
		// Validation check for username.
		// Validation check for email address
		// Validation check for password
		// Validation check for confirm password
		// Check whether or not the both values: password and confrim password are equal.
		// 
	}else{
		
	}
}