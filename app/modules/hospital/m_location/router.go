package m_location

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMLocationRepository(initializer.DB)
	service := NewMLocationService(repo)
	handler := NewMLocationHandler(service, validate)

	api.Post("/m-location", handler.MLocationCreate)
	api.Put("/m-location", handler.MLocationUpdate)
	api.Get("/m-location/:id", handler.MLocationIndex)
	api.Get("/m-location", handler.MLocationPage)
	api.Delete("/m-location/:id", handler.MLocationDelete)

	//api.Get("/generator/m-location/:size", handler.GenerateMLocation)
}
