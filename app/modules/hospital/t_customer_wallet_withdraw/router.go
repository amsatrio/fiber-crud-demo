package t_customer_wallet_withdraw

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTCustomerWalletWithdrawRepository(initializer.DB)
	service := NewTCustomerWalletWithdrawService(repo)
	handler := NewTCustomerWalletWithdrawHandler(service, validate)

	api.Post("/t-customer-wallet-withdraw", handler.TCustomerWalletWithdrawCreate)
	api.Put("/t-customer-wallet-withdraw", handler.TCustomerWalletWithdrawUpdate)
	api.Get("/t-customer-wallet-withdraw/:id", handler.TCustomerWalletWithdrawIndex)
	api.Get("/t-customer-wallet-withdraw", handler.TCustomerWalletWithdrawPage)
	api.Delete("/t-customer-wallet-withdraw/:id", handler.TCustomerWalletWithdrawDelete)

	//api.Get("/generator/t-customer-wallet-withdraw/:size", handler.GenerateTCustomerWalletWithdraw)
}
