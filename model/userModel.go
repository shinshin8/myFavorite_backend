/*
	This file handles CRUD for User Table.
*/

package model

import (
	"log"

	"../dto"
	"../utils"
)

// UserName struct
type UserName struct {
	Username string
}

// LoginUser judges wheather the recieved login information is corrent or not.
// At first parameter, username is recieved and its type is string.
// At second parameter, hashed password is recieved and its type is string.
// The function return true or false.
func LoginUser(username string, hashedPassword string) bool {

	sql := utils.DBInit()

	// at the end, sql will be closed.
	defer sql.Close()

	// SQL syntax
	findUserSyntax := "SELECT user_id FROM user_table WHERE user_name = ? AND password = ?;"

	// execute SQL manipulation
	rows, err := sql.Query(findUserSyntax, username, hashedPassword)

	if err != nil {
		log.Fatal(err)
	}

	var userNameArr []UserName

	listedUser := UserName{}

	for rows.Next() {
		if err := rows.Scan(&listedUser.Username); err != nil {
			log.Fatal(err)
		}

		// appending an array stores selected username into struct:listedUser.
		userNameArr = append(userNameArr, listedUser)
	}

	// expected array's length.
	expectedLength := 1

	// array's length
	arrLength := len(userNameArr)

	return arrLength == expectedLength

}

// SignUp returns sign-up resutl in JSON format.
// Username is in the first parameter with string type.
// Email Address is in the second parameter with string type.
// Password is in the third parameter with string type.
func SignUp(username string, emailAddress string, password string) dto.SignUpResultJSON {
	// Initalize DB Connection
	sql := utils.DBInit()
	// Close DB connection at the end.
	defer sql.Close()
	// SQL syntax
	registerNewUser := "INSERT into user_table(user_name, password, mail_address) values(?, ?, ?)"
	rows, err := sql.Prepare(registerNewUser)
	httpOk := 200
	if err != nil {
		log.Fatal(err)
		sqlErrorStatus := 8
		res := dto.SignUpResultJSON{
			Status:       httpOk,
			ErrorCode:    sqlErrorStatus,
			Username:     username,
			EmailAddress: emailAddress,
		}
		return res
	}
	rows.Exec(username, password, emailAddress)
	successStatus := 0

	res := dto.SignUpResultJSON{
		Status:       httpOk,
		ErrorCode:    successStatus,
		Username:     username,
		EmailAddress: emailAddress,
	}
	return res

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
