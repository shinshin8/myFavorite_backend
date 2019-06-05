package utils

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shinshin8/myFavorite/dto"
)

var dbConfig dto.DbConfig

// DBInit initialize MySQL connection.
func DBInit() *sql.DB {

	// decoding toml
	_, err := toml.DecodeFile(ConfigFile, &dbConfig)
	if err != nil {
		fmt.Println(err)
	}

	dataSourceName := dbConfig.Database.User + ":" + dbConfig.Database.Password + "@tcp(" + dbConfig.Database.Host + ":" + dbConfig.Database.DbPort + ")/" + dbConfig.Database.Database

	sql, err := sql.Open(dbConfig.Database.DriverName, dataSourceName)

	if err != nil {
		log.Fatal(err)
	}

	return sql
}
