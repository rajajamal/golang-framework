package routes

import (
	"github.com/rajajamal/golang-framework/configs"
	"github.com/rajajamal/golang-framework/controllers"
	"github.com/rajajamal/golang-framework/models"
	"github.com/rajajamal/golang-framework/repositories"
	"github.com/rajajamal/golang-framework/services"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/gofiber/jwt/v3"
)

func RegisterUserRoutes(router fiber.Router) {
	service := services.User{
		Repository: repositories.NewUserRepository(configs.Db),
	}
	user := controllers.User{
		Service: service,
	}

	router.Post("/users/login", user.Login)
	router.Post("/users/seed", func(ctx *fiber.Ctx) error {
		service.Create(models.CreateUser{
			Username: "admin",
			Password: "12345",
		})

		return ctx.JSON(map[string]interface{}{
			"message": "Seeding initial successfully",
		})
	})
	router.Use(jwt.New(jwt.Config{
		SigningKey: []byte(configs.Env.SecretKey),
	}))

	router.Post("/users", user.Create)
	router.Put("/users/:id", user.Update)
	router.Put("/users/:id/update-password", user.UpdatePassword)
	router.Delete("/users/:id", user.Delete)
	router.Get("/users", user.GetPaginated)
	router.Get("/users/:id", user.Get)
}
