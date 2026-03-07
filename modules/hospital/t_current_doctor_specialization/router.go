package t_current_doctor_specialization
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTCurrentDoctorSpecializationRepository(initializer.DB)
	service := NewTCurrentDoctorSpecializationService(repo)
	handler := NewTCurrentDoctorSpecializationHandler(service, validate)
	api := app.Group("/v1/t-current-doctor-specialization")
	api.Post("", handler.TCurrentDoctorSpecializationCreate)
	api.Put("", handler.TCurrentDoctorSpecializationUpdate)
	api.Get(":id", handler.TCurrentDoctorSpecializationIndex)
	api.Get("", handler.TCurrentDoctorSpecializationPage)
	api.Delete(":id", handler.TCurrentDoctorSpecializationDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-current-doctor-specialization/:size", handler.GenerateTCurrentDoctorSpecialization)
}

