package t_customer_wallet_withdraw
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTCustomerWalletWithdrawRepository(initializer.DB)
	service := NewTCustomerWalletWithdrawService(repo)
	handler := NewTCustomerWalletWithdrawHandler(service, validate)
	api := app.Group("/v1/t-customer-wallet-withdraw")
	api.Post("", handler.TCustomerWalletWithdrawCreate)
	api.Put("", handler.TCustomerWalletWithdrawUpdate)
	api.Get(":id", handler.TCustomerWalletWithdrawIndex)
	api.Get("", handler.TCustomerWalletWithdrawPage)
	api.Delete(":id", handler.TCustomerWalletWithdrawDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-customer-wallet-withdraw/:size", handler.GenerateTCustomerWalletWithdraw)
}

