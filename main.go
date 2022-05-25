package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"shrading/routes"
)

func main() {
	app := fiber.New()

	routes.RegisterAPI(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	errA := app.Listen(":3000")
	if errA != nil {
		log.Println("SERVICE START ERROR: " + errA.Error())
	} else {
		log.Println("SERVICE RUNNING ON PORT 3000")
	}

}
