package t_customer_wallet_top_up
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTCustomerWalletTopUpRepository(initializer.DB)
	service := NewTCustomerWalletTopUpService(repo)
	handler := NewTCustomerWalletTopUpHandler(service, validate)
	api := app.Group("/v1/t-customer-wallet-top-up")
	api.Post("", handler.TCustomerWalletTopUpCreate)
	api.Put("", handler.TCustomerWalletTopUpUpdate)
	api.Get(":id", handler.TCustomerWalletTopUpIndex)
	api.Get("", handler.TCustomerWalletTopUpPage)
	api.Delete(":id", handler.TCustomerWalletTopUpDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-customer-wallet-top-up/:size", handler.GenerateTCustomerWalletTopUp)
}

