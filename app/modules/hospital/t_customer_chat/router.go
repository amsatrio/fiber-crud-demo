package t_customer_chat

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTCustomerChatRepository(initializer.DB)
	service := NewTCustomerChatService(repo)
	handler := NewTCustomerChatHandler(service, validate)

	api.Post("/t-customer-chat", handler.TCustomerChatCreate)
	api.Put("/t-customer-chat", handler.TCustomerChatUpdate)
	api.Get("/t-customer-chat/:id", handler.TCustomerChatIndex)
	api.Get("/t-customer-chat", handler.TCustomerChatPage)
	api.Delete("/t-customer-chat/:id", handler.TCustomerChatDelete)

	//api.Get("/generator/t-customer-chat/:size", handler.GenerateTCustomerChat)
}
