package exception

import (
	"usedeall/model"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	_, dataNotFoundError := err.(DataNotFoundError)
	if dataNotFoundError {
		return ctx.Status(fiber.StatusNotFound).JSON(model.WebResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    err.Error(),
		})
	}

	_, dataExsistError := err.(DataExsistError)
	if dataExsistError {
		return ctx.Status(fiber.StatusConflict).JSON(model.WebResponse{
			StatusCode: fiber.StatusConflict,
			Message:    err.Error(),
		})
	}

	_, generalError := err.(GeneralError)
	if generalError {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	_, forbiddenError := err.(ForbiddenError)
	if forbiddenError {
		return ctx.Status(fiber.StatusForbidden).JSON(model.WebResponse{
			StatusCode: fiber.StatusForbidden,
			Message:    err.Error(),
		})
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
		StatusCode: fiber.StatusBadRequest,
		Message:    err.Error(),
	})
}
