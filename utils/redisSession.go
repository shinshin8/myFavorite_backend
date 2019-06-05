package utils

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/gomodule/redigo/redis"
	"github.com/shinshin8/myFavorite/dto"
)

// Cache stores Redis connection
var Cache redis.Conn

var redisIPConfig dto.RedisIPConfig

// RedisConnection is connecting with Redis.
func RedisConnection() {
	// decoding toml
	_, err := toml.DecodeFile(ConfigFile, &redisIPConfig)
	if err != nil {
		fmt.Println(err)
	}
	redisIP := redisIPConfig.RedisIP
	conn, err := redis.DialURL(redisIP)
	if err != nil {
		panic(err)
	}
	// Assign the connection to the package level `cache` variable
	Cache = conn
}
