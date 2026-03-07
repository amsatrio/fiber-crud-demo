package m_courier
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMCourierRepository(initializer.DB)
	service := NewMCourierService(repo)
	handler := NewMCourierHandler(service, validate)
	api := app.Group("/v1/m-courier")
	api.Post("", handler.MCourierCreate)
	api.Put("", handler.MCourierUpdate)
	api.Get(":id", handler.MCourierIndex)
	api.Get("", handler.MCourierPage)
	api.Delete(":id", handler.MCourierDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-courier/:size", handler.GenerateMCourier)
}

