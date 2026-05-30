package m_customer

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMCustomerRepository(initializer.DB)
	service := NewMCustomerService(repo)
	handler := NewMCustomerHandler(service, validate)

	api.Post("/m-customer", handler.MCustomerCreate)
	api.Put("/m-customer", handler.MCustomerUpdate)
	api.Get("/m-customer/:id", handler.MCustomerIndex)
	api.Get("/m-customer", handler.MCustomerPage)
	api.Delete("/m-customer/:id", handler.MCustomerDelete)

	//api.Get("/generator/m-customer/:size", handler.GenerateMCustomer)
}
