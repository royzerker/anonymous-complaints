package auth

import (
	"anonymous-complaints/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(router fiber.Router, controller *AuthController) {
	router.Post("/register", controller.Register)
	router.Post("/login", controller.Login)

	protected := router.Group("")
	protected.Use(middleware.JwtProtected())
	// protected.Get("/profile", controller.Profile)
}
