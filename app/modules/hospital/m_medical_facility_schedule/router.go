package m_medical_facility_schedule

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMMedicalFacilityScheduleRepository(initializer.DB)
	service := NewMMedicalFacilityScheduleService(repo)
	handler := NewMMedicalFacilityScheduleHandler(service, validate)

	api.Post("/m-medical-facility-schedule", handler.MMedicalFacilityScheduleCreate)
	api.Put("/m-medical-facility-schedule", handler.MMedicalFacilityScheduleUpdate)
	api.Get("/m-medical-facility-schedule/:id", handler.MMedicalFacilityScheduleIndex)
	api.Get("/m-medical-facility-schedule", handler.MMedicalFacilitySchedulePage)
	api.Delete("/m-medical-facility-schedule/:id", handler.MMedicalFacilityScheduleDelete)

	//api.Get("/generator/m-medical-facility-schedule/:size", handler.GenerateMMedicalFacilitySchedule)
}
