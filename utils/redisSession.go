package utils

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/gomodule/redigo/redis"
	"github.com/shinshin8/myFavorite_backend/dto"
)

// Cache stores Redis connection
var Cache redis.Conn

var redisConfig dto.RedisConfig

// RedisConnection is connecting with Redis.
func RedisConnection() {

	// decoding toml
	_, err := toml.DecodeFile(ConfigFile, &redisConfig)
	if err != nil {
		fmt.Println(err)
	}

	conn, err := redis.DialURL(redisConfig.Redis.RedisAddress)
	if err != nil {
		panic(err)
	}
	// Assign the connection to the package level `cache` variable
	Cache = conn
}
