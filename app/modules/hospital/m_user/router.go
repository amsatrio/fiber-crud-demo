package m_user

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMUserRepository(initializer.DB)
	service := NewMUserService(repo)
	handler := NewMUserHandler(service, validate)

	api.Post("/m-user", handler.MUserCreate)
	api.Put("/m-user", handler.MUserUpdate)
	api.Get("/m-user/:id", handler.MUserIndex)
	api.Get("/m-user", handler.MUserPage)
	api.Delete("/m-user/:id", handler.MUserDelete)

	//api.Get("/generator/m-user/:size", handler.GenerateMUser)
}
