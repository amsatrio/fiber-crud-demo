package m_doctor_education
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMDoctorEducationRepository(initializer.DB)
	service := NewMDoctorEducationService(repo)
	handler := NewMDoctorEducationHandler(service, validate)
	api := app.Group("/v1/m-doctor-education")
	api.Post("", handler.MDoctorEducationCreate)
	api.Put("", handler.MDoctorEducationUpdate)
	api.Get(":id", handler.MDoctorEducationIndex)
	api.Get("", handler.MDoctorEducationPage)
	api.Delete(":id", handler.MDoctorEducationDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-doctor-education/:size", handler.GenerateMDoctorEducation)
}

