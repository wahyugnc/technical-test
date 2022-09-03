package controller

import (
	"strconv"
	"usedeall/exception"
	"usedeall/middleware"
	"usedeall/model/entity"
	"usedeall/model/request"
	"usedeall/service"
	"usedeall/utils"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (controller *UserController) Route(app *fiber.App) {
	v1 := app.Group("/v1")
	user := v1.Group("/user")
	user.Get("", middleware.JWTProtected(), controller.FindAll)
	user.Post("", middleware.JWTProtected(), controller.Save)
	user.Post("/login", controller.Login)
	user.Get("/:userid", middleware.JWTProtected(), controller.FindById)
	user.Put("", middleware.JWTProtected(), controller.Update)
	user.Delete("/:userid", middleware.JWTProtected(), controller.Delete)
	user.Post("/refresh-token", controller.RefreshToken)
}

func (controller *UserController) FindAll(ctx *fiber.Ctx) error {
	pageSize, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil {
		panic(exception.GeneralError{
			Message: utils.Message_Invalid_Number,
		})
	}
	pageNumber, err := strconv.Atoi(ctx.Query("page_number"))
	if err != nil {
		panic(exception.GeneralError{
			Message: utils.Message_Invalid_Number,
		})
	}
	response, err := controller.UserService.FindAll(ctx, pageNumber, pageSize)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (controller *UserController) Save(ctx *fiber.Ctx) error {
	var entity entity.User
	err := ctx.BodyParser(&entity)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	response, err := controller.UserService.Save(ctx, &entity)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (controller *UserController) Login(ctx *fiber.Ctx) error {
	var request request.LoginRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	response, err := controller.UserService.UserLogin(ctx, &request)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (controller *UserController) FindById(ctx *fiber.Ctx) error {
	userId, err := strconv.Atoi(ctx.Params("userid"))
	if err != nil {
		panic(exception.GeneralError{
			Message: utils.Message_Invalid_Number,
		})
	}

	response, err := controller.UserService.FindById(ctx, userId)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (controller *UserController) Update(ctx *fiber.Ctx) error {
	var entity entity.User
	err := ctx.BodyParser(&entity)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	response, err := controller.UserService.Update(ctx, &entity)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (controller *UserController) Delete(ctx *fiber.Ctx) error {
	userId, err := strconv.Atoi(ctx.Params("userid"))
	if err != nil {
		panic(exception.GeneralError{
			Message: utils.Message_Invalid_Number,
		})
	}

	response, err := controller.UserService.Delete(ctx, userId)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (controller *UserController) RefreshToken(ctx *fiber.Ctx) error {
	var request request.RefreshTokenRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	response, err := controller.UserService.RefreshToken(ctx, &request)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
