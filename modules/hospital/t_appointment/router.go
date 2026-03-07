package t_appointment
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTAppointmentRepository(initializer.DB)
	service := NewTAppointmentService(repo)
	handler := NewTAppointmentHandler(service, validate)
	api := app.Group("/v1/t-appointment")
	api.Post("", handler.TAppointmentCreate)
	api.Put("", handler.TAppointmentUpdate)
	api.Get(":id", handler.TAppointmentIndex)
	api.Get("", handler.TAppointmentPage)
	api.Delete(":id", handler.TAppointmentDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-appointment/:size", handler.GenerateTAppointment)
}

