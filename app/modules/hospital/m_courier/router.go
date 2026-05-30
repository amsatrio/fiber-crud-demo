package m_courier

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMCourierRepository(initializer.DB)
	service := NewMCourierService(repo)
	handler := NewMCourierHandler(service, validate)

	api.Post("/m-courier", handler.MCourierCreate)
	api.Put("/m-courier", handler.MCourierUpdate)
	api.Get("/m-courier/:id", handler.MCourierIndex)
	api.Get("/m-courier", handler.MCourierPage)
	api.Delete("/m-courier/:id", handler.MCourierDelete)

	//api.Get("/generator/m-courier/:size", handler.GenerateMCourier)
}
