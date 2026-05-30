package m_courier_type

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMCourierTypeRepository(initializer.DB)
	service := NewMCourierTypeService(repo)
	handler := NewMCourierTypeHandler(service, validate)

	api.Post("/m-courier-type", handler.MCourierTypeCreate)
	api.Put("/m-courier-type", handler.MCourierTypeUpdate)
	api.Get("/m-courier-type/:id", handler.MCourierTypeIndex)
	api.Get("/m-courier-type", handler.MCourierTypePage)
	api.Delete("/m-courier-type/:id", handler.MCourierTypeDelete)

	//api.Get("/generator/m-courier-type/:size", handler.GenerateMCourierType)
}
