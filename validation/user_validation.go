package validation

import (
	"usedeall/exception"
	"usedeall/model/entity"
	"usedeall/model/request"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
)

func UserLoginValidate(ctx *fiber.Ctx, params request.LoginRequest) {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.Username, validation.Required),
		validation.Field(&params.Password, validation.Required),
	)

	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}
}

func UserSaveValidate(ctx *fiber.Ctx, params entity.User) {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.Username, validation.Required),
		validation.Field(&params.Password, validation.Required),
		validation.Field(&params.FirstName, validation.Required),
		validation.Field(&params.LastName, validation.Required),
	)

	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}
}

func UserUpdateValidate(ctx *fiber.Ctx, params entity.User) {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.Id, validation.Required),
		validation.Field(&params.Username, validation.Required),
		validation.Field(&params.Password, validation.Required),
		validation.Field(&params.FirstName, validation.Required),
		validation.Field(&params.LastName, validation.Required),
	)

	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}
}

func RefreshTokenValidate(ctx *fiber.Ctx, params request.RefreshTokenRequest) {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.RefreshToken, validation.Required),
	)

	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}
}
