package m_courier_type
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMCourierTypeRepository(initializer.DB)
	service := NewMCourierTypeService(repo)
	handler := NewMCourierTypeHandler(service, validate)
	api := app.Group("/v1/m-courier-type")
	api.Post("", handler.MCourierTypeCreate)
	api.Put("", handler.MCourierTypeUpdate)
	api.Get(":id", handler.MCourierTypeIndex)
	api.Get("", handler.MCourierTypePage)
	api.Delete(":id", handler.MCourierTypeDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-courier-type/:size", handler.GenerateMCourierType)
}

