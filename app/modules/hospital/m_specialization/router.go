package m_specialization

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMSpecializationRepository(initializer.DB)
	service := NewMSpecializationService(repo)
	handler := NewMSpecializationHandler(service, validate)

	api.Post("/m-specialization", handler.MSpecializationCreate)
	api.Put("/m-specialization", handler.MSpecializationUpdate)
	api.Get("/m-specialization/:id", handler.MSpecializationIndex)
	api.Get("/m-specialization", handler.MSpecializationPage)
	api.Delete("/m-specialization/:id", handler.MSpecializationDelete)

	//api.Get("/generator/m-specialization/:size", handler.GenerateMSpecialization)
}
