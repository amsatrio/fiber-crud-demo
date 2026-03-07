package t_customer_va_history
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTCustomerVaHistoryRepository(initializer.DB)
	service := NewTCustomerVaHistoryService(repo)
	handler := NewTCustomerVaHistoryHandler(service, validate)
	api := app.Group("/v1/t-customer-va-history")
	api.Post("", handler.TCustomerVaHistoryCreate)
	api.Put("", handler.TCustomerVaHistoryUpdate)
	api.Get(":id", handler.TCustomerVaHistoryIndex)
	api.Get("", handler.TCustomerVaHistoryPage)
	api.Delete(":id", handler.TCustomerVaHistoryDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-customer-va-history/:size", handler.GenerateTCustomerVaHistory)
}

