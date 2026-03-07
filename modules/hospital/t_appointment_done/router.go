package t_appointment_done
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTAppointmentDoneRepository(initializer.DB)
	service := NewTAppointmentDoneService(repo)
	handler := NewTAppointmentDoneHandler(service, validate)
	api := app.Group("/v1/t-appointment-done")
	api.Post("", handler.TAppointmentDoneCreate)
	api.Put("", handler.TAppointmentDoneUpdate)
	api.Get(":id", handler.TAppointmentDoneIndex)
	api.Get("", handler.TAppointmentDonePage)
	api.Delete(":id", handler.TAppointmentDoneDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-appointment-done/:size", handler.GenerateTAppointmentDone)
}

