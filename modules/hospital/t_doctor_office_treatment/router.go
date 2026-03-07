package t_doctor_office_treatment
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTDoctorOfficeTreatmentRepository(initializer.DB)
	service := NewTDoctorOfficeTreatmentService(repo)
	handler := NewTDoctorOfficeTreatmentHandler(service, validate)
	api := app.Group("/v1/t-doctor-office-treatment")
	api.Post("", handler.TDoctorOfficeTreatmentCreate)
	api.Put("", handler.TDoctorOfficeTreatmentUpdate)
	api.Get(":id", handler.TDoctorOfficeTreatmentIndex)
	api.Get("", handler.TDoctorOfficeTreatmentPage)
	api.Delete(":id", handler.TDoctorOfficeTreatmentDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-doctor-office-treatment/:size", handler.GenerateTDoctorOfficeTreatment)
}

