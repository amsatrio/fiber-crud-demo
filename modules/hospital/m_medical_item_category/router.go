package m_medical_item_category
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMMedicalItemCategoryRepository(initializer.DB)
	service := NewMMedicalItemCategoryService(repo)
	handler := NewMMedicalItemCategoryHandler(service, validate)
	api := app.Group("/v1/m-medical-item-category")
	api.Post("", handler.MMedicalItemCategoryCreate)
	api.Put("", handler.MMedicalItemCategoryUpdate)
	api.Get(":id", handler.MMedicalItemCategoryIndex)
	api.Get("", handler.MMedicalItemCategoryPage)
	api.Delete(":id", handler.MMedicalItemCategoryDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-medical-item-category/:size", handler.GenerateMMedicalItemCategory)
}

