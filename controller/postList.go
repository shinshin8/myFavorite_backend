package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/shinshin8/myFavorite/dto"
	"github.com/shinshin8/myFavorite/model"
	"github.com/shinshin8/myFavorite/utils"
)

// PostList is controller file for get all post with JSON format.
func PostList(w http.ResponseWriter, r *http.Request) {
	// listening port
	var localHostConfig dto.IPAddressConfig
	// decoding toml
	_, err := toml.DecodeFile(utils.ConfigFile, &localHostConfig)
	if err != nil {
		fmt.Println(err)
	}
	ipAddress := localHostConfig.IPAddress
	// Set CORS
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	w.Header().Set(utils.Cors, ipAddress)
	w.Header().Set(utils.ArrowHeader, utils.ContentType)
	w.Header().Set(utils.Credential, utils.True)

	successfulCode := 0
	// DB result array
	dbResultArray := model.GetPosts()

	resStruct := dto.PostList{
		Status:    true,
		ErrorCode: successfulCode,
		Posts:     dbResultArray,
	}

	res, err := json.Marshal(resStruct)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Response JSON
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
