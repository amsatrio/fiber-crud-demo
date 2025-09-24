package middleware

import (
	"fiber-crud-demo/util"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	err := c.Next()

	duration := time.Since(start)
	util.Log("INFO", "middleware", "LoggerMiddleware", "destination: "+c.Path()+", elapsed: "+duration.String())
	return err
}
