package m_bank

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMBankRepository(initializer.DB)
	service := NewMBankService(repo)
	handler := NewMBankHandler(service, validate)

	api.Post("/m-bank", handler.MBankCreate)
	api.Put("/m-bank", handler.MBankUpdate)
	api.Get("/m-bank/:id", handler.MBankIndex)
	api.Get("/m-bank", handler.MBankPage)
	api.Delete("/m-bank/:id", handler.MBankDelete)

	//api.Get("/generator/m-bank/:size", handler.GenerateMBank)
}
