package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	config "chanombude/super-hexagonal/config"
	userDomain "chanombude/super-hexagonal/internal/domain/user"
	lib "chanombude/super-hexagonal/pkg"

	userRest "chanombude/super-hexagonal/internal/handler/rest/user"
	userRepo "chanombude/super-hexagonal/internal/repository"
	userSvc "chanombude/super-hexagonal/internal/service"
)

func main() {
	cfg := config.Load()

	db := lib.ConnectDB(cfg.DBDSN)
	db.AutoMigrate(&userDomain.User{})
	fmt.Println("=== Migrate Successs ===")

	app := fiber.New()

	// User
	userRepository := userRepo.NewUserRepository(db)
	userService := userSvc.NewUserService(userRepository)
	userHandler := userRest.NewUserHandler(userService)
	userHandler.RegisterRoutes(app)	

	app.Listen(":" + cfg.Port)
}