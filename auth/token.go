package auth

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type jwtCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var key string

func init() {
	valuesEnv := godotenv.Load()
	if valuesEnv != nil {
		log.Fatal(valuesEnv)
	}
	key = os.Getenv("key_token")
}

func CreateToken(username string) string {
	claims := &jwtCustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), //token exiresin 1 hour
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	resToken, err := token.SignedString([]byte(key))
	if err != nil {
		return err.Error()
	}
	return resToken
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	justToken := strings.Split(tokenString, " ")
	tokenString = justToken[1]
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}
