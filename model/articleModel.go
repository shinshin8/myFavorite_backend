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
	getPosts := `SELECT 
					article_table.article_id, 
					user_table.user_name, 
					article_table.title, 
					article_table.content,
					article_table.created_time, 
					article_table.modified_time 
				FROM 
					article_table 
				INNER JOIN 
					user_table 
				ON 
					article_table.user_id = user_table.user_id 
				ORDER BY 
					article_table.created_time DESC`

	row, err := sql.Query(getPosts)

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

// UserPostsList gets a specific user's posts list from DB and convert its result into JSON.
// At the parameter, user id will be put in and its type is int.
func UserPostsList(userID int) []dto.Posts {
	// Initalize DB Connection
	sql := utils.DBInit()
	// Close DB connection at the end.
	defer sql.Close()
	// SQL syntax
	getPosts := `SELECT 
					article_table.article_id, 
					user_table.user_name, 
					article_table.title, 
					article_table.content,
					article_table.created_time, 
					article_table.modified_time 
				FROM 
					article_table 
				INNER JOIN 
					user_table 
				ON 
					article_table.user_id = user_table.user_id
				WHERE
					user_table.user_id = ?
				ORDER BY 
					article_table.created_time DESC`

	row, err := sql.Query(getPosts, userID)

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
