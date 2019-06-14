package model

import (
	"io"
	"log"
	"os"

	"github.com/shinshin8/myFavorite_backend/dto"
	"github.com/shinshin8/myFavorite_backend/utils"
)

// ShowLikedPosts returns the result of selected liked posts in JSON format.
func ShowLikedPosts(userID int) []dto.Posts {
	logfile, er := os.OpenFile("./all-the-logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if er != nil {
		panic(er.Error())
	}
	defer logfile.Close()

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
							liked_table.user_id = ?
						ORDER BY article_table.created_time DESC`

	row, err := sql.Query(getLikedPosts, userID)

	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}

	// Prepare an array which save JSON results.
	var postArray []dto.Posts

	for row.Next() {
		posts := dto.Posts{}
		if err := row.Scan(&posts.ArticleID, &posts.UserName, &posts.Title, &posts.Content, &posts.CreatedTime, &posts.ModifiedTime); err != nil {
			log.SetOutput(io.MultiWriter(logfile, os.Stdout))
			log.SetFlags(log.Ldate | log.Ltime)
			log.Fatal(err)
		}

		// Appending JSON in array.
		postArray = append(postArray, posts)
	}

	return postArray
}

// LikePost create new like post record in MySQL and returns the result in boolean.
// In the first parameter, user-id will be set with int type.
// In the second paraeter, article-id will be set witn int type.
func LikePost(userID int, articleID int) bool {
	logfile, er := os.OpenFile("./all-the-logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if er != nil {
		panic(er.Error())
	}
	defer logfile.Close()

	// Initalize DB Connection
	sql := utils.DBInit()
	// Close DB connection at the end.
	defer sql.Close()

	insertSyntax := `INSERT INTO 
						liked_table
						(user_id, 
						article_id) 
					VALUES
						(?,?)`

	rows, err := sql.Prepare(insertSyntax)

	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}

	rows.Exec(userID, articleID)
	return true
}

// DeleteLikedPost deletes a specific liked post record from MySQL and return boolean.
// At the first parameter, user id will be set with int type.
// At the second parameter, article id will be set with int type.
func DeleteLikedPost(userID int, articleID int) bool {
	logfile, er := os.OpenFile("./all-the-logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if er != nil {
		panic(er.Error())
	}
	defer logfile.Close()

	// Initalize DB Connection
	sql := utils.DBInit()
	// Close DB connection at the end.
	defer sql.Close()
	// SQL syntax
	delRec := `DELETE FROM 
					liked_table 
				WHERE 
					user_id = ? 
				AND 
					article_id = ?`

	rows, err := sql.Prepare(delRec)

	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}

	rows.Exec(userID, articleID)
	return true
}
