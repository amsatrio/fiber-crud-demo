package m_biodata_address

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMBiodataAddressRepository(initializer.DB)
	service := NewMBiodataAddressService(repo)
	handler := NewMBiodataAddressHandler(service, validate)

	api.Post("/m-biodata-address", handler.MBiodataAddressCreate)
	api.Put("/m-biodata-address", handler.MBiodataAddressUpdate)
	api.Get("/m-biodata-address/:id", handler.MBiodataAddressIndex)
	api.Get("/m-biodata-address", handler.MBiodataAddressPage)
	api.Delete("/m-biodata-address/:id", handler.MBiodataAddressDelete)

	//api.Get("/generator/m-biodata-address/:size", handler.GenerateMBiodataAddress)
}
