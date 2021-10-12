package main

import (
	"go-razorpay/routes"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowMethods:     "POST, GET, OPTIONS, PUT, DELETE",
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
		AllowCredentials: true,
		ExposeHeaders:    "true",
	}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "helll world"})
	})

	routes.Install(app)

	log.Fatal(app.Listen(":8080"))

}
