package t_customer_va

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTCustomerVaRepository(initializer.DB)
	service := NewTCustomerVaService(repo)
	handler := NewTCustomerVaHandler(service, validate)

	api.Post("/t-customer-va", handler.TCustomerVaCreate)
	api.Put("/t-customer-va", handler.TCustomerVaUpdate)
	api.Get("/t-customer-va/:id", handler.TCustomerVaIndex)
	api.Get("/t-customer-va", handler.TCustomerVaPage)
	api.Delete("/t-customer-va/:id", handler.TCustomerVaDelete)

	//api.Get("/generator/t-customer-va/:size", handler.GenerateTCustomerVa)
}
