package m_medical_facility_category

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMMedicalFacilityCategoryRepository(initializer.DB)
	service := NewMMedicalFacilityCategoryService(repo)
	handler := NewMMedicalFacilityCategoryHandler(service, validate)

	api.Post("/m-medical-facility-category", handler.MMedicalFacilityCategoryCreate)
	api.Put("/m-medical-facility-category", handler.MMedicalFacilityCategoryUpdate)
	api.Get("/m-medical-facility-category/:id", handler.MMedicalFacilityCategoryIndex)
	api.Get("/m-medical-facility-category", handler.MMedicalFacilityCategoryPage)
	api.Delete("/m-medical-facility-category/:id", handler.MMedicalFacilityCategoryDelete)

	//api.Get("/generator/m-medical-facility-category/:size", handler.GenerateMMedicalFacilityCategory)
}
