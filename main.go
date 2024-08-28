package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"kementan.com/database"
)

func main() {
	var total int64
	app := fiber.New()
	database.SetupDatabase()
	database.DB.Db.Table("submissions_simluhtan").Count(&total)
	hasil := strconv.FormatInt(total, 10)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"hasil": hasil,
		})
	})

	app.Listen(":3000")
}
