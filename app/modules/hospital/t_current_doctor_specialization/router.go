package t_current_doctor_specialization

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTCurrentDoctorSpecializationRepository(initializer.DB)
	service := NewTCurrentDoctorSpecializationService(repo)
	handler := NewTCurrentDoctorSpecializationHandler(service, validate)

	api.Post("/t-current-doctor-specialization", handler.TCurrentDoctorSpecializationCreate)
	api.Put("/t-current-doctor-specialization", handler.TCurrentDoctorSpecializationUpdate)
	api.Get("/t-current-doctor-specialization/:id", handler.TCurrentDoctorSpecializationIndex)
	api.Get("/t-current-doctor-specialization", handler.TCurrentDoctorSpecializationPage)
	api.Delete("/t-current-doctor-specialization/:id", handler.TCurrentDoctorSpecializationDelete)

	//api.Get("/generator/t-current-doctor-specialization/:size", handler.GenerateTCurrentDoctorSpecialization)
}
