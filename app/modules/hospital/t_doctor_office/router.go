package t_doctor_office

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTDoctorOfficeRepository(initializer.DB)
	service := NewTDoctorOfficeService(repo)
	handler := NewTDoctorOfficeHandler(service, validate)

	api.Post("/t-doctor-office", handler.TDoctorOfficeCreate)
	api.Put("/t-doctor-office", handler.TDoctorOfficeUpdate)
	api.Get("/t-doctor-office/:id", handler.TDoctorOfficeIndex)
	api.Get("/t-doctor-office", handler.TDoctorOfficePage)
	api.Delete("/t-doctor-office/:id", handler.TDoctorOfficeDelete)

	//api.Get("/generator/t-doctor-office/:size", handler.GenerateTDoctorOffice)
}
