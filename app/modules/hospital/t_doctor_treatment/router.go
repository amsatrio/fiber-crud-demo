package t_doctor_treatment

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTDoctorTreatmentRepository(initializer.DB)
	service := NewTDoctorTreatmentService(repo)
	handler := NewTDoctorTreatmentHandler(service, validate)

	api.Post("/t-doctor-treatment", handler.TDoctorTreatmentCreate)
	api.Put("/t-doctor-treatment", handler.TDoctorTreatmentUpdate)
	api.Get("/t-doctor-treatment/:id", handler.TDoctorTreatmentIndex)
	api.Get("/t-doctor-treatment", handler.TDoctorTreatmentPage)
	api.Delete("/t-doctor-treatment/:id", handler.TDoctorTreatmentDelete)

	//api.Get("/generator/t-doctor-treatment/:size", handler.GenerateTDoctorTreatment)
}
