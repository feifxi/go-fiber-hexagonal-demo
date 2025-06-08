package user

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"

	userDomain "chanombude/super-hexagonal/internal/domain/user"
	lib "chanombude/super-hexagonal/pkg"
)

type UserHandler struct {
	service userDomain.Service
}

func NewUserHandler(s userDomain.Service) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) RegisterRoutes(app *fiber.App) {
	r := app.Group("/users")
	r.Post("/register", h.Register)
	r.Get("/", h.GetAll)
	r.Get("/:id", h.GetById)
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var body RegisterUserRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := lib.ValidateDTO.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "validation failed",
			"details": err.Error(),
		})
	}  
	user := &userDomain.User{
		Name:  body.Name,
		Email: body.Email,
		Password: body.Password,
	}
	if err := h.service.Register(user); err != nil {
		if errors.Is(err, userDomain.ErrEmailAlreadyExists) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": err.Error(),
			})
		} 
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusCreated)
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return c.JSON(users)
}

func (h *UserHandler) GetById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	user, err := h.service.GetById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	if user == nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}
	resUser := UserResponse{
		ID: user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	return c.JSON(resUser)
}