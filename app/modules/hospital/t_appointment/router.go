package t_appointment

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTAppointmentRepository(initializer.DB)
	service := NewTAppointmentService(repo)
	handler := NewTAppointmentHandler(service, validate)

	api.Post("/t-appointment", handler.TAppointmentCreate)
	api.Put("/t-appointment", handler.TAppointmentUpdate)
	api.Get("/t-appointment/:id", handler.TAppointmentIndex)
	api.Get("/t-appointment", handler.TAppointmentPage)
	api.Delete("/t-appointment/:id", handler.TAppointmentDelete)

	//api.Get("/generator/t-appointment/:size", handler.GenerateTAppointment)
}
