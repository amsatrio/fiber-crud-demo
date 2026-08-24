package main

import (
	"log"
	"os"

	"github.com/amsatrio/fiber-crud-demo/app/initializer"
	"github.com/amsatrio/fiber-crud-demo/app/util"
	"github.com/amsatrio/fiber-crud-demo/cmd"
	"github.com/gofiber/fiber/v3"
)

func init() {
	initializer.LoadEnvironmentVariables()
	initializer.LoggerInit()
	initializer.InitializeDatabase()
	initializer.InitializeDatabaseFileManagement()
}

func main() {
	// runtime.GOMAXPROCS(1)

	app := fiber.New(cmd.Config())

	cmd.Middleware(app)

	cmd.Routes(app)

	port := os.Getenv("SERVER_PORT")
	host := os.Getenv("SERVER_HOST")
	util.Log("INFO", "root", "main", "listen and serve on "+host+" port "+port)

	// ### Run
	log.Fatal(app.Listen(host+":"+port, fiber.ListenConfig{EnablePrefork: true}))
}
