package like

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"
	"github.com/amsatrio/fiber-crud-demo/app/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewLikeRepository(initializer.DB_TWITTER_CLONE)
	service := NewLikeService(repo)
	handler := NewLikeHandler(service, validate)

	api.Post("/like", middleware.AuthenticationMiddleware(), handler.LikeCreate)
	api.Get("/like/:userId/:tweetId", middleware.AuthenticationMiddleware(), handler.LikeIndex)
	api.Get("/like", middleware.AuthenticationMiddleware(), handler.LikePage)
	api.Delete("/like/:userId/:tweetId", middleware.AuthenticationMiddleware(), handler.LikeDelete)
}
