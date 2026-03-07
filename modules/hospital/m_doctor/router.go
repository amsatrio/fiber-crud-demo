package m_doctor
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMDoctorRepository(initializer.DB)
	service := NewMDoctorService(repo)
	handler := NewMDoctorHandler(service, validate)
	api := app.Group("/v1/m-doctor")
	api.Post("", handler.MDoctorCreate)
	api.Put("", handler.MDoctorUpdate)
	api.Get(":id", handler.MDoctorIndex)
	api.Get("", handler.MDoctorPage)
	api.Delete(":id", handler.MDoctorDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-doctor/:size", handler.GenerateMDoctor)
}

