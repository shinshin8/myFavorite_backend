package utils

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/shinshin8/myFavorite_backend/dto"
)

var logFileConfig dto.LogConfig
var tokenStringConfig dto.TokenStringConfig

// CreateToken creates JWT.
func CreateToken(userID int) string {

	// decoding toml
	_, ers := toml.DecodeFile(ConfigFile, &logFileConfig)
	if ers != nil {
		panic(ers.Error())
	}

	_, tokenErr := toml.DecodeFile(ConfigFile, &tokenStringConfig)
	if tokenErr != nil {
		panic(tokenErr.Error())
	}

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
	tokenString, err := token.SignedString([]byte(tokenStringConfig.TokenString.TokenString))
	if err != nil {
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal(err)
	}
	return tokenString
}

// VerifyToken checks if token is valid.
func VerifyToken(reqHead string) int {
	_, tokenErr := toml.DecodeFile(ConfigFile, &tokenStringConfig)
	if tokenErr != nil {
		panic(tokenErr.Error())
	}

	token, _ := jwt.Parse(reqHead, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenStringConfig.TokenString.TokenString), nil
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
