package m_doctor_education

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMDoctorEducationRepository(initializer.DB)
	service := NewMDoctorEducationService(repo)
	handler := NewMDoctorEducationHandler(service, validate)

	api.Post("/m-doctor-education", handler.MDoctorEducationCreate)
	api.Put("/m-doctor-education", handler.MDoctorEducationUpdate)
	api.Get("/m-doctor-education/:id", handler.MDoctorEducationIndex)
	api.Get("/m-doctor-education", handler.MDoctorEducationPage)
	api.Delete("/m-doctor-education/:id", handler.MDoctorEducationDelete)

	//api.Get("/generator/m-doctor-education/:size", handler.GenerateMDoctorEducation)
}
