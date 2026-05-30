package m_location_level

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMLocationLevelRepository(initializer.DB)
	service := NewMLocationLevelService(repo)
	handler := NewMLocationLevelHandler(service, validate)

	api.Post("/m-location-level", handler.MLocationLevelCreate)
	api.Put("/m-location-level", handler.MLocationLevelUpdate)
	api.Get("/m-location-level/:id", handler.MLocationLevelIndex)
	api.Get("/m-location-level", handler.MLocationLevelPage)
	api.Delete("/m-location-level/:id", handler.MLocationLevelDelete)

	//api.Get("/generator/m-location-level/:size", handler.GenerateMLocationLevel)
}
