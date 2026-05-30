package t_customer_wallet_top_up

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTCustomerWalletTopUpRepository(initializer.DB)
	service := NewTCustomerWalletTopUpService(repo)
	handler := NewTCustomerWalletTopUpHandler(service, validate)

	api.Post("/t-customer-wallet-top-up", handler.TCustomerWalletTopUpCreate)
	api.Put("/t-customer-wallet-top-up", handler.TCustomerWalletTopUpUpdate)
	api.Get("/t-customer-wallet-top-up/:id", handler.TCustomerWalletTopUpIndex)
	api.Get("/t-customer-wallet-top-up", handler.TCustomerWalletTopUpPage)
	api.Delete("/t-customer-wallet-top-up/:id", handler.TCustomerWalletTopUpDelete)

	//api.Get("/generator/t-customer-wallet-top-up/:size", handler.GenerateTCustomerWalletTopUp)
}
