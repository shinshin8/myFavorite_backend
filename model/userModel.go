/*
	This file handles CRUD for User Table.
*/

package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// configuration file
type Config struct {
	Port DbConfig
}

// DB information in configuration file
type DbConfig struct {
	User     string
	Password string
	Database string
	Hostname string
	DbPort   string
}

var config Config

type UserName struct {
	Username string
}

// This function judges wheather the recieved login information is corrent or not.
// At first parameter, username is recieved and its type is string.
// At second parameter, hashed password is recieved and its type is string.
// The function return true or false.
func LoginUser(username string, hashedPassword string) bool {

	// // configuration file
	// configFile := "./config/development.toml"

	// // decoding toml
	// _, err := toml.DecodeFile(configFile, &config)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// driverName := "mysql"

	// // dataSourceName := config.Port.User + ":" + config.Port.Password + "@tcp(" + config.Port.Hostname + ":" + config.Port.DbPort + ")/" + config.Port.Database

	// dataSourceName := "my_favorite:my_favorite@tcp(127.0.0.1:3306)/myfavorite"

	// // open sql
	// sql, err := sql.Open(driverName, dataSourceName)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	sql, err := sql.Open("mysql", "my_favorite:my_favorite@tcp(127.0.0.1:3306)/my_favorite")

	if err != nil {
		log.Fatal(err)
	}

	// at the end, sql will be closed.
	defer sql.Close()

	// SQL syntax
	findUserSyntax := "SELECT user_id FROM user_table WHERE user_name = ? AND password = ?;"

	// execute SQL manipulation
	rows, err := sql.Query(findUserSyntax, username, hashedPassword)

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

	if arrLength == expectedLength {
		return true
	} else {
		return false
	}

}
