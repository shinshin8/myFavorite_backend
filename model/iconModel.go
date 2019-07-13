package model

import (
	"io"
	"log"
	"os"

	"github.com/shinshin8/myFavorite_backend/utils"
)

// RegisterIcon is the method to register icon path to DB.
func RegisterIcon(iconURL string, userID int) bool {
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
						icon_table
						(icon_url, 
						user_id) 
					VALUES
						(?,?)`

	rows, err := sql.Prepare(insertSyntax)

	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}

	rows.Exec(iconURL, userID)
	return true
}

// UpdateIcon delete from icon url from DB.
func UpdateIcon(newIconURL string, userID int) bool {
	logfile, er := os.OpenFile(utils.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if er != nil {
		panic(er.Error())
	}
	defer logfile.Close()
	// Initalize DB Connection
	sql := utils.DBInit()
	// Close DB connection at the end.
	defer sql.Close()

	updateSyntax := `UPDATE 
						icon_table 
					SET 
						icon_url = ? 
					WHERE 
						user_id = ?`

	rows, err := sql.Prepare(updateSyntax)

	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}

	rows.Exec(newIconURL, userID)
	return true
}
