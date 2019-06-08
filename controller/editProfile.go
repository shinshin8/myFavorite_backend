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
	// Set CORS
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	w.Header().Set(utils.Cors, utils.CorsWildCard)
	w.Header().Set(utils.ArrowHeader, utils.ContentType)
	w.Header().Set(utils.Credential, utils.True)
	// Session
	c, err := r.Cookie(utils.CookieName)
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
	userIDCache, err := utils.Cache.Do(utils.SessionGet, sessionToken)
	userID, _ := redis.Int(userIDCache, err)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if userIDCache == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var editProfileBody dto.EditProfileBody

	er := json.NewDecoder(r.Body).Decode(&editProfileBody)

	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get user name
	userName := editProfileBody.UserName
	// Get birthday
	birthday := editProfileBody.Birthday
	//Get mail address
	mailAddress := editProfileBody.MailAddress
	// Get comment
	comment := editProfileBody.Comment

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
			Status:    false,
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
		w.WriteHeader(http.StatusOK)
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
			Status:    false,
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
		w.WriteHeader(http.StatusOK)
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
			Status:    false,
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
		w.WriteHeader(http.StatusOK)
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
			Status:    false,
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
		w.WriteHeader(http.StatusOK)
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
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
