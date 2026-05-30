package t_courier_discount

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTCourierDiscountRepository(initializer.DB)
	service := NewTCourierDiscountService(repo)
	handler := NewTCourierDiscountHandler(service, validate)

	api.Post("/t-courier-discount", handler.TCourierDiscountCreate)
	api.Put("/t-courier-discount", handler.TCourierDiscountUpdate)
	api.Get("/t-courier-discount/:id", handler.TCourierDiscountIndex)
	api.Get("/t-courier-discount", handler.TCourierDiscountPage)
	api.Delete("/t-courier-discount/:id", handler.TCourierDiscountDelete)

	//api.Get("/generator/t-courier-discount/:size", handler.GenerateTCourierDiscount)
}
