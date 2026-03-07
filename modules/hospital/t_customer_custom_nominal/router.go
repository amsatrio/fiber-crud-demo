package t_customer_custom_nominal
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTCustomerCustomNominalRepository(initializer.DB)
	service := NewTCustomerCustomNominalService(repo)
	handler := NewTCustomerCustomNominalHandler(service, validate)
	api := app.Group("/v1/t-customer-custom-nominal")
	api.Post("", handler.TCustomerCustomNominalCreate)
	api.Put("", handler.TCustomerCustomNominalUpdate)
	api.Get(":id", handler.TCustomerCustomNominalIndex)
	api.Get("", handler.TCustomerCustomNominalPage)
	api.Delete(":id", handler.TCustomerCustomNominalDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-customer-custom-nominal/:size", handler.GenerateTCustomerCustomNominal)
}

