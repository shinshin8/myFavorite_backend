package utils

// Application form type
const (
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json"
)

// HTTP request method
const (
	Post = "POST"
	Get  = "GET"
	Put  = "PUT"
)

// Each path
const (
	LoginPath      = "/login"
	SignInPath     = "/signIn"
	PostList       = "/"
	ShowLikedPosts = "/likedPostsList"
	LikePost       = "/likePost"
)

// The directory of configuration file
const ConfigFile = "./config/development.toml"
