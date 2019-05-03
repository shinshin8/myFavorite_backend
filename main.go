/*
	This file is the entry point of this application's back-end.
	The file defines routes and HTTP method.
*/
package main

import (
	"fmt"
	"net/http"

	"./controller"
	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
)

// configuration file
type Config struct {
	Port PortConfig
}

// a port part in configuration file
type PortConfig struct {
	Port string
}

var config Config

// HTTP method
var (
	get    = "GET"
	post   = "POST"
	put    = "PUT"
	delete = "DELETE"
)

// path
var (
	loginPath  = "/login"
	signInPath = "/signIn"
	postList   = "/postList"
)

func main() {

	// configuration file
	configFile := "./config/development.toml"

	// decoding toml
	_, err := toml.DecodeFile(configFile, &config)
	if err != nil {
		fmt.Println(err)
	}

	// initialize router
	r := mux.NewRouter()

	// Login
	r.HandleFunc(loginPath, controller.Login)

	// Sign-In
	r.HandleFunc(signInPath, controller.SignUp)

	// Post list
	r.HandleFunc(postList, controller.PostList)

	// listening port
	port := config.Port.Port
	// listenre
	http.ListenAndServe(port, r)
}
