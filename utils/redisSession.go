package utils

import (
	"io"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gomodule/redigo/redis"
	"github.com/shinshin8/myFavorite_backend/dto"
)

// Cache stores Redis connection
var Cache redis.Conn

var redisConfig dto.RedisConfig

// RedisConnection is connecting with Redis.
func RedisConnection() {

	logfile, er := os.OpenFile("./all-the-logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if er != nil {
		panic(er.Error())
	}
	defer logfile.Close()

	// decoding toml
	_, err := toml.DecodeFile(ConfigFile, &redisConfig)
	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}

	conn, err := redis.DialURL(redisConfig.Redis.RedisAddress)
	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}
	// Assign the connection to the package level `cache` variable
	Cache = conn
}
