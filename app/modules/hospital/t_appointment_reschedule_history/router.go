package t_appointment_reschedule_history

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTAppointmentRescheduleHistoryRepository(initializer.DB)
	service := NewTAppointmentRescheduleHistoryService(repo)
	handler := NewTAppointmentRescheduleHistoryHandler(service, validate)

	api.Post("/t-appointment-reschedule-history", handler.TAppointmentRescheduleHistoryCreate)
	api.Put("/t-appointment-reschedule-history", handler.TAppointmentRescheduleHistoryUpdate)
	api.Get("/t-appointment-reschedule-history/:id", handler.TAppointmentRescheduleHistoryIndex)
	api.Get("/t-appointment-reschedule-history", handler.TAppointmentRescheduleHistoryPage)
	api.Delete("/t-appointment-reschedule-history/:id", handler.TAppointmentRescheduleHistoryDelete)

	//api.Get("/generator/t-appointment-reschedule-history/:size", handler.GenerateTAppointmentRescheduleHistory)
}
