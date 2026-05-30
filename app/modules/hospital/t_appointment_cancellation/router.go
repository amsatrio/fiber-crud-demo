package t_appointment_cancellation

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTAppointmentCancellationRepository(initializer.DB)
	service := NewTAppointmentCancellationService(repo)
	handler := NewTAppointmentCancellationHandler(service, validate)

	api.Post("/t-appointment-cancellation", handler.TAppointmentCancellationCreate)
	api.Put("/t-appointment-cancellation", handler.TAppointmentCancellationUpdate)
	api.Get("/t-appointment-cancellation/:id", handler.TAppointmentCancellationIndex)
	api.Get("/t-appointment-cancellation", handler.TAppointmentCancellationPage)
	api.Delete("/t-appointment-cancellation/:id", handler.TAppointmentCancellationDelete)

	//api.Get("/generator/t-appointment-cancellation/:size", handler.GenerateTAppointmentCancellation)
}
