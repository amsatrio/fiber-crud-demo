package auth

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"
	"github.com/amsatrio/fiber-crud-demo/app/middleware"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_user"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMAdminRepository(initializer.DB)
	mUserRepository := m_user.NewMUserRepository(initializer.DB)
	service := NewAuthService(repo, mUserRepository)
	handler := NewAuthHandler(service, validate)

	api.Post("/auth/login", handler.AuthLogin)
	api.Post("/auth/register", handler.AuthRegister)
	api.Post("/auth/refresh-token", handler.AuthRefreshToken)
	api.Post("/auth/reset-password", middleware.AuthenticationMiddleware(), handler.AuthResetPassword)

	api.Get("/auth/dokter", middleware.AuthenticationMiddleware(), middleware.AuthorizationMiddleware(1, 2), handler.AuthDokter)
	api.Get("/auth/faskes", middleware.AuthenticationMiddleware(), middleware.AuthorizationMiddleware(1, 4), handler.AuthFaskes)
	api.Get("/auth/admin", middleware.AuthenticationMiddleware(), middleware.AuthorizationMiddleware(1), handler.AuthAdmin)
	api.Get("/auth/public", handler.AuthPublic)
	api.Get("/auth/pasien", middleware.AuthenticationMiddleware(), middleware.AuthorizationMiddleware(1, 3), handler.AuthPasien)
}
