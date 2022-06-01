package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"shrading/connection"
	"shrading/routes"
)

func main() {
	app := fiber.New()

	routes.RegisterAPI(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	a := connection.ConnectLivechatElastic()
	fmt.Println(a)

	errA := app.Listen(":3000")
	if errA != nil {
		log.Println("SERVICE START ERROR: " + errA.Error())
	} else {
		log.Println("SERVICE RUNNING ON PORT 3000")
	}
}
