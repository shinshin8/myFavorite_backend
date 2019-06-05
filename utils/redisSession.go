package utils

import (
	"github.com/gomodule/redigo/redis"
)

// Cache stores Redis connection
var Cache redis.Conn

// RedisConnection is connecting with Redis.
func RedisConnection(){
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	// Assign the connection to the package level `cache` variable
	Cache = conn
}