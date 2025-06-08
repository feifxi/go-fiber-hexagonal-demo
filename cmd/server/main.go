package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"chanombude/super-hexagonal/config"
	"chanombude/super-hexagonal/internal/api/rest/handler"
	"chanombude/super-hexagonal/internal/domain"
	"chanombude/super-hexagonal/internal/repository"
	"chanombude/super-hexagonal/internal/service"
	"chanombude/super-hexagonal/pkg"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	db := pkg.ConnectDB(cfg.DBDSN)
	db.AutoMigrate(&domain.User{}, &domain.Product{}, &domain.Order{}, &domain.OrderItem{})
	fmt.Println("=== Database migration completed ===")

	// Initialize Fiber app
	app := fiber.New()

	// Initialize dependencies
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// Register routes
	userHandler.RegisterRoutes(app)
	productHandler.RegisterRoutes(app)

	// Start server
	fmt.Printf("Server starting on port %s...\n", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
