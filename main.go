package main

import (
	"github.com/gofiber/fiber/v2"
	"kementan.com/config"
	"kementan.com/database"
)

func main() {
	// runtime.GOMAXPROCS(5) //Limiting MAX CPU usage
	app := fiber.New()
	database.SetupDatabase()
	config.InitRoutes(app)
	app.Listen(":3000")
}
