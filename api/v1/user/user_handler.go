package user

import (
	"anonymous-complaints/internal/shared"
	"anonymous-complaints/internal/user"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service *user.UserService
}

func NewUserHandler(service *user.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Register godoc
// @Summary Registrar un nuevo usuario
// @Description Crea un nuevo usuario con email, contraseña y rol
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body struct{Email string `json:"email"`; Password string `json:"password"`; Role string `json:"role"`} true "Datos del usuario"
// @Success 201 {string} string "User created successfully"
// @Failure 400 {object} fiber.Map
// @Router /api/auth/register [post]
func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	err := h.service.Register(req.Email, req.Password, shared.RoleUser(req.Role))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}

// Login godoc
// @Summary Iniciar sesión
// @Description Retorna un token JWT si las credenciales son válidas
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body struct{Email string `json:"email"`; Password string `json:"password"`} true "Credenciales"
// @Success 200 {object} map[string]string
// @Failure 401 {object} fiber.Map
// @Router /api/auth/login [post]
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	token, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return c.JSON(fiber.Map{"token": token})
}
