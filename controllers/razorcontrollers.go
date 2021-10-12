package controllers

import (
	"fmt"
	"log"
	"os"

	"github.com/razorpay/razorpay-go"
)

var razorclient = Razor()

func CreateOrder(amount int64, currency string, receipt string, payment_capture int16) map[string]interface{} {

	data := map[string]interface{}{
		"amount":          amount,
		"currency":        "INR",
		"receipt":         "some_receipt_id",
		"payment_capture": 1,
	}
	body, err := razorclient.Order.Create(data, nil)
	if err != nil {
		fmt.Println(err, "fucked ")
	}

	return body

}

func Getpayment(payment_id string) {
	fmt.Println("it fucking workd")
	body, err := razorclient.Payment.Fetch(payment_id, nil, nil)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(body)
}

// // fmt.Println("////////////////////////")
// fmt.Println(client.Invoice)
// fmt.Println("////////////////////////")
// fmt.Println(body["id"])
// fmt.Println("////////////////////////")

// if err != nil {
// 	log.Fatal(err)
// }

func GetURL() (string, string) {
	key := os.Getenv("KEY")
	if key == "" {
		log.Fatal("cant get .env key")
	}

	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Fatal("cant get .env Secret")
	}

	return key, secret
}

func Razor() *razorpay.Client {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Panicln(err)
	// }

	key := "rzp_test_wZsCHShozJ8ncL"
	secret := "CAAzHHQuc2WQn8KgEJygBRpP"

	client := razorpay.NewClient(key, secret)

	return client

}
