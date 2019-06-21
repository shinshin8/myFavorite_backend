package utils

import (
	"database/sql"
	"io"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shinshin8/myFavorite_backend/dto"
)

var dbConfig dto.DbConfig

// DBInit initialize MySQL connection.
func DBInit() *sql.DB {

	// decoding toml
	_, ers := toml.DecodeFile(ConfigFile, &logFileConfig)
	if ers != nil {
		panic(ers.Error())
	}

	logfile, er := os.OpenFile(ConfigFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if er != nil {
		panic(er.Error())
	}
	defer logfile.Close()

	// decoding toml
	_, err := toml.DecodeFile(ConfigFile, &dbConfig)
	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}

	dataSourceName := dbConfig.Database.User + ":" + dbConfig.Database.Password + "@tcp(" + dbConfig.Database.Host + ":" + dbConfig.Database.DbPort + ")/" + dbConfig.Database.Database

	sql, err := sql.Open(dbConfig.Database.DriverName, dataSourceName)

	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}

	return sql
}
