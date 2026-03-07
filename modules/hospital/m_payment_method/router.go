package m_payment_method
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMPaymentMethodRepository(initializer.DB)
	service := NewMPaymentMethodService(repo)
	handler := NewMPaymentMethodHandler(service, validate)
	api := app.Group("/v1/m-payment-method")
	api.Post("", handler.MPaymentMethodCreate)
	api.Put("", handler.MPaymentMethodUpdate)
	api.Get(":id", handler.MPaymentMethodIndex)
	api.Get("", handler.MPaymentMethodPage)
	api.Delete(":id", handler.MPaymentMethodDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-payment-method/:size", handler.GenerateMPaymentMethod)
}

