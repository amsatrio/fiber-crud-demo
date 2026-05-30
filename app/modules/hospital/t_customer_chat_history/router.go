package t_customer_chat_history

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTCustomerChatHistoryRepository(initializer.DB)
	service := NewTCustomerChatHistoryService(repo)
	handler := NewTCustomerChatHistoryHandler(service, validate)

	api.Post("/t-customer-chat-history", handler.TCustomerChatHistoryCreate)
	api.Put("/t-customer-chat-history", handler.TCustomerChatHistoryUpdate)
	api.Get("/t-customer-chat-history/:id", handler.TCustomerChatHistoryIndex)
	api.Get("/t-customer-chat-history", handler.TCustomerChatHistoryPage)
	api.Delete("/t-customer-chat-history/:id", handler.TCustomerChatHistoryDelete)

	//api.Get("/generator/t-customer-chat-history/:size", handler.GenerateTCustomerChatHistory)
}
