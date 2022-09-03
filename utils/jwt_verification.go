package utils

import (
	"time"
	"usedeall/exception"

	"github.com/gofiber/fiber/v2"
)

func JwtVerification(ctx *fiber.Ctx) {
	now := time.Now().Unix()

	claims, err := ExtractTokenMetadata(ctx)
	if err != nil {
		panic(exception.UnauthorizedError{
			Message: err.Error(),
		})
	}

	expires := claims.Expires
	if now > expires {
		message := Message_TokenExceptionMessage
		panic(exception.UnauthorizedError{
			Message: message,
		})
	}

}

func JwtVerificationWithClaim(ctx *fiber.Ctx) *TokenMetadata {
	now := time.Now().Unix()

	claims, err := ExtractTokenMetadata(ctx)
	if err != nil {
		panic(exception.UnauthorizedError{
			Message: err.Error(),
		})
	}

	expires := claims.Expires
	if now > expires {
		message := Message_TokenExceptionMessage
		panic(exception.UnauthorizedError{
			Message: message,
		})
	}

	return claims
}

func JwtVerificationWithClaimEmail(ctx *fiber.Ctx, refreshToken string) *TokenMetadata {
	now := time.Now().Unix()

	claims, err := ExtractRefreshTokenMetadata(ctx, refreshToken)
	if err != nil {
		panic(exception.RefreshTokenError{
			Message: err.Error(),
		})
	}

	expires := claims.Expires
	if now > expires {
		panic(exception.RefreshTokenError{
			Message: Message_RefreshToken,
		})
	}

	return claims
}
