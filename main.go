package main

import (
	"7solutions/backend/common/authorization"
	"7solutions/backend/config"
	"7solutions/backend/core/handlers"
	"7solutions/backend/core/middlewares"
	"7solutions/backend/core/repositories"
	"7solutions/backend/core/services"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func init() {
	config.NewAppInitEnvironment()
}

func main() {
	db := config.NewAppDatabase()
	auth := authorization.NewJWT_HS256()
	userRepo := repositories.NewUserRepository(db, "users")

	userSrv := services.NewUserService(auth, userRepo)

	userHand := handlers.NewUserHandler(userSrv)

	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New(config.CorsConfig()))

	go func(userRepo repositories.UserRepository) {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			count, err := userRepo.CountUser()
			if err != nil {
				fmt.Printf("Background task error: Failed to count users: %v\n", err)
			} else {
				fmt.Printf("Background task: Current users: %d\n", count)
			}
			fmt.Printf("Background task: Completed at %s\n", time.Now().Format("2006-01-02 15:04:05"))
		}
	}(userRepo)

	app.Post("/signin", userHand.SignIn)
	app.Post("/user", middlewares.AccessToken, userHand.CreateUser)
	app.Get("/user/:id", middlewares.AccessToken, userHand.GetUserByID)
	app.Get("/users", middlewares.AccessToken, userHand.GetUsers)
	app.Put("/user/:id", middlewares.AccessToken, userHand.UpdateUser)
	app.Delete("/user/:id", middlewares.AccessToken, userHand.DeleteUser)

	app.Listen(":" + config.Env.Port)
}
