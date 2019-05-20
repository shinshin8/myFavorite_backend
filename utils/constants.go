package utils

// Application form type
const (
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json"
)

// HTTP request method
const (
	Post   = "POST"
	Get    = "GET"
	Put    = "PUT"
	Delete = "DELETE"
)

// Each path
const (
	LoginPath          = "/login"
	SignUpPath         = "/signUn"
	PostList           = "/"
	ShowLikedPosts     = "/likedPostsList"
	LikePost           = "/likePost"
	DeleteLikedPost    = "/deleteLikedPost"
	ShowFavoritePosts  = "/favoritePostsList"
	FavoritePost       = "/favoritePost"
	DeleteFavoritePost = "/deleteFavorite"
	UserPostsList      = "/userPostsList"
	SinglePost         = "/post"
	NewPost            = "/newPost"
)

// The directory of configuration file
const ConfigFile = "./config/development.toml"
