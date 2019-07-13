package model

import (
	"io"
	"log"
	"os"

	"github.com/shinshin8/myFavorite_backend/utils"
)

// LikePost create new like post record in MySQL and returns the result in boolean.
// In the first parameter, user-id will be set with int type.
// In the second paraeter, article-id will be set witn int type.
func LikePost(userID int, articleID int) bool {
	logfile, er := os.OpenFile(utils.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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
	logfile, er := os.OpenFile(utils.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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
