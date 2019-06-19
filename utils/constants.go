package utils

// Application form type
const (
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json;charset=utf-8"
	Cors            = "Access-Control-Allow-Origin"
	ArrowHeader     = "Access-Control-Allow-Headers"
	Credential      = "Access-Control-Allow-Credentials"
	ArrowMethods    = "Access-Control-Allow-Methods"
	Methods         = "PUT,POST,GET,DELETE,OPTIONS"
	True            = "true"
	CorsWildCard    = "*"
)

// HTTP request method
const (
	Post   = "POST"
	Get    = "GET"
	Put    = "PUT"
	Delete = "DELETE"
)

// Session
const (
	SessionName       = "session-name"
	ContextSessionKey = "session"
)

// Each path
const (
	LoginPath          = "/login"
	SignUpPath         = "/signUp"
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
	EditPost           = "/editPost"
	DeletePost         = "/deletePost"
	ShowProfile        = "/showProfile"
	EditProfile        = "/editProfile"
	Logout             = "/logout"
	DeleteAccount      = "/deleteAccount"
)

// ConfigFile indicate configuration file path.
const ConfigFile = "./config/config.toml"

// Session method
const (
	SessionGet     = "GET"
	SessionSet     = "SETEX"
	SessionDelete  = "DEL"
	CookieName     = "session_token"
	SessionTimeOut = "1200"
	SessionExpire  = 1200
)
