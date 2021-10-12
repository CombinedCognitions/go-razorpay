package routehandlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"go-razorpay/controllers"
	"go-razorpay/models"

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
	fmt.Println("from frontend", data.Order_id)
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

		fmt.Println("it fucking workd")
		// url := "https://jsonplaceholder.typicode.com/todos/1"
		// res, err := http.Get(url)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		// data, _ := ioutil.ReadAll(res.Body)
		// res.Body.Close()
		// var parser models.User
		// c.BodyParser(res.Body)

		// fmt.Printf("%s, body:%s", data, parse)
		return c.JSON(fiber.Map{"status": "payment-sucess"})

	}

	return c.JSON(fiber.Map{"status": "failed"})

}
