package test

import (
	"reflect"
	"testing"

	"github.com/shinshin8/myFavorite_backend/utils"
)

func TestCreateToken(t *testing.T) {
	var userID int
	userID = 1
	var example string
	example = "example"
	res := utils.CreateToken(userID)
	if reflect.TypeOf(res) != reflect.TypeOf(example) && len(res) == 0 {
		t.Fatal("failed test")
	}
}

func TestVerifyToken(t *testing.T) {
	var userID int
	userID = 1
	toke := utils.CreateToken(userID)
	res := utils.VerifyToken(toke)
	if res != userID {
		t.Fatal("failed test")
	}
}
