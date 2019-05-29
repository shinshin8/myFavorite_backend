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
	// initialize router
	r := mux.NewRouter()
	// Login
	utils.HandlerFunc(r, utils.LoginPath, controller.Login).Methods(utils.Post)
	// Sign-Up
	utils.HandlerFunc(r, utils.SignUpPath, controller.SignUp).Methods(utils.Post)
	// Post list
	utils.HandlerFunc(r, utils.PostList, controller.PostList).Methods(utils.Get)
	// Show liked post lists
	utils.HandlerFunc(r, utils.ShowLikedPosts, controller.ShowLikedPosts).Methods(utils.Get)
	// Like post
	utils.HandlerFunc(r, utils.LikePost, controller.LikePost).Methods(utils.Put)
	// Delete liked post
	utils.HandlerFunc(r, utils.DeleteLikedPost, controller.DeleteLikedPost).Methods(utils.Delete)
	// Show favorite posts
	utils.HandlerFunc(r, utils.ShowFavoritePosts, controller.ShowFavoritePosts).Methods(utils.Get)
	// Create a favorite post
	utils.HandlerFunc(r, utils.FavoritePost, controller.FavoritePost).Methods(utils.Put)
	// Delete a favorite post
	utils.HandlerFunc(r, utils.DeleteFavoritePost, controller.DeleteFavoritePost).Methods(utils.Delete)
	// User's posts list
	utils.HandlerFunc(r, utils.UserPostsList, controller.UserPostsList).Methods(utils.Get)
	// Single post
	utils.HandlerFunc(r, utils.SinglePost, controller.SinglePost).Methods(utils.Get)
	// Create a post
	utils.HandlerFunc(r, utils.NewPost, controller.CreatePost).Methods(utils.Post)
	// Edit a post
	utils.HandlerFunc(r, utils.EditPost, controller.EditPost).Methods(utils.Put)
	// Delete a post
	utils.HandlerFunc(r, utils.DeletePost, controller.DeletePost).Methods(utils.Delete)
	// Show user's profile.
	utils.HandlerFunc(r, utils.ShowProfile, controller.ShowProfile).Methods(utils.Get)
	// Edit user's profile.
	utils.HandlerFunc(r, utils.EditProfile, controller.EditProfile).Methods(utils.Put)
	// listening port
	port := portConfig.Port.Port
	// listener
	http.ListenAndServe(port, r)
}
