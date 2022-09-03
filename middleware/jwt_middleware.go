package middleware

import (
	"usedeall/config"
	"usedeall/model"

	"github.com/gofiber/fiber/v2"

	jwtMiddleware "github.com/gofiber/jwt/v3"
)

func JWTProtected() func(*fiber.Ctx) error {
	configuration := config.New()
	config := jwtMiddleware.Config{
		SigningKey:    []byte(configuration.Get("JWT_SECRET_KEY")),
		ContextKey:    "jwt",
		SigningMethod: "HS512",
		ErrorHandler:  jwtError,
	}

	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(model.WebResponse{
		Message:    err.Error(),
		StatusCode: fiber.StatusUnauthorized,
	})
}
