package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"shrading/routes"
	"shrading/shard"
)

func main() {
	app := fiber.New()

	routes.RegisterAPI(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	fmt.Println(shard.Cluster)

	errA := app.Listen(":3000")
	if errA != nil {
		log.Println("SERVICE START ERROR: " + errA.Error())
	} else {
		log.Println("SERVICE RUNNING ON PORT 3000")
	}

}
