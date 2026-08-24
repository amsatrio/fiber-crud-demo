package api

import (
	"net/http"

	"github.com/amsatrio/fiber-crud-demo/cmd"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

func handler() http.HandlerFunc {
	app := fiber.New(cmd.Config())

	cmd.Middleware(app)

	cmd.Routes(app)
	return adaptor.FiberApp(app)
}
