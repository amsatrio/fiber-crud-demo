package t_customer_va_history

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTCustomerVaHistoryRepository(initializer.DB)
	service := NewTCustomerVaHistoryService(repo)
	handler := NewTCustomerVaHistoryHandler(service, validate)

	api.Post("/t-customer-va-history", handler.TCustomerVaHistoryCreate)
	api.Put("/t-customer-va-history", handler.TCustomerVaHistoryUpdate)
	api.Get("/t-customer-va-history/:id", handler.TCustomerVaHistoryIndex)
	api.Get("/t-customer-va-history", handler.TCustomerVaHistoryPage)
	api.Delete("/t-customer-va-history/:id", handler.TCustomerVaHistoryDelete)

	//api.Get("/generator/t-customer-va-history/:size", handler.GenerateTCustomerVaHistory)
}
