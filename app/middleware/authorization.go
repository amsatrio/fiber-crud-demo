package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
	"github.com/amsatrio/fiber-crud-demo/app/initializer"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_menu"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_menu_role"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("AUTH_SECRET_TOKEN"))

func AuthenticationMiddleware() fiber.Handler {
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

		if userId, exists := claims["userId"]; exists {
			c.Locals("userId", userId)
		}

		return c.Next()

	}
}

func AuthorizationMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		res := &response.Response{}

		localRole := c.Locals("userRole")
		if localRole == nil {
			res.ErrMessage(c.Path(), fiber.StatusForbidden, "Access Denied: No role found.")
			return c.Status(res.Status).JSON(res)
		}

		userRole, ok := localRole.(float64)
		if !ok {
			res.ErrMessage(c.Path(), fiber.StatusForbidden, "Access Denied: Invalid role format.")
			return c.Status(res.Status).JSON(res)
		}

		roleID := uint(userRole)
		urlPath := c.OriginalURL()

		var menu m_menu.MMenu
		result := initializer.DB_HOSPITAL.Where("url = ? AND is_delete = 0", urlPath).First(&menu)
		if result.Error != nil {
			return c.Next()
		}

		var count int64
		initializer.DB_HOSPITAL.Model(&m_menu_role.MMenuRole{}).
			Where("menu_id = ? AND role_id = ? AND is_delete = 0", menu.Id, roleID).
			Count(&count)

		if count > 0 {
			return c.Next()
		}

		res.ErrMessage(c.Path(), fiber.StatusForbidden, "Access Denied: Insufficient permissions.")
		return c.Status(res.Status).JSON(res)
	}
}
