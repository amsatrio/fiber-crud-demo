package t_doctor_office_schedule

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTDoctorOfficeScheduleRepository(initializer.DB)
	service := NewTDoctorOfficeScheduleService(repo)
	handler := NewTDoctorOfficeScheduleHandler(service, validate)

	api.Post("/t-doctor-office-schedule", handler.TDoctorOfficeScheduleCreate)
	api.Put("/t-doctor-office-schedule", handler.TDoctorOfficeScheduleUpdate)
	api.Get("/t-doctor-office-schedule/:id", handler.TDoctorOfficeScheduleIndex)
	api.Get("/t-doctor-office-schedule", handler.TDoctorOfficeSchedulePage)
	api.Delete("/t-doctor-office-schedule/:id", handler.TDoctorOfficeScheduleDelete)

	//api.Get("/generator/t-doctor-office-schedule/:size", handler.GenerateTDoctorOfficeSchedule)
}
