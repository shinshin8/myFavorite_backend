package dto

// PostList is a struct for the DB query result JSON.
type PostList struct {
	Status    int     `json:"status"`
	ErrorCode int     `json:"error_code"`
	Posts     []Posts `json:posts`
}

// Posts is a struct for listing posts.
type Posts struct {
	ArticleID    int    `json:article_id`
	UserName     string `json:user_name`
	Title        string `json:title`
	Content      string `json:content`
	CreatedTime  string `json:created_time`
	ModifiedTime string `json:modified_time`
}

// SiglePost saves a single post data in JSON.
type SiglePost struct {
	Status      int   `json:"status"`
	ErrorCode   int   `json:"error_code"`
	LikedFlg    bool  `json:liked_flg`
	FavoriteFlg bool  `json:favorite_flg`
	Post        Posts `json:post`
}

// NewPost is a struct for a new post.
type NewPost struct {
	Status    int    `json:"status"`
	ErrorCode int    `json:"error_code"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}
