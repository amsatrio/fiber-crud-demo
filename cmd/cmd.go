package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
	"github.com/amsatrio/fiber-crud-demo/app/middleware"
	"github.com/amsatrio/fiber-crud-demo/app/modules/health"
	hello_world "github.com/amsatrio/fiber-crud-demo/app/modules/hello_world"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital"
	_ "github.com/amsatrio/fiber-crud-demo/docs"
	"github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cache"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func Config() fiber.Config {
	// engine := html.NewFileSystem(http.Dir("./html"), ".html")
	return fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		BodyLimit:     4 * 1024 * 1024,
		// DisableStartupMessage: true,
		ServerHeader: "Fiber Vue CRUD Demo",
		AppName:      "Fiber Vue CRUD Demo v0.0.1",
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
		// Views: engine,
	}
}

func Middleware(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // v3 expects a slice of strings for origins
		AllowHeaders: []string{"*"},
	}))
	app.Use(cache.New())
	app.Use(recover.New())
	app.Use(middleware.LoggerMiddleware)

	app.Use(limiter.New(limiter.Config{
		Max:               10,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
		LimitReached: func(c fiber.Ctx) error {
			retryAfter := c.GetRespHeader("Retry-After")
			res := &response.Response{}
			res.ErrMessage(c.Path(), fiber.StatusTooManyRequests, "Too many request. retry after: "+retryAfter)
			return c.Status(fiber.StatusTooManyRequests).JSON(res)
		},
	}))
}

func Routes(app *fiber.App) {
	app.Get("/swagger/*", swaggo.HandlerDefault)

	api := app.Group("/v1")
	hello_world.Router(api)
	health.Router(api)
	hospital.Router(app)

	api.Get("/config", func(ctx fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"SERVER_PORT": os.Getenv("SERVER_PORT"),
		})
	})

	// Static assets
	app.Get("/*", static.New("./public"))

	// SPA fallback  -  catches everything else
	app.Get("*", static.New("./public/index.html"))
}

func LoadEnv() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Set to System Environment so os.Getenv works
		os.Setenv(key, value)
	}

	return scanner.Err()
}
