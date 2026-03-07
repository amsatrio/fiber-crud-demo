package m_customer
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMCustomerRepository(initializer.DB)
	service := NewMCustomerService(repo)
	handler := NewMCustomerHandler(service, validate)
	api := app.Group("/v1/m-customer")
	api.Post("", handler.MCustomerCreate)
	api.Put("", handler.MCustomerUpdate)
	api.Get(":id", handler.MCustomerIndex)
	api.Get("", handler.MCustomerPage)
	api.Delete(":id", handler.MCustomerDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-customer/:size", handler.GenerateMCustomer)
}

