package t_reset_password

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTResetPasswordRepository(initializer.DB)
	service := NewTResetPasswordService(repo)
	handler := NewTResetPasswordHandler(service, validate)

	api.Post("/t-reset-password", handler.TResetPasswordCreate)
	api.Put("/t-reset-password", handler.TResetPasswordUpdate)
	api.Get("/t-reset-password/:id", handler.TResetPasswordIndex)
	api.Get("/t-reset-password", handler.TResetPasswordPage)
	api.Delete("/t-reset-password/:id", handler.TResetPasswordDelete)

	//api.Get("/generator/t-reset-password/:size", handler.GenerateTResetPassword)
}
