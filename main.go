package main

import (
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"kementan.com/config"
	"kementan.com/database"
)

func main() {
	// runtime.GOMAXPROCS(5) //Limiting MAX CPU usage
	file, _ := os.OpenFile("status.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	iw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(iw)
	app := fiber.New()
	database.SetupDatabase()
	config.InitRoutes(app)
	app.Listen(":3000")
}
