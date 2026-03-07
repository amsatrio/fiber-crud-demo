package m_medical_item
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMMedicalItemRepository(initializer.DB)
	service := NewMMedicalItemService(repo)
	handler := NewMMedicalItemHandler(service, validate)
	api := app.Group("/v1/m-medical-item")
	api.Post("", handler.MMedicalItemCreate)
	api.Put("", handler.MMedicalItemUpdate)
	api.Get(":id", handler.MMedicalItemIndex)
	api.Get("", handler.MMedicalItemPage)
	api.Delete(":id", handler.MMedicalItemDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-medical-item/:size", handler.GenerateMMedicalItem)
}

