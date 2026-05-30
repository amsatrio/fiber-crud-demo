package t_doctor_office_treatment

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTDoctorOfficeTreatmentRepository(initializer.DB)
	service := NewTDoctorOfficeTreatmentService(repo)
	handler := NewTDoctorOfficeTreatmentHandler(service, validate)

	api.Post("/t-doctor-office-treatment", handler.TDoctorOfficeTreatmentCreate)
	api.Put("/t-doctor-office-treatment", handler.TDoctorOfficeTreatmentUpdate)
	api.Get("/t-doctor-office-treatment/:id", handler.TDoctorOfficeTreatmentIndex)
	api.Get("/t-doctor-office-treatment", handler.TDoctorOfficeTreatmentPage)
	api.Delete("/t-doctor-office-treatment/:id", handler.TDoctorOfficeTreatmentDelete)

	//api.Get("/generator/t-doctor-office-treatment/:size", handler.GenerateTDoctorOfficeTreatment)
}
