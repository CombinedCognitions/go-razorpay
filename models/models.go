package models

import "time"

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
	Name                string `json:"name" bson:"name"`
	Amount              int64  `json:"amount" bson:"amount"`
}

type User struct {
	User_id   int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Donator struct {
	Name   string    `json:"name" bson:"name"`
	Amount int64     `json:"amount" bson:"amount"`
	Time   time.Time `json:"time" bson:"time"`
}
