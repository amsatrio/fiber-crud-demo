package user

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"
	"github.com/amsatrio/fiber-crud-demo/app/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewUserRepository(initializer.DB_TWITTER_CLONE)
	service := NewUserService(repo)
	handler := NewUserHandler(service, validate)

	api.Post("/user", middleware.AuthenticationMiddleware(), handler.UserCreate)
	api.Put("/user", middleware.AuthenticationMiddleware(), handler.UserUpdate)
	api.Get("/user/:id", middleware.AuthenticationMiddleware(), handler.UserIndex)
	api.Get("/user", middleware.AuthenticationMiddleware(), handler.UserPage)
	api.Delete("/user/:id", middleware.AuthenticationMiddleware(), handler.UserDelete)
}
