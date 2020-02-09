package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	collection := client.Database("testing").Collection("numbers")

	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	id := res.InsertedID

	fmt.Println(id)
}
