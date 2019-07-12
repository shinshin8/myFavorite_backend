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

// Login function
func Login(w http.ResponseWriter, r *http.Request) {
	// Set CORS
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	w.Header().Set(utils.Cors, "http://localhost:3000")
	w.Header().Set(utils.ArrowHeader, utils.ContentType)
	w.Header().Set(utils.ArrowMethods, utils.Methods)
	w.Header().Set(utils.Credential, utils.True)

	var loginBody dto.LoginBody

	err := json.NewDecoder(r.Body).Decode(&loginBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get username
	username := loginBody.UserName
	// get password
	password := loginBody.Password

	// Hashing password
	hash := sha256.New()
	hash.Write([]byte(password))
	hexPassword := hash.Sum(nil)
	hashedPassword := hex.EncodeToString(hexPassword)

	// get login result
	loginRes := model.LoginUser(username, hashedPassword)

	if loginRes == 0 {
		// set values in structs
		resultjson := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: utils.WrongUserNamePassword,
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
		// Creating jwt
		token := utils.CreateToken(loginRes)
		resultjson := dto.LoginResult{
			Status:    true,
			ErrorCode: utils.SuccessCode,
			Token:     token,
		}
		// convert structs to json
		res, err := json.Marshal(resultjson)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
