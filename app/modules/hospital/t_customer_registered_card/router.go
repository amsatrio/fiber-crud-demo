package t_customer_registered_card

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTCustomerRegisteredCardRepository(initializer.DB)
	service := NewTCustomerRegisteredCardService(repo)
	handler := NewTCustomerRegisteredCardHandler(service, validate)

	api.Post("/t-customer-registered-card", handler.TCustomerRegisteredCardCreate)
	api.Put("/t-customer-registered-card", handler.TCustomerRegisteredCardUpdate)
	api.Get("/t-customer-registered-card/:id", handler.TCustomerRegisteredCardIndex)
	api.Get("/t-customer-registered-card", handler.TCustomerRegisteredCardPage)
	api.Delete("/t-customer-registered-card/:id", handler.TCustomerRegisteredCardDelete)

	//api.Get("/generator/t-customer-registered-card/:size", handler.GenerateTCustomerRegisteredCard)
}
