package t_customer_registered_card
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTCustomerRegisteredCardRepository(initializer.DB)
	service := NewTCustomerRegisteredCardService(repo)
	handler := NewTCustomerRegisteredCardHandler(service, validate)
	api := app.Group("/v1/t-customer-registered-card")
	api.Post("", handler.TCustomerRegisteredCardCreate)
	api.Put("", handler.TCustomerRegisteredCardUpdate)
	api.Get(":id", handler.TCustomerRegisteredCardIndex)
	api.Get("", handler.TCustomerRegisteredCardPage)
	api.Delete(":id", handler.TCustomerRegisteredCardDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-customer-registered-card/:size", handler.GenerateTCustomerRegisteredCard)
}

