package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("AUTH_SECRET_TOKEN"))

func AuthenticationMiddleware(allowedRoles ...string) fiber.Handler {
	return func(c fiber.Ctx) error {
		res := &response.Response{}

		authHeader := c.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			res.ErrMessage(c.Path(), fiber.StatusUnauthorized, "Missing or invalid Authorization header")
			return c.Status(res.Status).JSON(res)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			res.ErrMessage(c.Path(), fiber.StatusUnauthorized, "Invalid or expired token")
			return c.Status(res.Status).JSON(res)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			res.ErrMessage(c.Path(), fiber.StatusUnauthorized, "Invalid token claims")
			return c.Status(res.Status).JSON(res)
		}

		if role, exists := claims["role"]; exists {
			c.Locals("userRole", role)
		}

		return c.Next()

	}
}

func AuthorizationMiddleware(allowedRoles ...int64) fiber.Handler {
	return func(c fiber.Ctx) error {
		res := &response.Response{}

		localRole := c.Locals("userRole")
		if localRole == nil {
			res.ErrMessage(c.Path(), fiber.StatusForbidden, "Access Denied: No role found.")
			return c.Status(res.Status).JSON(res)
		}

		userRole, _ := localRole.(float64)

		for _, role := range allowedRoles {
			if int64(userRole) == role {
				return c.Next()
			}
		}

		res.ErrMessage(c.Path(), fiber.StatusForbidden, "Access Denied: Insufficient permissions.")
		return c.Status(res.Status).JSON(res)
	}
}
