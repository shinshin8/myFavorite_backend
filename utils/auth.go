package utils

import (
	"io"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken creates JWT.
func CreateToken(userID int) string {
	logfile, er := os.OpenFile(ConfigFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if er != nil {
		panic(er.Error())
	}
	defer logfile.Close()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"iat":     time.Now(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_STRING")))
	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}
	return tokenString
}

// VerifyToken checks if token is valid.
func VerifyToken(reqHead string) int {
	token, _ := jwt.Parse(reqHead, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_STRING")), nil
	})
	if token == nil {
		return 0
	}
	// Get user id from jwt.
	claims := token.Claims.(jwt.MapClaims)
	userIDkey := "user_id"
	userIDFloat64, _ := claims[userIDkey]
	userID := int(userIDFloat64.(float64))

	return userID
}
