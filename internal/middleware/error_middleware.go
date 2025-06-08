package middleware

import (
	"github.com/gofiber/fiber/v2"
	"chanombude/super-hexagonal/internal/pkg/errors"
)

func ErrorHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err == nil {
			return nil
		}

		// Check if it's our custom error type
		if appErr, ok := err.(*errors.AppError); ok {
			switch appErr.Type {
			case errors.NotFoundError:
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": appErr.Message,
					"code":  appErr.Code,
				})
			case errors.ValidationError:
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": appErr.Message,
					"code":  appErr.Code,
				})
			case errors.ConflictError:
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{
					"error": appErr.Message,
					"code":  appErr.Code,
				})
			case errors.DomainError:
				return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"error": appErr.Message,
					"code":  appErr.Code,
				})
			}
		}

		// Handle other errors
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}
} 