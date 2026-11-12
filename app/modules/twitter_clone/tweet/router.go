package tweet

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"
	"github.com/amsatrio/fiber-crud-demo/app/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTweetRepository(initializer.DB_TWITTER_CLONE)
	service := NewTweetService(repo)
	handler := NewTweetHandler(service, validate)

	api.Post("/tweet", middleware.AuthenticationMiddleware(), handler.TweetCreate)
	api.Put("/tweet", middleware.AuthenticationMiddleware(), handler.TweetUpdate)
	api.Get("/tweet/:id", middleware.AuthenticationMiddleware(), handler.TweetIndex)
	api.Get("/tweet", middleware.AuthenticationMiddleware(), handler.TweetPage)
	api.Delete("/tweet/:id", middleware.AuthenticationMiddleware(), handler.TweetDelete)
}
