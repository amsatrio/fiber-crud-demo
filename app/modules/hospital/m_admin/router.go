package m_admin

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMAdminRepository(initializer.DB)
	service := NewMAdminService(repo)
	handler := NewMAdminHandler(service, validate)

	api.Post("/m-admin", handler.MAdminCreate)
	api.Put("/m-admin", handler.MAdminUpdate)
	api.Get("/m-admin/:id", handler.MAdminIndex)
	api.Get("/m-admin", handler.MAdminPage)
	api.Delete("/m-admin/:id", handler.MAdminDelete)

	//api.Get("/generator/m-admin/:size", handler.GenerateMAdmin)
}
