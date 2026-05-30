package m_biodata

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMBiodataRepository(initializer.DB)
	service := NewMBiodataService(repo)
	handler := NewMBiodataHandler(service, validate)

	api.Post("/m-biodata", handler.MBiodataCreate)
	api.Put("/m-biodata", handler.MBiodataUpdate)
	api.Get("/m-biodata/:id", handler.MBiodataIndex)
	api.Get("/m-biodata", handler.MBiodataPage)
	api.Delete("/m-biodata/:id", handler.MBiodataDelete)

	//api.Get("/generator/m-biodata/:size", handler.GenerateMBiodata)
}
