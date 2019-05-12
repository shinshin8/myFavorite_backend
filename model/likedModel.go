package model

import (
	"log"

	"../dto"
	"../utils"
)

// ShowLikedPosts returns the result of selected liked posts in JSON format.
func ShowLikedPosts(userID int) []dto.Posts {
	// Initalize DB Connection
	sql := utils.DBInit()
	// Close DB connection at the end.
	defer sql.Close()
	// SQL syntax

	getLikedPosts := `SELECT 
							article_table.article_id, 
							user_table.user_name, 
							article_table.title, 
							article_table.content, 
							article_table.created_time, 
							article_table.modified_time 
						FROM 
							(liked_table 
						INNER JOIN 
							user_table 
						ON 
							user_table.user_id = liked_table.user_id) 
						INNER JOIN 
							article_table 
						ON 
							article_table.article_id = liked_table.article_id 
						WHERE 
							liked_table.user_id = ?`

	row, err := sql.Query(getLikedPosts, userID)

	if err != nil {
		log.Fatal(err)
	}

	// Prepare an array which save JSON results.
	var postArray []dto.Posts

	for row.Next() {
		posts := dto.Posts{}
		if err := row.Scan(&posts.ArticleID, &posts.UserName, &posts.Title, &posts.Content, &posts.CreatedTime, &posts.ModifiedTime); err != nil {
			log.Fatal(err)
		}

		// Appending JSON in array.
		postArray = append(postArray, posts)
	}

	return postArray
}
