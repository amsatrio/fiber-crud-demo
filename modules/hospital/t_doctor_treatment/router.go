package t_doctor_treatment
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTDoctorTreatmentRepository(initializer.DB)
	service := NewTDoctorTreatmentService(repo)
	handler := NewTDoctorTreatmentHandler(service, validate)
	api := app.Group("/v1/t-doctor-treatment")
	api.Post("", handler.TDoctorTreatmentCreate)
	api.Put("", handler.TDoctorTreatmentUpdate)
	api.Get(":id", handler.TDoctorTreatmentIndex)
	api.Get("", handler.TDoctorTreatmentPage)
	api.Delete(":id", handler.TDoctorTreatmentDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-doctor-treatment/:size", handler.GenerateTDoctorTreatment)
}

