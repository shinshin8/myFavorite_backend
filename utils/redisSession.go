package utils

import (
	"gopkg.in/boj/redistore.v1"
)

// ResisSesssion connects Redis to handle sessions.
func ResisSesssion() *redistore.RediStore {
	store, err := redistore.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
	if err != nil {
		panic(err)
	}
	store.SetMaxAge(1 * 24 * 3600)
	return store
}
