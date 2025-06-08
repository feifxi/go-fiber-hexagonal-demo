package handler

import (
	"github.com/gofiber/fiber/v2"

	"chanombude/super-hexagonal/internal/api/rest/dto"
	"chanombude/super-hexagonal/internal/domain"
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := pkg.ValidateDTO.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": err.Error(),
		})
	}
	user := &domain.User{
		Name: body.Name,
		Email: body.Email,
		Password: body.Password,
	}
	err := h.userService.Register(user)
	if err != nil {
		switch err {
		case domain.ErrEmailAlreadyExists:
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
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	user, err := h.userService.GetById(uint(id))
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
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
