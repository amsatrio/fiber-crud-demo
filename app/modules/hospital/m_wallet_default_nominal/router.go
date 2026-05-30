package m_wallet_default_nominal

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMWalletDefaultNominalRepository(initializer.DB)
	service := NewMWalletDefaultNominalService(repo)
	handler := NewMWalletDefaultNominalHandler(service, validate)

	api.Post("/m-wallet-default-nominal", handler.MWalletDefaultNominalCreate)
	api.Put("/m-wallet-default-nominal", handler.MWalletDefaultNominalUpdate)
	api.Get("/m-wallet-default-nominal/:id", handler.MWalletDefaultNominalIndex)
	api.Get("/m-wallet-default-nominal", handler.MWalletDefaultNominalPage)
	api.Delete("/m-wallet-default-nominal/:id", handler.MWalletDefaultNominalDelete)

	//api.Get("/generator/m-wallet-default-nominal/:size", handler.GenerateMWalletDefaultNominal)
}
