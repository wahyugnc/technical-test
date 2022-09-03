package main

import (
	"log"
	"usedeall/config"
	"usedeall/controller"
	"usedeall/exception"
	"usedeall/mapper"
	"usedeall/repository"
	"usedeall/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	configuration := config.New()
	db, err := config.Open(configuration.Get("MARIADB_CONNECTION_STRING"))
	if err != nil {
		log.Fatalf("error opening db: %v", err)
	}

	mapper := mapper.NewMapper()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, mapper)
	userController := controller.NewUserController(userService)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	app.Use(logger.New())

	userController.Route(app)

	err = app.Listen(":3000")
	exception.PanicIfNeeded(err)

}
