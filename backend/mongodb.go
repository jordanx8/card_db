package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func seedToMongo(cards []Card, collectionName string) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	secret := os.Args[1]
	clientOptions := options.Client().ApplyURI("mongodb+srv://" + secret + "@carddatabase.3nsgdce.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPIOptions)
	// Connect to the MongoDB and return Client instance
	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("mongo.Connect() ERROR: %v", err)
	}
	mongoCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	col := mongoClient.Database("Cards").Collection(collectionName)
	if err = col.Drop(mongoCtx); err != nil {
		log.Fatal(err)
	}
	col = mongoClient.Database("Cards").Collection(collectionName) //check for error again
	if err != nil {
		log.Fatal(err)
	} //else mongo has been connected
	fmt.Println("Successfully connected to Mongo")
	for i := range cards {
		card := cards[i]
		fmt.Println("inserting", card)
		// Call the InsertOne() method and pass the context and doc objects
		col.InsertOne(mongoCtx, card)
	}
}
