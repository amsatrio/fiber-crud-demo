package t_customer_chat
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTCustomerChatRepository(initializer.DB)
	service := NewTCustomerChatService(repo)
	handler := NewTCustomerChatHandler(service, validate)
	api := app.Group("/v1/t-customer-chat")
	api.Post("", handler.TCustomerChatCreate)
	api.Put("", handler.TCustomerChatUpdate)
	api.Get(":id", handler.TCustomerChatIndex)
	api.Get("", handler.TCustomerChatPage)
	api.Delete(":id", handler.TCustomerChatDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-customer-chat/:size", handler.GenerateTCustomerChat)
}

