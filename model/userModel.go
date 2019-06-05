package model

import (
	"log"

	"github.com/shinshin8/myFavorite/dto"
	"github.com/shinshin8/myFavorite/utils"
)

// UserID is for logined user's id.
type UserID struct {
	UserID int
}

// LoginUser judges wheather the recieved login information is corrent or not.
// At first parameter, username is recieved and its type is string.
// At second parameter, hashed password is recieved and its type is string.
// The function return true or false.
func LoginUser(username string, hashedPassword string) UserID {

	sql := utils.DBInit()

	// at the end, sql will be closed.
	defer sql.Close()

	// SQL syntax
	findUserSyntax := "SELECT user_id FROM user_table WHERE user_name = ? AND password = ?;"

	var userID UserID

	err := sql.QueryRow(findUserSyntax, username, hashedPassword).Scan(&userID.UserID)

	if err != nil {
		log.Fatal(err)
	}

	return userID

}

// SignUp returns inserted user's id in int64.
// Username is in the first parameter with string type.
// Email Address is in the second parameter with string type.
// Password is in the third parameter with string type.
func SignUp(username string, emailAddress string, password string) int64 {
	// Initalize DB Connection
	sql := utils.DBInit()
	// Close DB connection at the end.
	defer sql.Close()
	// SQL syntax
	registerNewUser := "INSERT into user_table(user_name, password, mail_address) values(?, ?, ?)"

	res, err := sql.Exec(registerNewUser, username, password, emailAddress)

	if err != nil {
		println("Exec err:", err.Error())
	}

	id, err := res.LastInsertId()

	if err != nil {
		println("Error:", err.Error())
	}

	return id

}

// ShowProfile gets a user's profile from user_table and return its result in JSON.
// At the first parameter, user id is set in int type.
func ShowProfile(userID int) dto.Profile {
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
func EditProfile(userID int, userName, birthday, mailAddress, comment string) dto.SimpleResutlJSON {
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
	httpOk := 200
	if err != nil {
		log.Fatal(err)
		sqlErrorStatus := 8
		res := dto.SimpleResutlJSON{
			Status:    httpOk,
			ErrorCode: sqlErrorStatus,
		}
		return res
	}
	rows.Exec(userName, birthday, mailAddress, comment, userID)
	successStatus := 0

	res := dto.SimpleResutlJSON{
		Status:    httpOk,
		ErrorCode: successStatus,
	}
	return res
}
