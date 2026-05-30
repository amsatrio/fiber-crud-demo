package t_customer_wallet

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTCustomerWalletRepository(initializer.DB)
	service := NewTCustomerWalletService(repo)
	handler := NewTCustomerWalletHandler(service, validate)

	api.Post("/t-customer-wallet", handler.TCustomerWalletCreate)
	api.Put("/t-customer-wallet", handler.TCustomerWalletUpdate)
	api.Get("/t-customer-wallet/:id", handler.TCustomerWalletIndex)
	api.Get("/t-customer-wallet", handler.TCustomerWalletPage)
	api.Delete("/t-customer-wallet/:id", handler.TCustomerWalletDelete)

	//api.Get("/generator/t-customer-wallet/:size", handler.GenerateTCustomerWallet)
}
