package t_medical_item_purchase_detail
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTMedicalItemPurchaseDetailRepository(initializer.DB)
	service := NewTMedicalItemPurchaseDetailService(repo)
	handler := NewTMedicalItemPurchaseDetailHandler(service, validate)
	api := app.Group("/v1/t-medical-item-purchase-detail")
	api.Post("", handler.TMedicalItemPurchaseDetailCreate)
	api.Put("", handler.TMedicalItemPurchaseDetailUpdate)
	api.Get(":id", handler.TMedicalItemPurchaseDetailIndex)
	api.Get("", handler.TMedicalItemPurchaseDetailPage)
	api.Delete(":id", handler.TMedicalItemPurchaseDetailDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-medical-item-purchase-detail/:size", handler.GenerateTMedicalItemPurchaseDetail)
}

