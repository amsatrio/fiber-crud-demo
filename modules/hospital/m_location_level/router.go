package m_location_level
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMLocationLevelRepository(initializer.DB)
	service := NewMLocationLevelService(repo)
	handler := NewMLocationLevelHandler(service, validate)
	api := app.Group("/v1/m-location-level")
	api.Post("", handler.MLocationLevelCreate)
	api.Put("", handler.MLocationLevelUpdate)
	api.Get(":id", handler.MLocationLevelIndex)
	api.Get("", handler.MLocationLevelPage)
	api.Delete(":id", handler.MLocationLevelDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-location-level/:size", handler.GenerateMLocationLevel)
}

