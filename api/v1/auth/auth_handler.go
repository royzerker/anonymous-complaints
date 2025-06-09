package auth

import (
	"anonymous-complaints/internal/shared"
	"anonymous-complaints/internal/user"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authUseCase *user.UserService
}

func NewAuthController(authUseCase *user.UserService) *AuthController {
	return &AuthController{authUseCase: authUseCase}
}

// @Summary Register a new user
// @Description Register with email, password and role
// @Tags auth
// @Accept json
// @Produce json
// @Param user body request true "User info"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/register [post]
func (ac *AuthController) Register(c *fiber.Ctx) error {
	type request struct {
		Email    string          `json:"email"`
		Password string          `json:"_"`
		Role     shared.RoleUser `json:"role"`
	}

	req := new(request)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	err := ac.authUseCase.Register(req.Email, req.Password, req.Role)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "user created"})
}

// @Summary Login user
// @Description Login and receive JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body request true "Login info"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/login [post]
func (ac *AuthController) Login(c *fiber.Ctx) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	req := new(request)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	token, err := ac.authUseCase.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
	}

	return c.JSON(fiber.Map{"token": token})
}
