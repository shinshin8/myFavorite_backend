/*
	dbInit.go is handling DB connection.
*/

package utils

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DBInit initialize MySQL connection.
func DBInit() *sql.DB {
	sql, err := sql.Open("mysql", "my_favorite:my_favorite@tcp(127.0.0.1:3306)/my_favorite")

	if err != nil {
		log.Fatal(err)
	}

	return sql
}
