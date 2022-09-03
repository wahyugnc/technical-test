package utils

import (
	"strings"
	"usedeall/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type TokenMetadata struct {
	Expires  int64
	UserId   int
	Username string
	Role     string
}

func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		expires := int64(claims["exp"].(float64))
		username := string(claims["username"].(string))
		role := string(claims["role"].(string))
		userId := int(claims["userId"].(float64))

		return &TokenMetadata{
			Expires:  expires,
			Username: username,
			Role:     role,
			UserId:   userId,
		}, nil
	}

	return nil, err
}

func ExtractRefreshTokenMetadata(c *fiber.Ctx, refreshToken string) (*TokenMetadata, error) {
	token, err := verifyRefreshToken(c, refreshToken)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		expires := int64(claims["exp"].(float64))
		username := string(claims["username"].(string))
		role := string(claims["role"].(string))
		userId := int(claims["userId"].(float64))

		return &TokenMetadata{
			Expires:  expires,
			Username: username,
			Role:     role,
			UserId:   userId,
		}, nil
	}

	return nil, err
}

func extractToken(c *fiber.Ctx) string {
	bearerToken := c.Get("Authorization")

	onlyToken := strings.Split(bearerToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func verifyRefreshToken(c *fiber.Ctx, refreshToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(refreshToken, jwtKeyFunc)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	configuration := config.New()
	return []byte(configuration.Get("JWT_SECRET_KEY")), nil
}
