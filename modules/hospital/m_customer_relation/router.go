package m_customer_relation
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMCustomerRelationRepository(initializer.DB)
	service := NewMCustomerRelationService(repo)
	handler := NewMCustomerRelationHandler(service, validate)
	api := app.Group("/v1/m-customer-relation")
	api.Post("", handler.MCustomerRelationCreate)
	api.Put("", handler.MCustomerRelationUpdate)
	api.Get(":id", handler.MCustomerRelationIndex)
	api.Get("", handler.MCustomerRelationPage)
	api.Delete(":id", handler.MCustomerRelationDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-customer-relation/:size", handler.GenerateMCustomerRelation)
}

