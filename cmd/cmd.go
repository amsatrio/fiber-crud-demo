package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
	"github.com/amsatrio/fiber-crud-demo/app/initializer"
	"github.com/amsatrio/fiber-crud-demo/app/middleware"
	"github.com/amsatrio/fiber-crud-demo/app/modules/german_vocabulary/nomen"
	"github.com/amsatrio/fiber-crud-demo/app/modules/health"
	hello_world "github.com/amsatrio/fiber-crud-demo/app/modules/hello_world"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital"
	twitterclone "github.com/amsatrio/fiber-crud-demo/app/modules/twitter_clone"
	_ "github.com/amsatrio/fiber-crud-demo/docs"
	"github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/extractors"
	"github.com/gofiber/fiber/v3/middleware/cache"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/csrf"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func Initialize() {
	initializer.LoadEnvironmentVariables()
	initializer.LoggerInit()
	initializer.InitializeDatabaseGerman()
	initializer.InitializeDatabaseHospital()
	initializer.InitializeDatabaseTwitterClone()
}

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
	// 1. Security headers on every response
	app.Use(helmet.New(helmet.Config{
		HSTSMaxAge:         63072000,
		HSTSPreloadEnabled: true,
		ContentSecurityPolicy: "default-src 'self'; " +
			"script-src 'self' 'unsafe-inline'; " +
			"style-src 'self' 'unsafe-inline' https://fonts.googleapis.com; " +
			"font-src 'self' https://fonts.gstatic.com; " +
			"img-src 'self' data:; " +
			"connect-src 'self'",
	}))

	// 2. CORS for cross-origin access
	/*
		AllowOrigins lists the exact domains that need access. No wildcards when using credentials.
		AllowHeaders must include X-Csrf-Token if you are using CSRF with header extraction.
		AllowCredentials is true because your frontend sends cookies.
		MaxAge is 3600 (1 hour) so browsers cache preflight results instead of making an OPTIONS request before every POST.
	*/
	app.Use(cors.New(cors.Config{
		AllowOrigins:     parseOrigins(os.Getenv("CORS_ALLOW_ORIGINS")),
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Csrf-Token"},
		AllowCredentials: false,
		MaxAge:           3600,
	}))

	// 3. CSRF protection for state-changing requests
	/*
		__Host- prefix: Browsers enforce that __Host- cookies are secure, have no domain, and have a path of /. This prevents cookie injection attacks.
		CookieSecure: true: The cookie is only sent over HTTPS.
		CookieHTTPOnly: false: Your JavaScript needs to read the cookie to send the token in a header. This is the necessary trade-off for SPAs.
		CookieSameSite: "Lax": The cookie is sent with top-level navigations but not with cross-site requests. This is the baseline CSRF protection.
		Extractor from header: The CSRF token comes from the X-Csrf-Token header, not from the cookie itself. This is the Double Submit Cookie pattern.
		TrustedOrigins: needed when the SPA origin differs from the API host (e.g. direct calls without a same-origin proxy).
	*/
	app.Use(csrf.New(csrf.Config{
		CookieName:        "__Host-csrf_",
		CookieSecure:      true,
		CookieHTTPOnly:    false, // SPAs need JavaScript access
		CookieSameSite:    "Lax",
		CookieSessionOnly: true,
		Extractor:         extractors.FromHeader("X-Csrf-Token"),
		TrustedOrigins: []string{
			"https://localhost:3001",
			"http://localhost:3001",
		},
	}))

	app.Use(cache.New(cache.Config{
		Methods: []string{fiber.MethodGet},
	}))
	app.Use(recover.New())
	app.Use(middleware.LoggerMiddleware)

	app.Use(limiter.New(limiter.Config{
		Max:               100,
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
	nomen.Router(api)
	twitterclone.Router(app)

	api.Get("/config", func(ctx fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"SERVER_PORT": os.Getenv("SERVER_PORT"),
		})
	})

	app.Get("/*", static.New("./public"))
	app.Get("/*", func(c fiber.Ctx) error {
		return c.SendFile("./public/index.html")
	})
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

func parseOrigins(raw string) []string {
	if strings.TrimSpace(raw) == "" {
		return []string{"*"}
	}

	parts := strings.Split(raw, ",")
	origins := make([]string, 0, len(parts))
	for _, part := range parts {
		if origin := strings.TrimSpace(part); origin != "" {
			origins = append(origins, origin)
		}
	}
	if len(origins) == 0 {
		return []string{"*"}
	}
	return origins
}
