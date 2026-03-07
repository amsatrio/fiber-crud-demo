package t_medical_item_purchase
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTMedicalItemPurchaseRepository(initializer.DB)
	service := NewTMedicalItemPurchaseService(repo)
	handler := NewTMedicalItemPurchaseHandler(service, validate)
	api := app.Group("/v1/t-medical-item-purchase")
	api.Post("", handler.TMedicalItemPurchaseCreate)
	api.Put("", handler.TMedicalItemPurchaseUpdate)
	api.Get(":id", handler.TMedicalItemPurchaseIndex)
	api.Get("", handler.TMedicalItemPurchasePage)
	api.Delete(":id", handler.TMedicalItemPurchaseDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-medical-item-purchase/:size", handler.GenerateTMedicalItemPurchase)
}

