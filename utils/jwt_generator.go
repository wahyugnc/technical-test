package utils

import (
	"os"
	"strconv"
	"time"
	"usedeall/config"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateNewAccessToken(username string, role string, userId int) (string, error) {
	configuration := config.New()
	secret := configuration.Get("JWT_SECRET_KEY")
	expTime, _ := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_EXPIRE"))
	claims := jwt.MapClaims{}

	claims["username"] = username
	claims["role"] = role
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expTime)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func GenerateNewRefreshToken(username string, role string, userId int) (string, error) {
	configuration := config.New()
	secret := configuration.Get("JWT_SECRET_KEY")
	expTime, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_EXPIRE"))
	claims := jwt.MapClaims{}

	claims["username"] = username
	claims["role"] = role
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expTime)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}
