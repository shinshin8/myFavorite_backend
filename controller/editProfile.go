package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/shinshin8/myFavorite/dto"
	"github.com/shinshin8/myFavorite/model"
	"github.com/shinshin8/myFavorite/utils"
)

// EditProfile edits user's profile.
func EditProfile(w http.ResponseWriter, r *http.Request) {
	// listening port
	port := portConfig.Port.Port
	// Set CORS
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	w.Header().Set(utils.Cors, "http://localhost"+port)
	w.Header().Set(utils.ArrowHeader, utils.ContentType)
	w.Header().Set(utils.Credential, utils.True)
	// Session
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sessionToken := c.Value

	// Get user id from cache.
	userIDCache, err := utils.Cache.Do("GET", sessionToken)
	userID, _ := redis.Int(userIDCache, err)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if userIDCache == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
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
		profile := dto.Profile{
			UserID:      userID,
			Birthday:    birthday,
			MailAddress: mailAddress,
			Comment:     comment,
		}
		// Set values into the struct
		resStruct := dto.ProfileResult{
			Status:    http.StatusOK,
			ErrorCode: invalidUserName,
			Profile:   profile,
		}

		// convert struct to JSON
		res, err := json.Marshal(resStruct)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Response JSON
		w.Write(res)
		return
	}
	// Check birthday
	if !utils.IsBirthday(birthday) {
		invalidBirthday := 22
		// Set values into the struct
		profile := dto.Profile{
			UserID:      userID,
			Birthday:    birthday,
			MailAddress: mailAddress,
			Comment:     comment,
		}
		resStruct := dto.ProfileResult{
			Status:    http.StatusOK,
			ErrorCode: invalidBirthday,
			Profile:   profile,
		}
		// convert struct to JSON
		res, err := json.Marshal(resStruct)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Response JSON
		w.Write(res)
		return
	}

	// Check mail address
	if !utils.IsEmailAddress(mailAddress) {
		invalidMailAddress := 23
		// Set values into the struct
		profile := dto.Profile{
			UserID:      userID,
			Birthday:    birthday,
			MailAddress: mailAddress,
			Comment:     comment,
		}
		resStruct := dto.ProfileResult{
			Status:    http.StatusOK,
			ErrorCode: invalidMailAddress,
			Profile:   profile,
		}
		// convert struct to JSON
		res, err := json.Marshal(resStruct)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Response JSON
		w.Write(res)
		return
	}
	// Check comment
	if !utils.IsComment(comment) {
		invalidComment := 24
		// Set values into the struct
		profile := dto.Profile{
			UserID:      userID,
			Birthday:    birthday,
			MailAddress: mailAddress,
			Comment:     comment,
		}
		resStruct := dto.ProfileResult{
			Status:    http.StatusOK,
			ErrorCode: invalidComment,
			Profile:   profile,
		}
		// convert struct to JSON
		res, err := json.Marshal(resStruct)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
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
	// Response JSON
	w.Write(res)
}
