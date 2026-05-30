package m_payment_method

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMPaymentMethodRepository(initializer.DB)
	service := NewMPaymentMethodService(repo)
	handler := NewMPaymentMethodHandler(service, validate)

	api.Post("/m-payment-method", handler.MPaymentMethodCreate)
	api.Put("/m-payment-method", handler.MPaymentMethodUpdate)
	api.Get("/m-payment-method/:id", handler.MPaymentMethodIndex)
	api.Get("/m-payment-method", handler.MPaymentMethodPage)
	api.Delete("/m-payment-method/:id", handler.MPaymentMethodDelete)

	//api.Get("/generator/m-payment-method/:size", handler.GenerateMPaymentMethod)
}
