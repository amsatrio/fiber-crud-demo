package main

import (
	"errors"
	"log"
	"os"

	_ "fiber-crud-demo/docs"
	"fiber-crud-demo/dto/response"
	"fiber-crud-demo/initializer"
	"fiber-crud-demo/middleware"
	"fiber-crud-demo/modules/health"
	"fiber-crud-demo/modules/hello_world"
	"fiber-crud-demo/modules/m_biodata"
	"fiber-crud-demo/modules/m_role"
	"fiber-crud-demo/modules/m_user"
	"fiber-crud-demo/util"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func init() {
	initializer.LoadEnvironmentVariables()
	initializer.LoggerInit()
	initializer.ConnectToDB()
}

func config() fiber.Config {
	return fiber.Config{
		Prefork:               true,
		CaseSensitive:         true,
		StrictRouting:         true,
		BodyLimit:             4 * 1024 * 1024,
		DisableStartupMessage: true,
		ServerHeader:          "Fiber Audio Management",
		AppName:               "Fiber Audio Management v0.0.1",
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		ErrorHandler: func(c *fiber.Ctx, err error) error {

			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			res := &response.Response{}
			res.ErrMessage(c.Path(), code, err.Error())

			return c.Status(code).JSON(res)
		},
	}
}

func main() {
	app := fiber.New(config())

	// ### Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	app.Use(cache.New())
	app.Use(recover.New())
	app.Use(middleware.LoggerMiddleware)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// ### Routes
	routes(app)

	port := os.Getenv("SERVER_PORT")
	host := os.Getenv("SERVER_HOST")
	util.Log("INFO", "root", "main", "listen and serve on "+host+" port "+port)

	// ### Run
	log.Fatal(app.Listen(host + ":" + port))
}

func routes(app *fiber.App) {
	// HEALTH
	health_api := app.Group("/health")
	health_api.Get("/status", health.Status)

	// HELLO WORLD
	hello_world_api := app.Group("/hello-world")
	hello_world_api.Get("", hello_world.HelloWorld)
	hello_world_api.Get("/path/:message", hello_world.HelloWorldPath)
	hello_world_api.Get("/query", hello_world.HelloWorldQuery)
	hello_world_api.Post("/payload", hello_world.HelloWorldPayload)
	hello_world_api.Get("/error/:type", hello_world.HelloWorldError)

	// MASTER ROLE
	m_role_api := app.Group("/m-role")
	m_role_api.Post("", m_role.MRoleCreate)
	m_role_api.Put("", m_role.MRoleUpdate)
	m_role_api.Get(":id", m_role.MRoleIndex)
	m_role_api.Get("", m_role.MRolePage)
	m_role_api.Delete(":id", m_role.MRoleDelete)

	// MASTER BIODATA
	m_biodata_api := app.Group("/m-biodata")
	m_biodata_api.Post("", m_biodata.MBiodataCreate)
	m_biodata_api.Put("", m_biodata.MBiodataUpdate)
	m_biodata_api.Get(":id", m_biodata.MBiodataIndex)
	m_biodata_api.Get("", m_biodata.MBiodataPage)
	m_biodata_api.Delete(":id", m_biodata.MBiodataDelete)

	// MASTER USER
	m_user_api := app.Group("/m-user")
	m_user_api.Post("", m_user.MUserCreate)
	m_user_api.Put("", m_user.MUserUpdate)
	m_user_api.Get(":id", m_user.MUserIndex)
	m_user_api.Get("", m_user.MUserPage)
	m_user_api.Delete(":id", m_user.MUserDelete)
}
