package models

type Dataforid struct {
	Name     string `json:"name" `
	Amount   int64  `json:"amount"`
	Selected string `json:"selected"`
	Currency string `json:"currency"`
}

type VerficationData struct {
	Razorpay_payment_id string `json:"razorpay_payment_id"`
	Razorpay_order_id   string `json:"razorpay_order_id"`
	Razorpay_signature  string `json:"razorpay_signature"`
	Order_id            string `json:"order_id"`
}

type User struct {
	User_id   int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
