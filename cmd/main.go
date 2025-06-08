package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"chanombude/super-hexagonal/config"
	"chanombude/super-hexagonal/internal/controller"
	"chanombude/super-hexagonal/internal/model"
	"chanombude/super-hexagonal/internal/middleware"
	"chanombude/super-hexagonal/internal/repository"
	"chanombude/super-hexagonal/internal/service"
	"chanombude/super-hexagonal/pkg/validator"
	"chanombude/super-hexagonal/pkg/database"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize validator
	validator.Init()

	// Initialize database
	db := database.ConnectDB(cfg.DBDSN)
	db.AutoMigrate(&model.User{})
	fmt.Println("=== Database migration completed ===")

	// Initialize Fiber app
	app := fiber.New()

	// Add middleware
	app.Use(middleware.ErrorHandler())

	// Initialize dependencies
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler :=  controller.NewUserController(userService)

	// Register routes
	userHandler.RegisterRoutes(app)

	// Start server
	fmt.Printf("Server starting on port %s...\n", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
