package m_doctor

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMDoctorRepository(initializer.DB)
	service := NewMDoctorService(repo)
	handler := NewMDoctorHandler(service, validate)

	api.Post("/m-doctor", handler.MDoctorCreate)
	api.Put("/m-doctor", handler.MDoctorUpdate)
	api.Get("/m-doctor/:id", handler.MDoctorIndex)
	api.Get("/m-doctor", handler.MDoctorPage)
	api.Delete("/m-doctor/:id", handler.MDoctorDelete)

	//api.Get("/generator/m-doctor/:size", handler.GenerateMDoctor)
}
