/*
	This file is the entry point of this application's back-end.
	The file defines routes and HTTP method.
*/
package main

import (
	"fmt"
	"net/http"

	"./controller"
	"./dto"
	"./utils"
	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
)

var portConfig dto.PortConfig

func main() {

	// decoding toml
	_, err := toml.DecodeFile(utils.ConfigFile, &portConfig)
	if err != nil {
		fmt.Println(err)
	}

	// initialize router
	r := mux.NewRouter()

	// Login
	r.HandleFunc(utils.LoginPath, controller.Login)

	// Sign-In
	r.HandleFunc(utils.SignInPath, controller.SignUp)

	// Post list
	r.HandleFunc(utils.PostList, controller.PostList)

	// Show liked post lists
	r.HandleFunc(utils.ShowLikedPosts, controller.ShowLikedPosts).Methods(utils.Get)

	// Like post
	r.HandleFunc(utils.LikePost, controller.LikePost).Methods(utils.Post)

	// listening port
	port := portConfig.Port.Port
	// listenre
	http.ListenAndServe(port, r)
}
