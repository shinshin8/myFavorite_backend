package dto

// PostList is a struct for the DB query result JSON.
type PostList struct {
	Status    int     `json:"status"`
	ErrorCode int     `json:"error_code"`
	Posts     []Posts `json:posts`
}

// Posts is a struct for listing posts.
type Posts struct {
	ArticleID int    `json:article_id`
	UserID    int    `json:user_id`
	Title     string `json:title`
}
