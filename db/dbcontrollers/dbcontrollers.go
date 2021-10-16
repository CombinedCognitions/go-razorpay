package dbcontrollers

import (
	"context"
	"fmt"
	"go-razorpay/db"
	"go-razorpay/models"

	"go.mongodb.org/mongo-driver/bson"
)

var dbclient = db.Db()

var donationcollection = dbclient.Database("donatorsDB").Collection("topdonators")

func Save(donator *models.Donator) error {

	_, err := donationcollection.InsertOne(context.Background(), donator)

	if err != nil {
		panic(err)
	}

	fmt.Println("Added New Donator ", donator)
	return err

}

func GetAllDonators() []models.Donator {

	cursor, err := donationcollection.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}

	var donatordoc []models.Donator

	for cursor.Next(context.Background()) {

		var single models.Donator

		err := cursor.Decode(&single)
		if err != nil {
			fmt.Println(err)
		}

		donatordoc = append(donatordoc, single)

	}

	return donatordoc

}

func Close() error {
	err := dbclient.Disconnect(context.Background())
	if err != nil {
		panic(err)

	}

	fmt.Println("db closed gg ez")
	return nil
}
