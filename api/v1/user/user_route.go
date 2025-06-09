package user

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App, userHandler *UserHandler) {
	api := app.Group("/api")
	auth := api.Group("/user")

	auth.Post("", userHandler.Register)
}
