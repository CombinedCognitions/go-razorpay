package routes

import (
	"go-razorpay/routes/routehandlers"

	"github.com/gofiber/fiber/v2"
)

func Install(app *fiber.App) {

	app.Post("/getorderid", routehandlers.CreateOrder)
	app.Post("/verifypayment", routehandlers.VerifyPayment)
	app.Get("/getalldonors", routehandlers.GetallDonators)

}
