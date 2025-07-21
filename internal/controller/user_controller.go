package controller

import (
	"github.com/gofiber/fiber/v2"

	"chanombude/super-hexagonal/internal/dto"
	"chanombude/super-hexagonal/internal/model"
	"chanombude/super-hexagonal/internal/port"
	"chanombude/super-hexagonal/pkg/errors"
	"chanombude/super-hexagonal/pkg/validator"
)

type UserController struct {
	userService port.UserService
}

func NewUserController(service port.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

func (h *UserController) RegisterRoutes(app *fiber.App) {
	r := app.Group("/users")
	r.Post("/register", h.Register)
	r.Get("/", h.GetAll)
	r.Get("/:id", h.GetById)
}

func (h *UserController) Register(c *fiber.Ctx) error {
	var body dto.RegisterUserRequest
	if err := c.BodyParser(&body); err != nil {
		return errors.NewValidationError("INVALID_REQUEST", "invalid request body")
	}

	// // Validate the request
    if errs := validator.ValidateStruct(body); errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errs,
		})
    }

    // if errs := validator.ValidateStruct(body); len(errs) > 0 {
    //     return errors.NewValidationError("VALIDATION_FAILED", errs[0].Message)
    // }
	
	user := &model.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	}

	if err := h.userService.Register(user); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *UserController) GetAll(c *fiber.Ctx) error {
	users, err := h.userService.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(users)
}

func (h *UserController) GetById(c *fiber.Ctx) error {
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
