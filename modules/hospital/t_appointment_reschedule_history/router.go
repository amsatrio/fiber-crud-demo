package t_appointment_reschedule_history
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTAppointmentRescheduleHistoryRepository(initializer.DB)
	service := NewTAppointmentRescheduleHistoryService(repo)
	handler := NewTAppointmentRescheduleHistoryHandler(service, validate)
	api := app.Group("/v1/t-appointment-reschedule-history")
	api.Post("", handler.TAppointmentRescheduleHistoryCreate)
	api.Put("", handler.TAppointmentRescheduleHistoryUpdate)
	api.Get(":id", handler.TAppointmentRescheduleHistoryIndex)
	api.Get("", handler.TAppointmentRescheduleHistoryPage)
	api.Delete(":id", handler.TAppointmentRescheduleHistoryDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-appointment-reschedule-history/:size", handler.GenerateTAppointmentRescheduleHistory)
}

