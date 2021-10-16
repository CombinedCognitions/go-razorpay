package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Db() *mongo.Client {
	uri := Geturl()

	fmt.Println(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Connected to DB gg ez ")

	return client
}

func Geturl() string {

	return fmt.Sprintf("mongodb+srv://akash:broimfucked343@cluster0.qgemk.mongodb.net/tes%s", "t")
}
