package auth

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"
	"github.com/amsatrio/fiber-crud-demo/app/middleware"
	"github.com/amsatrio/fiber-crud-demo/app/modules/twitter_clone/user"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewAuthRepository(initializer.DB_TWITTER_CLONE)
	mUserRepository := user.NewUserRepository(initializer.DB_TWITTER_CLONE)
	service := NewAuthService(repo, mUserRepository)
	handler := NewAuthHandler(service, validate)

	api.Post("/auth/login", handler.AuthLogin)
	api.Post("/auth/register", handler.AuthRegister)
	api.Post("/auth/refresh-token", handler.AuthRefreshToken)
	api.Post("/auth/reset-password", middleware.AuthenticationMiddleware(), handler.AuthResetPassword)

	api.Get("/auth/public", middleware.AuthenticationMiddleware(), handler.AuthPublic)
}
