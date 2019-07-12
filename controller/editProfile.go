package controller

import (
	"encoding/json"
	"net/http"

	"github.com/shinshin8/myFavorite_backend/dto"
	"github.com/shinshin8/myFavorite_backend/model"
	"github.com/shinshin8/myFavorite_backend/utils"
)

// EditProfile edits user's profile.
func EditProfile(w http.ResponseWriter, r *http.Request) {
	// Set CORS
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	w.Header().Set(utils.Cors, utils.CorsWildCard)
	w.Header().Set(utils.ArrowHeader, utils.ContentType)
	w.Header().Set(utils.ArrowMethods, utils.Methods)
	w.Header().Set(utils.Credential, utils.True)
	// Get jwt from header.
	reqToken := r.Header.Get(utils.Authorization)
	// Check if jwt is verified.
	userID := utils.VerifyToken(reqToken)
	if userID == 0 {
		resultjson := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: utils.InvalidToken,
		}
		// convert structs to json
		res, err := json.Marshal(resultjson)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(res)
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
		profile := dto.Profile{
			UserID:      userID,
			Birthday:    birthday,
			MailAddress: mailAddress,
			Comment:     comment,
		}
		// Set values into the struct
		resStruct := dto.ProfileResult{
			Status:    false,
			ErrorCode: utils.InvalidEditProfileUserName,
			Profile:   profile,
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
	// Check birthday
	if !utils.IsBirthday(birthday) {
		// Set values into the struct
		profile := dto.Profile{
			UserID:      userID,
			Birthday:    birthday,
			MailAddress: mailAddress,
			Comment:     comment,
		}
		resStruct := dto.ProfileResult{
			Status:    false,
			ErrorCode: utils.InvalidEditProfileBirthday,
			Profile:   profile,
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
	// Check mail address
	if !utils.IsEmailAddress(mailAddress) {
		// Set values into the struct
		profile := dto.Profile{
			UserID:      userID,
			Birthday:    birthday,
			MailAddress: mailAddress,
			Comment:     comment,
		}
		resStruct := dto.ProfileResult{
			Status:    false,
			ErrorCode: utils.InvalidEditProfileMailAddress,
			Profile:   profile,
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
	// Check comment
	if !utils.IsComment(comment) {
		// Set values into the struct
		profile := dto.Profile{
			UserID:      userID,
			Birthday:    birthday,
			MailAddress: mailAddress,
			Comment:     comment,
		}
		resStruct := dto.ProfileResult{
			Status:    false,
			ErrorCode: utils.InvalidEditProfileComment,
			Profile:   profile,
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

	// Execute edit user's profile.
	result := model.EditProfile(userID, userName, birthday, mailAddress, comment)

	if result {
		// set values in structs
		resultjson := dto.SimpleResutlJSON{
			Status:    true,
			ErrorCode: utils.SuccessCode,
		}
		// convert structs to json
		res, err := json.Marshal(resultjson)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		// set values in structs
		resultjson := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: utils.FailedEditProfile,
		}
		// convert structs to json
		res, err := json.Marshal(resultjson)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
	}
}
