package dto

// SiglePost saves a single post data in JSON.
type SiglePost struct {
	Status       int    `json:"status"`
	ErrorCode    int    `json:"error_code"`
	LikedFlg     bool   `json:liked_flg`
	FavoriteFlg  bool   `json:favorite_flg`
	ArticleID    int    `json:article_id`
	UserName     string `json:user_name`
	Title        string `json:title`
	Content      string `json:content`
	CreatedTime  string `json:created_time`
	ModifiedTime string `json:modified_time`
}
