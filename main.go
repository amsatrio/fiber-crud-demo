package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"runtime"

	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
	"github.com/amsatrio/fiber-crud-demo/app/initializer"
	"github.com/amsatrio/fiber-crud-demo/app/middleware"
	"github.com/amsatrio/fiber-crud-demo/app/modules/health"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hello_world"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital"
	_ "github.com/amsatrio/fiber-crud-demo/docs"

	"github.com/amsatrio/fiber-crud-demo/app/util"

	"github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/html/v2"
)

func init() {
	initializer.LoadEnvironmentVariables()
	initializer.LoggerInit()
	initializer.InitializeDatabase()
}

func config() fiber.Config {
	htmlEngine := html.New("./web/templates", ".html")
	return fiber.Config{
		// Prefork:               true,
		CaseSensitive: true,
		StrictRouting: true,
		BodyLimit:     4 * 1024 * 1024,
		// DisableStartupMessage: true,
		ServerHeader: "Fiber Audio Management",
		AppName:      "Fiber Audio Management v0.0.1",
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: func(c fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			res := &response.Response{}
			res.ErrMessage(c.Path(), code, err.Error())

			return c.Status(code).JSON(res)
		},
		Views:       htmlEngine,
		ViewsLayout: "layouts/main",
	}
}

func main() {
	runtime.GOMAXPROCS(1)

	app := fiber.New(config())

	// ### Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // v3 expects a slice of strings for origins
		AllowHeaders: []string{"*"},
	}))
	// app.Use(cache.New())
	app.Use(recover.New())
	app.Use(middleware.LoggerMiddleware)

	// app.Get("/swagger/*", swagger.HandlerDefault) // default
	app.Get("/swagger/*", swaggo.HandlerDefault)

	app.Get("/public/*", static.New("./web/public"))

	// ### Routes
	routes(app)

	port := os.Getenv("SERVER_PORT")
	host := os.Getenv("SERVER_HOST")
	util.Log("INFO", "root", "main", "listen and serve on "+host+" port "+port)

	// ### Run
	log.Fatal(app.Listen(host + ":" + port))
}

func routes(app *fiber.App) {

	api := app.Group("/v1")

	health.Router(api)
	hello_world.Router(api)
	hospital.Router(app)

}
