package model

import (
	"log"

	"../dto"
	"../utils"
)

// ShowFavoritePosts returns the list of specific user's favorite posts in JSON.
// At the first parameter, user id will be set with int type.
// At the second paramtere, article id will be set with int type.
func ShowFavoritePosts(userID int, articleID int) []dto.Posts {
	// Initalize DB Connection
	sql := utils.DBInit()
	// Close DB connection at the end.
	defer sql.Close()
	// SQL syntax
	getFavoritePosts := `SELECT 
							article_table.article_id, 
							user_table.user_name, 
							article_table.title, 
							article_table.content, 
							article_table.created_time, 
							article_table.modified_time 
						FROM 
							(favorite_table 
						INNER JOIN 
							user_table 
						ON 
							user_table.user_id = favorite_table.user_id) 
						INNER JOIN 
							article_table 
						ON 
							article_table.article_id = favorite_table.article_id 
						WHERE 
							favorite_table.user_id = ?
						ORDER BY article_table.created_time DESC`

	row, err := sql.Query(getFavoritePosts, userID)

	if err != nil {
		log.Fatal(err)
	}

	// Prepare an array which save JSON results.
	var favoritePostArray []dto.Posts

	for row.Next() {
		posts := dto.Posts{}
		if err := row.Scan(&posts.ArticleID, &posts.UserName, &posts.Title, &posts.Content, &posts.CreatedTime, &posts.ModifiedTime); err != nil {
			log.Fatal(err)
		}

		// Appending JSON in array.
		favoritePostArray = append(favoritePostArray, posts)
	}

	return favoritePostArray
}