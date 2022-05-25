package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"shrading/register_subject"
)

func RegisterAPI(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, code-otp, request-id, authorization, app-secret-key",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))
	registerTKB(app)
}

func registerTKB(app *fiber.App) {
	app.Post("/register_subject", register_subject.RegisterSubject)
}
