package handler

import (
	"github.com/gofiber/fiber/v2"

	"chanombude/super-hexagonal/internal/api/rest/dto"
	"chanombude/super-hexagonal/internal/domain"
	"chanombude/super-hexagonal/internal/pkg/errors"
	"chanombude/super-hexagonal/internal/service"
	"chanombude/super-hexagonal/pkg"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (h *UserHandler) RegisterRoutes(app *fiber.App) {
	r := app.Group("/users")
	r.Post("/register", h.Register)
	r.Get("/", h.GetAll)
	r.Get("/:id", h.GetById)
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var body dto.RegisterUserRequest
	if err := c.BodyParser(&body); err != nil {
		return errors.NewValidationError("INVALID_REQUEST", "invalid request body")
	}

	if err := pkg.ValidateDTO.Struct(body); err != nil {
		return errors.NewValidationError("VALIDATION_FAILED", err.Error())
	}

	user := &domain.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	}

	if err := h.userService.Register(user); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.userService.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(users)
}

func (h *UserHandler) GetById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errors.NewValidationError("INVALID_ID", "invalid user ID")
	}

	user, err := h.userService.GetById(uint(id))
	if err != nil {
		return err
	}

	return c.JSON(dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}
