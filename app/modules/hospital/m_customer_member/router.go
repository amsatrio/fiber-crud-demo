package m_customer_member

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMCustomerMemberRepository(initializer.DB)
	service := NewMCustomerMemberService(repo)
	handler := NewMCustomerMemberHandler(service, validate)

	api.Post("/m-customer-member", handler.MCustomerMemberCreate)
	api.Put("/m-customer-member", handler.MCustomerMemberUpdate)
	api.Get("/m-customer-member/:id", handler.MCustomerMemberIndex)
	api.Get("/m-customer-member", handler.MCustomerMemberPage)
	api.Delete("/m-customer-member/:id", handler.MCustomerMemberDelete)

	//api.Get("/generator/m-customer-member/:size", handler.GenerateMCustomerMember)
}
