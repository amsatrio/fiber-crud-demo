package t_customer_wallet
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTCustomerWalletRepository(initializer.DB)
	service := NewTCustomerWalletService(repo)
	handler := NewTCustomerWalletHandler(service, validate)
	api := app.Group("/v1/t-customer-wallet")
	api.Post("", handler.TCustomerWalletCreate)
	api.Put("", handler.TCustomerWalletUpdate)
	api.Get(":id", handler.TCustomerWalletIndex)
	api.Get("", handler.TCustomerWalletPage)
	api.Delete(":id", handler.TCustomerWalletDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-customer-wallet/:size", handler.GenerateTCustomerWallet)
}

