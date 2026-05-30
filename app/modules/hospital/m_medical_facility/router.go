package m_medical_facility

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMMedicalFacilityRepository(initializer.DB)
	service := NewMMedicalFacilityService(repo)
	handler := NewMMedicalFacilityHandler(service, validate)

	api.Post("/m-medical-facility", handler.MMedicalFacilityCreate)
	api.Put("/m-medical-facility", handler.MMedicalFacilityUpdate)
	api.Get("/m-medical-facility/:id", handler.MMedicalFacilityIndex)
	api.Get("/m-medical-facility", handler.MMedicalFacilityPage)
	api.Delete("/m-medical-facility/:id", handler.MMedicalFacilityDelete)

	//api.Get("/generator/m-medical-facility/:size", handler.GenerateMMedicalFacility)
}
