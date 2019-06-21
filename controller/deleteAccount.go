package controller

import (
	"encoding/json"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/shinshin8/myFavorite_backend/dto"
	"github.com/shinshin8/myFavorite_backend/model"
	"github.com/shinshin8/myFavorite_backend/utils"
)

// DeleteAccount delete loginned user's account.
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	// Set CORS
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	w.Header().Set(utils.Cors, utils.CorsWildCard)
	w.Header().Set(utils.ArrowHeader, utils.ContentType)
	w.Header().Set(utils.ArrowMethods, utils.Methods)
	w.Header().Set(utils.Credential, utils.True)
	// Get jwt from header.
	reqToken := r.Header.Get(utils.Authorization)
	// Check if jwt is verified.
	token, _ := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("boobar"), nil
	})
	if token == nil {
		// set values in structs
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
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
	// Get user id from jwt.
	claims := token.Claims.(jwt.MapClaims)
	userIDkey := "user_id"
	userIDFloat64, _ := claims[userIDkey]
	userID := int(userIDFloat64.(float64))

	// Execute delete user's account
	deleteAccount := model.DeleteAccount(userID)

	if deleteAccount {
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
			ErrorCode: utils.FailedDeleteAccount,
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
