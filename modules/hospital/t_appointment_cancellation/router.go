package t_appointment_cancellation
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTAppointmentCancellationRepository(initializer.DB)
	service := NewTAppointmentCancellationService(repo)
	handler := NewTAppointmentCancellationHandler(service, validate)
	api := app.Group("/v1/t-appointment-cancellation")
	api.Post("", handler.TAppointmentCancellationCreate)
	api.Put("", handler.TAppointmentCancellationUpdate)
	api.Get(":id", handler.TAppointmentCancellationIndex)
	api.Get("", handler.TAppointmentCancellationPage)
	api.Delete(":id", handler.TAppointmentCancellationDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-appointment-cancellation/:size", handler.GenerateTAppointmentCancellation)
}

