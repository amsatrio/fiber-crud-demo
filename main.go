package main

import (
	"log"
	"os"

	"github.com/amsatrio/fiber-crud-demo/app/util"
	"github.com/amsatrio/fiber-crud-demo/cmd"
	"github.com/gofiber/fiber/v3"
)

func init() {
	cmd.Initialize()
}

func main() {
	// runtime.GOMAXPROCS(1)

	app := fiber.New(cmd.Config())

	cmd.Middleware(app)

	cmd.Routes(app)

	port := os.Getenv("SERVER_PORT")
	host := os.Getenv("SERVER_HOST")
	util.Log("INFO", "root", "main", "listen and serve TLS on "+host+" port "+port)

	// ### Run
	// Prefork keeps per-process in-memory CSRF storage, so tokens from one
	// worker are rejected by another. Keep it off for local HTTPS + CSRF.
	log.Fatal(app.Listen(host+":"+port, fiber.ListenConfig{
		EnablePrefork: true,
		CertFile:      "../../certs/ssl.cert",
		CertKeyFile:   "../../certs/ssl.key",
	}))
}
