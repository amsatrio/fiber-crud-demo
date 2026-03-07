package m_wallet_default_nominal
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMWalletDefaultNominalRepository(initializer.DB)
	service := NewMWalletDefaultNominalService(repo)
	handler := NewMWalletDefaultNominalHandler(service, validate)
	api := app.Group("/v1/m-wallet-default-nominal")
	api.Post("", handler.MWalletDefaultNominalCreate)
	api.Put("", handler.MWalletDefaultNominalUpdate)
	api.Get(":id", handler.MWalletDefaultNominalIndex)
	api.Get("", handler.MWalletDefaultNominalPage)
	api.Delete(":id", handler.MWalletDefaultNominalDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-wallet-default-nominal/:size", handler.GenerateMWalletDefaultNominal)
}

