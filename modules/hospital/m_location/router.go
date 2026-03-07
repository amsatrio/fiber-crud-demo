package m_location
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMLocationRepository(initializer.DB)
	service := NewMLocationService(repo)
	handler := NewMLocationHandler(service, validate)
	api := app.Group("/v1/m-location")
	api.Post("", handler.MLocationCreate)
	api.Put("", handler.MLocationUpdate)
	api.Get(":id", handler.MLocationIndex)
	api.Get("", handler.MLocationPage)
	api.Delete(":id", handler.MLocationDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-location/:size", handler.GenerateMLocation)
}

