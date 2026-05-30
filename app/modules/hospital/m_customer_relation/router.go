package m_customer_relation

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMCustomerRelationRepository(initializer.DB)
	service := NewMCustomerRelationService(repo)
	handler := NewMCustomerRelationHandler(service, validate)

	api.Post("/m-customer-relation", handler.MCustomerRelationCreate)
	api.Put("/m-customer-relation", handler.MCustomerRelationUpdate)
	api.Get("/m-customer-relation/:id", handler.MCustomerRelationIndex)
	api.Get("/m-customer-relation", handler.MCustomerRelationPage)
	api.Delete("/m-customer-relation/:id", handler.MCustomerRelationDelete)

	//api.Get("/generator/m-customer-relation/:size", handler.GenerateMCustomerRelation)
}
