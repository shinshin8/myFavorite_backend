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
	Authorization   = "Authorization"
)

// HTTP request method
const (
	Post    = "POST"
	Get     = "GET"
	Put     = "PUT"
	Delete  = "DELETE"
	Options = "OPTIONS"
)

// Each path
const (
	Trending           = "/"
	Timeline           = "/timeline"
	LoginPath          = "/login"
	SignUpPath         = "/signUp"
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
	DeleteAccount      = "/deleteAccount"
)

// ConfigFile indicate configuration file path.
const ConfigFile = "./config/config.toml"

// Error code
const (
	SuccessCode                   = 0
	WrongUserNamePassword         = 1
	InvalidToken                  = 2
	InvalidSignUpUsername         = 3
	InvalidSignUpMailAddress      = 4
	InvalidSignUpPassword         = 5
	NotMatchPasswords             = 6
	FailedLoginCode               = 7
	FailedDeleteLike              = 8
	FailedFavoritePost            = 9
	FailedDeleteFavorite          = 10
	FailedDeletePost              = 11
	InvalidCreateTitle            = 12
	InvalidCreateContent          = 13
	InvalidEditTitle              = 14
	InvalidEditContent            = 15
	InvalidEditProfileUserName    = 16
	InvalidEditProfileBirthday    = 17
	InvalidEditProfileMailAddress = 18
	InvalidEditProfileComment     = 19
	FailedDeleteAccount           = 20
	FailedEditProfile             = 21
	FailedEditPost                = 22
)
