package routehandlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"go-razorpay/controllers"
	"go-razorpay/db/dbcontrollers"
	"go-razorpay/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateOrder(c *fiber.Ctx) error {
	var datax models.Dataforid

	err := c.BodyParser(&datax)
	if err != nil {
		fmt.Println(err)

	}

	datax.Amount = datax.Amount * 100

	body := controllers.CreateOrder(datax.Amount, "INR", "some_res", 1)
	fmt.Println("order_id from server", body["id"])

	return c.JSON(fiber.Map{"orderID": body["id"]})
}

func VerifyPayment(c *fiber.Ctx) error {

	var data models.VerficationData

	err := c.BodyParser(&data)
	fmt.Println("from frontend", data.Name, data.Amount)
	if err != nil {
		fmt.Println(err)

	}
	_, secret := controllers.GetURL()
	encode := fmt.Sprintf("%s|%s", data.Order_id, data.Razorpay_payment_id)
	fmt.Println("input", encode)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(encode))

	sha := hex.EncodeToString(h.Sum(nil))

	fmt.Println("Result: " + sha)
	fmt.Println("signature :", data.Razorpay_signature, "hash:", sha)

	if data.Razorpay_signature == sha {

		controllers.Getpayment(data.Razorpay_payment_id)

		var donator models.Donator
		donator.Name = data.Name
		donator.Amount = data.Amount
		donator.Time = time.Now()

		err = dbcontrollers.Save(&donator)
		if err != nil {
			panic(err)

		}

		return c.JSON(fiber.Map{"status": "payment-sucess"})

	}

	return c.JSON(fiber.Map{"status": "failed"})

}

func GetallDonators(c *fiber.Ctx) error {

	fmt.Println("geting all donations")
	var donors []models.Donator
	donors = dbcontrollers.GetAllDonators()
	fmt.Println(donors)

	return c.JSON(fiber.Map{"donorlist": donors})

}
