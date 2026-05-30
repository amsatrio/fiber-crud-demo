package middleware

import (
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/util"

	"github.com/gofiber/fiber/v3"
)

func LoggerMiddleware(c fiber.Ctx) error {
	start := time.Now()

	err := c.Next()

	duration := time.Since(start)
	util.Log("INFO", "middleware", "LoggerMiddleware", "destination: "+c.Path()+", elapsed: "+duration.String())
	return err
}
