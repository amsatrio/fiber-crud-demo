package main

import (
	"errors"
	"log"
	"os"
	"runtime"

	_ "fiber-crud-demo/docs"
	"fiber-crud-demo/dto/response"
	"fiber-crud-demo/initializer"
	"fiber-crud-demo/middleware"
	"fiber-crud-demo/modules/health"
	"fiber-crud-demo/modules/hello_world"
	"fiber-crud-demo/modules/m_biodata"
	"fiber-crud-demo/modules/m_file"
	"fiber-crud-demo/modules/m_module"
	"fiber-crud-demo/modules/m_role"
	"fiber-crud-demo/modules/m_user"

	"fiber-crud-demo/util"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
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
		Views:       htmlEngine,
		ViewsLayout: "layouts/main",
	}
}

func main() {
	runtime.GOMAXPROCS(1)

	app := fiber.New(config())

	// ### Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	// app.Use(cache.New())
	app.Use(recover.New())
	app.Use(middleware.LoggerMiddleware)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Static("/public", "./web/public")

	// ### Routes
	routes(app)

	port := os.Getenv("SERVER_PORT")
	host := os.Getenv("SERVER_HOST")
	util.Log("INFO", "root", "main", "listen and serve on "+host+" port "+port)

	// ### Run
	log.Fatal(app.Listen(host + ":" + port))
}

func routes(app *fiber.App) {
	mFileRepo := m_file.NewMFileRepository()

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

	// WEB MASTER BIODATA
	mBiodataWebHandler := m_biodata.NewMBiodataWebHandler()
	m_biodata_web := app.Group("/web/m-biodata")
	m_biodata_web.Get("/datatable", mBiodataWebHandler.MBiodataDatatableWebIndex)
	m_biodata_web.Get("/table", mBiodataWebHandler.MBiodataTableWebIndex)
	m_biodata_web.Get("", mBiodataWebHandler.MBiodataWebIndex)

	var validate = validator.New()

	// MASTER ROLE
	mRoleRepo := m_role.NewMRoleRepository(initializer.DB)
	mRoleService := m_role.NewMRoleService(mRoleRepo)
	mRoleHandler := m_role.NewMRoleHandler(mRoleService, validate)
	m_role_api := app.Group("/m-role")
	m_role_api.Post("", mRoleHandler.MRoleCreate)
	m_role_api.Put("", mRoleHandler.MRoleUpdate)
	m_role_api.Get(":id", mRoleHandler.MRoleIndex)
	m_role_api.Get("", mRoleHandler.MRolePage)
	m_role_api.Delete(":id", mRoleHandler.MRoleDelete)

	// MASTER BIODATA
	mBiodataRepo := m_biodata.NewMBiodataRepository(initializer.DB)
	mBiodataService := m_biodata.NewMBiodataService(mBiodataRepo, mFileRepo)
	mBiodataHandler := m_biodata.NewMBiodataHandler(mBiodataService, validate)
	m_biodata_api := app.Group("/m-biodata")
	m_biodata_api.Post("", mBiodataHandler.MBiodataCreate)
	m_biodata_api.Put("", mBiodataHandler.MBiodataUpdate)
	m_biodata_api.Get(":id", mBiodataHandler.MBiodataIndex)
	m_biodata_api.Get("", mBiodataHandler.MBiodataPage)
	m_biodata_api.Delete(":id", mBiodataHandler.MBiodataDelete)

	// MASTER USER
	mUserRepo := m_user.NewMUserRepository(initializer.DB)
	mUserService := m_user.NewMUserService(mUserRepo)
	mUserHandler := m_user.NewMUserHandler(mUserService, validate)
	m_user_api := app.Group("/m-user")
	m_user_api.Post("", mUserHandler.MUserCreate)
	m_user_api.Put("", mUserHandler.MUserUpdate)
	m_user_api.Get(":id", mUserHandler.MUserIndex)
	m_user_api.Get("", mUserHandler.MUserPage)
	m_user_api.Delete(":id", mUserHandler.MUserDelete)

	// MASTER MODULE
	mModuleRepo := m_module.NewMModuleRepository(initializer.DB)
	mModuleService := m_module.NewMModuleService(mModuleRepo)
	mModuleHandler := m_module.NewMModuleHandler(mModuleService, validate)
	m_module_api := app.Group("/m-module")
	m_module_api.Post("", mModuleHandler.MModuleCreate)
	m_module_api.Put("", mModuleHandler.MModuleUpdate)
	m_module_api.Get(":id", mModuleHandler.MModuleIndex)
	m_module_api.Get("", mModuleHandler.MModulePage)
	m_module_api.Delete(":id", mModuleHandler.MModuleDelete)
}
