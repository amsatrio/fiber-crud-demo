package m_specialization
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMSpecializationRepository(initializer.DB)
	service := NewMSpecializationService(repo)
	handler := NewMSpecializationHandler(service, validate)
	api := app.Group("/v1/m-specialization")
	api.Post("", handler.MSpecializationCreate)
	api.Put("", handler.MSpecializationUpdate)
	api.Get(":id", handler.MSpecializationIndex)
	api.Get("", handler.MSpecializationPage)
	api.Delete(":id", handler.MSpecializationDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-specialization/:size", handler.GenerateMSpecialization)
}

