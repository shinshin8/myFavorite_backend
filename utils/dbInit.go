package utils

import (
	"database/sql"
	"io"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// DBInit initialize MySQL connection.
func DBInit() *sql.DB {
	logfile, er := os.OpenFile(ConfigFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if er != nil {
		panic(er.Error())
	}
	defer logfile.Close()

	dbUser := os.Getenv("USER")
	dbPassword := os.Getenv("PASSWORD")
	dbHost := os.Getenv("HOST")
	dbPort := os.Getenv("DB_PORT")
	database := os.Getenv("DATABASE")
	dbDriver := os.Getenv("DRIVER_NAME")

	dataSourceName := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + database

	sql, err := sql.Open(dbDriver, dataSourceName)

	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}

	return sql
}
