package follow

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"
	"github.com/amsatrio/fiber-crud-demo/app/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewFollowRepository(initializer.DB_TWITTER_CLONE)
	service := NewFollowService(repo)
	handler := NewFollowHandler(service, validate)

	api.Post("/follow", middleware.AuthenticationMiddleware(), handler.FollowCreate)
	api.Get("/follow/:followerId/:followedId", middleware.AuthenticationMiddleware(), handler.FollowIndex)
	api.Get("/follow", middleware.AuthenticationMiddleware(), handler.FollowPage)
	api.Delete("/follow/:followerId/:followedId", middleware.AuthenticationMiddleware(), handler.FollowDelete)
}
