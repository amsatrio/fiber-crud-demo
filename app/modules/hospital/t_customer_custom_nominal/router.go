package t_customer_custom_nominal

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTCustomerCustomNominalRepository(initializer.DB)
	service := NewTCustomerCustomNominalService(repo)
	handler := NewTCustomerCustomNominalHandler(service, validate)

	api.Post("/t-customer-custom-nominal", handler.TCustomerCustomNominalCreate)
	api.Put("/t-customer-custom-nominal", handler.TCustomerCustomNominalUpdate)
	api.Get("/t-customer-custom-nominal/:id", handler.TCustomerCustomNominalIndex)
	api.Get("/t-customer-custom-nominal", handler.TCustomerCustomNominalPage)
	api.Delete("/t-customer-custom-nominal/:id", handler.TCustomerCustomNominalDelete)

	//api.Get("/generator/t-customer-custom-nominal/:size", handler.GenerateTCustomerCustomNominal)
}
