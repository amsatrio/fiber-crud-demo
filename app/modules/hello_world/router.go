package hello_world

import "github.com/gofiber/fiber/v3"

func Router(api fiber.Router){
	hello_world_api := api.Group("/hello-world")
	hello_world_api.Get("", HelloWorld)
	hello_world_api.Get("/path/:message", HelloWorldPath)
	hello_world_api.Get("/query", HelloWorldQuery)
	hello_world_api.Post("/payload", HelloWorldPayload)
	hello_world_api.Get("/error/:type", HelloWorldError)
}