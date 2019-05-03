package model

import (
	"log"

	"../dto"
	"../utils"
)

// GetPosts is a function that returns an array which includes DB results with JSON format.
func GetPosts() []dto.Posts {
	// Initalize DB Connection
	sql := utils.DBInit()
	// Close DB connection at the end.
	defer sql.Close()
	// SQL syntax
	getPosts := "SELECT article_id, user_id, title, created_time, modified_time FROM article_table ORDER BY created_time DESC"

	row, err := sql.Query(getPosts)

	if err != nil {
		log.Fatal(err)
	}

	// Prepare an array which save JSON results.
	var postArray []dto.Posts

	for row.Next() {
		posts := dto.Posts{}
		if err := row.Scan(&posts.ArticleID, &posts.UserID, &posts.Title, &posts.CreatedTime, &posts.ModifiedTime); err != nil {
			log.Fatal(err)
		}

		// Appending JSON in array.
		postArray = append(postArray, posts)
	}

	return postArray
}
