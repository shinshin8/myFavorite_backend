package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../dto"
	"../model"
	"../utils"
)

// EditProfile edits user's profile.
func EditProfile(w http.ResponseWriter, r *http.Request) {
	// Get user id from URL query parameter and convert its type string to int.
	usrID := "user_id"
	userIDStr := r.URL.Query().Get(usrID)
	userID, _ := strconv.Atoi(userIDStr)
	// Input forms
	var (
		userNameForm    = "user_name"
		birthdayForm    = "birthday"
		mailAddressForm = "mailAddress"
		commentForm     = "comment"
	)

	// Get user name
	userName := r.PostFormValue(userNameForm)
	// Get birthday
	birthday := r.PostFormValue(birthdayForm)
	//Get mail address
	mailAddress := r.PostFormValue(mailAddressForm)
	// Get comment
	comment := r.PostFormValue(commentForm)

	// Check user name
	if !utils.IsName(userName) {
		invalidUserName := 21
		// Set values into the struct
		resStruct := dto.Profile{
			Status:      http.StatusOK,
			ErrorCode:   invalidUserName,
			UserID:      userID,
			UserName:    userName,
			Birthday:    birthday,
			MailAddress: mailAddress,
			Comment:     comment,
		}
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
	// Check birthday
	if !utils.IsBirthday(birthday) {
		invalidBirthday := 22
		// Set values into the struct
		resStruct := dto.Profile{
			Status:      http.StatusOK,
			ErrorCode:   invalidBirthday,
			UserID:      userID,
			UserName:    userName,
			Birthday:    birthday,
			MailAddress: mailAddress,
			Comment:     comment,
		}
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

	// Check mail address
	if !utils.IsEmailAddress(mailAddress) {
		invalidMailAddress := 23
		// Set values into the struct
		resStruct := dto.Profile{
			Status:      http.StatusOK,
			ErrorCode:   invalidMailAddress,
			UserID:      userID,
			UserName:    userName,
			Birthday:    birthday,
			MailAddress: mailAddress,
			Comment:     comment,
		}
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
	// Check comment
	if !utils.IsComment(comment) {
		invalidComment := 24
		// Set values into the struct
		resStruct := dto.Profile{
			Status:      http.StatusOK,
			ErrorCode:   invalidComment,
			UserID:      userID,
			UserName:    userName,
			Birthday:    birthday,
			MailAddress: mailAddress,
			Comment:     comment,
		}
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

	// Execute edit user's profile.
	result := model.EditProfile(userID, userName, birthday, mailAddress, comment)

	// In the Model, the function returns JSON in other way.
	// So in this part, just response result.

	// convert struct to JSON
	res, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Set HTTP header and defined MIME type
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	// Response JSON
	w.Write(res)
}
