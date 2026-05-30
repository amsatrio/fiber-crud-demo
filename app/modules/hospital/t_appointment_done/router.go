package t_appointment_done

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTAppointmentDoneRepository(initializer.DB)
	service := NewTAppointmentDoneService(repo)
	handler := NewTAppointmentDoneHandler(service, validate)

	api.Post("/t-appointment-done", handler.TAppointmentDoneCreate)
	api.Put("/t-appointment-done", handler.TAppointmentDoneUpdate)
	api.Get("/t-appointment-done/:id", handler.TAppointmentDoneIndex)
	api.Get("/t-appointment-done", handler.TAppointmentDonePage)
	api.Delete("/t-appointment-done/:id", handler.TAppointmentDoneDelete)

	//api.Get("/generator/t-appointment-done/:size", handler.GenerateTAppointmentDone)
}
