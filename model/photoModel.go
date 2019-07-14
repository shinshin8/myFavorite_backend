package model

import (
	"io"
	"log"
	"os"

	"github.com/shinshin8/myFavorite_backend/dto"
	"github.com/shinshin8/myFavorite_backend/utils"
)

// UploadImage registers image path from s3 into DB.
func UploadImage(imageData []dto.UploadImage) bool {
	logfile, er := os.OpenFile(utils.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if er != nil {
		panic(er.Error())
	}
	defer logfile.Close()
	// Initalize DB Connection
	sql := utils.DBInit()
	// Close DB connection at the end.
	defer sql.Close()

	insertSyntax := `INSERT INTO photo_table(photo_url, user_id, article_id) VALUES`

	vals := []interface{}{}

	for _, row := range imageData {
		insertSyntax += "(?, ?, ?),"
		vals = append(vals, row.ImageURL, row.UserID, row.ArticleID)
	}
	//trim the last ,
	insertSyntax = insertSyntax[0 : len(insertSyntax)-1]

	rows, err := sql.Prepare(insertSyntax)
	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}
	_, insertErr := rows.Exec(vals...)
	if insertErr != nil {
		return false
	}
	return true
}
