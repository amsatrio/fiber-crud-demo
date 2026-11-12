package twitterclone

import (
	"github.com/amsatrio/fiber-crud-demo/app/modules/twitter_clone/auth"
	"github.com/amsatrio/fiber-crud-demo/app/modules/twitter_clone/follow"
	"github.com/amsatrio/fiber-crud-demo/app/modules/twitter_clone/like"
	"github.com/amsatrio/fiber-crud-demo/app/modules/twitter_clone/tweet"
	"github.com/amsatrio/fiber-crud-demo/app/modules/twitter_clone/user"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func Router(app *fiber.App) {
	var validate = validator.New()

	api := app.Group("/v1/twitter-clone")

	auth.GetRouter(api, validate)
	user.GetRouter(api, validate)
	tweet.GetRouter(api, validate)
	follow.GetRouter(api, validate)
	like.GetRouter(api, validate)
}
