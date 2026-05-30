package m_education_level

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMEducationLevelRepository(initializer.DB)
	service := NewMEducationLevelService(repo)
	handler := NewMEducationLevelHandler(service, validate)

	api.Post("/m-education-level", handler.MEducationLevelCreate)
	api.Put("/m-education-level", handler.MEducationLevelUpdate)
	api.Get("/m-education-level/:id", handler.MEducationLevelIndex)
	api.Get("/m-education-level", handler.MEducationLevelPage)
	api.Delete("/m-education-level/:id", handler.MEducationLevelDelete)

	//api.Get("/generator/m-education-level/:size", handler.GenerateMEducationLevel)
}
