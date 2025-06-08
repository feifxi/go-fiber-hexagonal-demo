package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"chanombude/super-hexagonal/internal/api/rest/dto"
	"chanombude/super-hexagonal/internal/domain/errors"
	"chanombude/super-hexagonal/internal/domain/ports/primary"
	"chanombude/super-hexagonal/pkg"
)

type UserHandler struct {
	userService primary.UserService
}

func NewUserHandler(service primary.UserService) *UserHandler {
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
	var req dto.RegisterUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := pkg.ValidateDTO.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": err.Error(),
		})
	}

	err := h.userService.Register(req.Name, req.Email, req.Password)
	if err != nil {
		switch err {
		case errors.ErrEmailAlreadyExists:
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": err.Error(),
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.userService.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(users)
}

func (h *UserHandler) GetById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	user, err := h.userService.GetById(uint(id))
	if err != nil {
		switch err {
		case errors.ErrUserNotFound:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
	}

	return c.JSON(dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}
