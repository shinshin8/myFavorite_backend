package dto

// PostList is a struct for the DB query result JSON.
type PostList struct {
	Status    bool    `json:"status"`
	ErrorCode int     `json:"error_code"`
	UserID    int     `json:"user_id"`
	Posts     []Posts `json:"posts"`
}

// Article is a struct for article data from DB.
type Article struct {
	UserID       int
	ArticleID    int
	LikedSum     int
	UserName     string
	Title        string
	Content      string
	CreatedTime  string
	ModifiedTime string
}

// SiglePost saves a single post data in JSON.
type SiglePost struct {
	Status      bool             `json:"status"`
	ErrorCode   int              `json:"error_code"`
	UserID      int              `json:"user_id"`
	LikedFlg    bool             `json:"liked_flg"`
	FavoriteFlg bool             `json:"favorite_flg"`
	Post        SinglePostDetail `json:"post"`
}

// NewPost is a struct for a new post.
type NewPost struct {
	Status    bool   `json:"status"`
	ErrorCode int    `json:"error_code"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

// EditPostBody is a struct for request body of create post.
type EditPostBody struct {
	Title   string
	Content string
}

// Posts is a struct for post data.
type Posts struct {
	ArticleID   int    `json:"article_id"`
	LikedSum    int    `json:"liked_sum"`
	ImageURL    string `json:"image_url"`
	IconURL     string `json:"icon_url"`
	UserName    string `json:"user_name"`
	Title       string `json:"title"`
	CreatedTime string `json:"created_time"`
}

// SinglePostDetail is a struct for single post.
type SinglePostDetail struct {
	ArticleID    int      `json:"article_id"`
	LikedSum     int      `json:"liked_sum"`
	ImageURL     []string `json:"image_url"`
	IconURL      string   `json:"icon_url"`
	UserName     string   `json:"user_name"`
	Title        string   `json:"title"`
	Content      string   `json:"content"`
	CreatedTime  string   `json:"created_time"`
	ModifiedTime string   `json:"modified_time"`
}

// PostImage is a struct for images.
type PostImage struct {
	ImageURL  []string
	UserID    int
	ArticleID int
}
