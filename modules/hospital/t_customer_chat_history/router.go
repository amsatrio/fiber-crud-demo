package t_customer_chat_history
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTCustomerChatHistoryRepository(initializer.DB)
	service := NewTCustomerChatHistoryService(repo)
	handler := NewTCustomerChatHistoryHandler(service, validate)
	api := app.Group("/v1/t-customer-chat-history")
	api.Post("", handler.TCustomerChatHistoryCreate)
	api.Put("", handler.TCustomerChatHistoryUpdate)
	api.Get(":id", handler.TCustomerChatHistoryIndex)
	api.Get("", handler.TCustomerChatHistoryPage)
	api.Delete(":id", handler.TCustomerChatHistoryDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-customer-chat-history/:size", handler.GenerateTCustomerChatHistory)
}

