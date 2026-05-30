package health

import "github.com/gofiber/fiber/v3"

func Router(api fiber.Router) {
	hello_world_api := api.Group("/health")
	hello_world_api.Get("/status", Status)
}
