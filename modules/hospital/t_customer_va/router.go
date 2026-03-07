package t_customer_va
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTCustomerVaRepository(initializer.DB)
	service := NewTCustomerVaService(repo)
	handler := NewTCustomerVaHandler(service, validate)
	api := app.Group("/v1/t-customer-va")
	api.Post("", handler.TCustomerVaCreate)
	api.Put("", handler.TCustomerVaUpdate)
	api.Get(":id", handler.TCustomerVaIndex)
	api.Get("", handler.TCustomerVaPage)
	api.Delete(":id", handler.TCustomerVaDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-customer-va/:size", handler.GenerateTCustomerVa)
}

