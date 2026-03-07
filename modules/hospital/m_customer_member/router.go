package m_customer_member
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMCustomerMemberRepository(initializer.DB)
	service := NewMCustomerMemberService(repo)
	handler := NewMCustomerMemberHandler(service, validate)
	api := app.Group("/v1/m-customer-member")
	api.Post("", handler.MCustomerMemberCreate)
	api.Put("", handler.MCustomerMemberUpdate)
	api.Get(":id", handler.MCustomerMemberIndex)
	api.Get("", handler.MCustomerMemberPage)
	api.Delete(":id", handler.MCustomerMemberDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-customer-member/:size", handler.GenerateMCustomerMember)
}

