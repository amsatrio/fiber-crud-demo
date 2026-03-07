package m_education_level
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMEducationLevelRepository(initializer.DB)
	service := NewMEducationLevelService(repo)
	handler := NewMEducationLevelHandler(service, validate)
	api := app.Group("/v1/m-education-level")
	api.Post("", handler.MEducationLevelCreate)
	api.Put("", handler.MEducationLevelUpdate)
	api.Get(":id", handler.MEducationLevelIndex)
	api.Get("", handler.MEducationLevelPage)
	api.Delete(":id", handler.MEducationLevelDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-education-level/:size", handler.GenerateMEducationLevel)
}

