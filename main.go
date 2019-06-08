package main

import (
	"fmt"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
	"github.com/shinshin8/myFavorite/controller"
	"github.com/shinshin8/myFavorite/dto"
	"github.com/shinshin8/myFavorite/utils"
)

var portConfig dto.PortConfig

func main() {
	// decoding toml
	_, err := toml.DecodeFile(utils.ConfigFile, &portConfig)
	if err != nil {
		fmt.Println(err)
	}
	utils.RedisConnection()
	// initialize router
	r := mux.NewRouter()
	// Login
	r.HandleFunc(utils.LoginPath, controller.Login).Methods(utils.Post)
	// Sign-Up
	r.HandleFunc(utils.SignUpPath, controller.SignUp).Methods(utils.Post)
	// Post list
	r.HandleFunc(utils.PostList, controller.PostList).Methods(utils.Get)
	// Show liked post lists
	r.HandleFunc(utils.ShowLikedPosts, controller.ShowLikedPosts).Methods(utils.Get)
	// Like post
	r.HandleFunc(utils.LikePost, controller.LikePost).Methods(utils.Put)
	// Delete liked post
	r.HandleFunc(utils.DeleteLikedPost, controller.DeleteLikedPost).Methods(utils.Delete)
	// Show favorite posts
	r.HandleFunc(utils.ShowFavoritePosts, controller.ShowFavoritePosts).Methods(utils.Get)
	// Create a favorite post
	r.HandleFunc(utils.FavoritePost, controller.FavoritePost).Methods(utils.Put)
	// Delete a favorite post
	r.HandleFunc(utils.DeleteFavoritePost, controller.DeleteFavoritePost).Methods(utils.Delete)
	// User's posts list
	r.HandleFunc(utils.UserPostsList, controller.UserPostsList).Methods(utils.Get)
	// Single post
	r.HandleFunc(utils.SinglePost, controller.SinglePost).Methods(utils.Get)
	// Create a post
	r.HandleFunc(utils.NewPost, controller.CreatePost).Methods(utils.Post)
	// Edit a post
	r.HandleFunc(utils.EditPost, controller.EditPost).Methods(utils.Put)
	// Delete a post
	r.HandleFunc(utils.DeletePost, controller.DeletePost).Methods(utils.Delete)
	// Show user's profile.
	r.HandleFunc(utils.ShowProfile, controller.ShowProfile).Methods(utils.Get)
	// Edit user's profile.
	r.HandleFunc(utils.EditProfile, controller.EditProfile).Methods(utils.Put)
	// Logout
	r.HandleFunc(utils.Logout, controller.Logout).Methods(utils.Post)
	// Delete Account
	r.HandleFunc(utils.DeleteAccount, controller.DeleteAccount).Methods(utils.Delete)
	// listening port
	port := portConfig.Port.Port
	// listener
	http.ListenAndServe(port, r)
}
