package model

import (
	"io"
	"log"
	"os"

	"github.com/shinshin8/myFavorite_backend/dto"
	"github.com/shinshin8/myFavorite_backend/utils"
)

// UploadImage registers image path from s3 into DB.
func UploadImage(imageData []dto.ImageStruct) bool {
	logfile, er := os.OpenFile(utils.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if er != nil {
		panic(er.Error())
	}
	defer logfile.Close()
	// Initalize DB Connection
	sql, sqlErr := utils.DBInit()
	if sqlErr != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(sqlErr)
	}
	// Close DB connection at the end.
	defer sql.Close()

	insertSyntax := `INSERT INTO 
						photo_table(
							photo_url, 
							user_id, 
							article_id
							) 
					VALUES`

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
	res, insertErr := rows.Exec(vals...)
	if res == nil || insertErr != nil {
		return false
	}
	return true
}

// DeleteImage delete serveral image records from DB.
func DeleteImage(imageData []dto.ImageStruct) bool {
	logfile, er := os.OpenFile(utils.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if er != nil {
		panic(er.Error())
	}
	defer logfile.Close()
	// Initalize DB Connection
	sql, sqlErr := utils.DBInit()
	if sqlErr != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(sqlErr)
	}
	// Close DB connection at the end.
	defer sql.Close()

	insertSyntax := `DELETE FROM 
							photo_table 
					WHERE (
						photo_url, 
						user_id, 
						article_id
						) 
					IN (`

	vals := []interface{}{}

	for _, row := range imageData {
		insertSyntax += "(?, ?, ?),"
		vals = append(vals, row.ImageURL, row.UserID, row.ArticleID)
	}
	//trim the last ,
	insertSyntax = insertSyntax[0 : len(insertSyntax)-1]
	insertSyntax += ")"

	rows, err := sql.Prepare(insertSyntax)
	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}
	res, insertErr := rows.Exec(vals...)
	if res == nil || insertErr != nil {
		return false
	}
	return true
}

// GetAllImages gets all images data from DB.
func GetAllImages() []dto.ImageStruct {
	logfile, er := os.OpenFile(utils.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if er != nil {
		panic(er.Error())
	}
	defer logfile.Close()
	// Initalize DB Connection
	sql, sqlErr := utils.DBInit()
	if sqlErr != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(sqlErr)
	}
	// Close DB connection at the end.
	defer sql.Close()
	// SQL syntax
	getImages := `SELECT 
						photo_url, 
						user_id, 
						article_id 
					FROM 
						photo_table;`

	row, err := sql.Query(getImages)

	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}
	// Prepare an array which save JSON results.
	var imageArray []dto.ImageStruct

	for row.Next() {
		posts := dto.ImageStruct{}
		if err := row.Scan(&posts.ImageURL, &posts.UserID, &posts.ArticleID); err != nil {
			log.SetOutput(io.MultiWriter(logfile, os.Stdout))
			log.SetFlags(log.Ldate | log.Ltime)
			log.Fatal(err)
		}
		// Appending JSON in array.
		imageArray = append(imageArray, posts)
	}
	return imageArray
}

// GetUserPostImageList get image list of user's post from DB.
func GetUserPostImageList(userID int) []dto.PostImage {
	logfile, er := os.OpenFile(utils.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if er != nil {
		panic(er.Error())
	}
	defer logfile.Close()
	// Initalize DB Connection
	sql, sqlErr := utils.DBInit()
	if sqlErr != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(sqlErr)
	}
	// Close DB connection at the end.
	defer sql.Close()
	// SQL syntax
	getImages := `SELECT 
						photo_url, 
						user_id, 
						article_id 
					FROM 
						photo_table
					WHERE
						user_id = ?;`

	row, err := sql.Query(getImages, userID)

	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}
	// Prepare an array which save JSON results.
	var imageArray []dto.PostImage

	for row.Next() {
		posts := dto.PostImage{}
		if err := row.Scan(&posts.ImageURL, &posts.UserID, &posts.ArticleID); err != nil {
			log.SetOutput(io.MultiWriter(logfile, os.Stdout))
			log.SetFlags(log.Ldate | log.Ltime)
			log.Fatal(err)
		}
		// Appending JSON in array.
		imageArray = append(imageArray, posts)
	}
	return imageArray
}

// GetSiglePostImages gets images for a single post from DB.
func GetSiglePostImages(articleID int) []string {
	logfile, er := os.OpenFile(utils.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if er != nil {
		panic(er.Error())
	}
	defer logfile.Close()
	// Initalize DB Connection
	sql, sqlErr := utils.DBInit()
	if sqlErr != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(sqlErr)
	}
	// Close DB connection at the end.
	defer sql.Close()
	// SQL syntax
	getImages := `SELECT 
						photo_url 
					FROM 
						photo_table 
					WHERE 
						article_id = ?`

	row, err := sql.Query(getImages, articleID)

	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}
	// Prepare an array which save JSON results.
	var imageArray []string

	for row.Next() {
		var image string
		if err := row.Scan(&image); err != nil {
			log.SetOutput(io.MultiWriter(logfile, os.Stdout))
			log.SetFlags(log.Ldate | log.Ltime)
			log.Fatal(err)
		}
		// Appending JSON in array.
		imageURL := os.Getenv("S3_URL") + image
		imageArray = append(imageArray, imageURL)
	}
	return imageArray
}

// GetFavoriteListImage gets user's favorite post list images from DB.
func GetFavoriteListImage(userID int) []dto.PostImage {
	logfile, er := os.OpenFile(utils.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if er != nil {
		panic(er.Error())
	}
	defer logfile.Close()
	// Initalize DB Connection
	sql, sqlErr := utils.DBInit()
	if sqlErr != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(sqlErr)
	}
	// Close DB connection at the end.
	defer sql.Close()
	// SQL syntax
	getImages := `SELECT 
						photo_table.photo_url, 
						favorite_table.user_id, 
						favorite_table.article_id 
					FROM 
						photo_table 
					INNER JOIN 
						favorite_table 
					ON 
						photo_table.article_id = favorite_table.article_id 
					WHERE favorite_table.user_id = ?;`

	row, err := sql.Query(getImages, userID)

	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}
	// Prepare an array which save JSON results.
	var imageArray []dto.PostImage

	for row.Next() {
		posts := dto.PostImage{}
		if err := row.Scan(&posts.ImageURL, &posts.UserID, &posts.ArticleID); err != nil {
			log.SetOutput(io.MultiWriter(logfile, os.Stdout))
			log.SetFlags(log.Ldate | log.Ltime)
			log.Fatal(err)
		}
		// Appending JSON in array.
		imageArray = append(imageArray, posts)
	}
	return imageArray
}
