package model

import (
	"io"
	"log"
	"os"

	"github.com/shinshin8/myFavorite_backend/dto"
	"github.com/shinshin8/myFavorite_backend/utils"
)

// LoginUser judges wheather the recieved login information is corrent or not.
// At first parameter, username is recieved and its type is string.
// At second parameter, hashed password is recieved and its type is string.
// The function return userID or 0.
func LoginUser(username string, hashedPassword string) int {
	logfile, er := os.OpenFile(utils.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if er != nil {
		panic(er.Error())
	}
	defer logfile.Close()

	sql := utils.DBInit()

	// at the end, sql will be closed.
	defer sql.Close()

	// SQL syntax
	findUserSyntax := `SELECT 
							user_id 
						FROM 
							user_table 
						WHERE user_name = ? AND password = ?`

	var userID int

	err := sql.QueryRow(findUserSyntax, username, hashedPassword).Scan(&userID)

	if err != nil {
		return 0
	}

	return userID

}

// SignUp returns inserted user's id in int64.
// Username is in the first parameter with string type.
// Email Address is in the second parameter with string type.
// Password is in the third parameter with string type.
func SignUp(username string, emailAddress string, password string) int {
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
	registerNewUser := "INSERT into user_table(user_name, password, mail_address) values(?, ?, ?)"

	res, err := sql.Exec(registerNewUser, username, password, emailAddress)

	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}

	return int(id)

}

// ShowProfile gets a user's profile from user_table and return its result in JSON.
// At the first parameter, user id is set in int type.
func ShowProfile(userID int) dto.Profile {
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
	selectProfile := `SELECT
							user_id,
							user_name, 
							mail_address, 
							birthday, 
							comment 
						FROM 
							user_table 
						WHERE 
							user_id = ?`

	var profile dto.Profile

	err := sql.QueryRow(selectProfile, userID).Scan(&profile.UserID, &profile.UserName, &profile.MailAddress, &profile.Birthday, &profile.Comment)

	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}

	return profile
}

// EditProfile updates user_table and return the result in JSON.
// At first parameter, user id is set in int type.
// At first parameter, user name is set in string type.
// At second paramter, birthday is set in string type.
// At third paramter, mail address is set in string type.
// At fourth paramter, comment is set in string type.
func EditProfile(userID int, userName, birthday, mailAddress, comment string) bool {
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
	update := `UPDATE 
					user_table 
				SET 
					user_name = ?, 
					birthday = ?, 
					mail_address = ?, 
					comment = ? 
				WHERE 
					user_id = ?`
	rows, err := sql.Prepare(update)
	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}
	res, err := rows.Exec(userName, birthday, mailAddress, comment, userID)

	if res == nil || err != nil {
		return false
	}

	return true
}

// DeleteAccount delete user from DB
// At first parameter, user id is set in int type.
// The function returns result in boolean.
func DeleteAccount(userID int) bool {
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
	deleteUser := `DELETE FROM 
						user_table 
					WHERE 
						user_id = ?`
	rows, err := sql.Prepare(deleteUser)
	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}
	rows.Exec(userID)

	return true
}
