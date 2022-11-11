package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/jordanx8/card_db/backend/protos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s CardService) AddCard(ctx context.Context, card *pb.Card) (*pb.Response, error) {
	fmt.Println("AddCard()")
	fmt.Println(card)
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	secret := os.Args[1]
	clientOptions := options.Client().ApplyURI("mongodb+srv://" + secret + "@carddatabase.3nsgdce.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPIOptions)
	// Connect to the MongoDB and return Client instance
	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("mongo.Connect() ERROR: %v", err)
	}
	//check the connection
	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	db := mongoClient.Database("Cards").Collection("Pelicans")
	var search bson.M
	if err = db.FindOne(ctx, bson.M{"firstname": card.FirstName, "lastname": card.LastName}).Decode(&search); err != nil {
		fmt.Println(err)
		fmt.Println("Player does not exist in database")
	} else {
		var search2 bson.M
		fmt.Println("Player exists in database; copying player info to new card.")
		var search3 bson.M
		if err = db.FindOne(ctx, bson.M{"season": card.Season}).Decode(&search3); err != nil {
			fmt.Println("Season exists as well; Copying that.")
			if err = db.FindOne(ctx, bson.M{"firstname": search3["firstname"], "lastname": search3["lastname"], "seasonsplayed": search3["seasonsplayed"], "season": search3["season"], "manufacturer": card.Manufacturer, "set": card.Set, "insert": card.Insert, "parallel": card.Parallel, "cardnumber": card.CardNumber, "notes": card.Notes}).Decode(&search2); err == nil {
				fmt.Println(search2)
				fmt.Println("This card already exists.")
				return &pb.Response{Success: 0}, nil
			}
			doc := bson.D{
				{Key: "firstname", Value: search3["firstname"]},
				{Key: "lastname", Value: search3["lastname"]},
				{Key: "seasonsplayed", Value: search3["seasonsplayed"]},
				{Key: "season", Value: search3["season"]},
				{Key: "manufacturer", Value: card.Manufacturer},
				{Key: "set", Value: card.Set},
				{Key: "insert", Value: card.Insert},
				{Key: "parallel", Value: card.Parallel},
				{Key: "cardnumber", Value: card.CardNumber},
				{Key: "notes", Value: card.Notes}}
			result, err := db.InsertOne(context.TODO(), doc)
			fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
			if err != nil {
				return &pb.Response{Success: 0}, nil
			}
			return &pb.Response{Success: 1}, err
		}
		fmt.Println(err)
		fmt.Println(search["season"])
		fmt.Println(card.Season)
		fmt.Println(search["firstname"])
		fmt.Println(card.FirstName)
		fmt.Println(search["lastname"])
		fmt.Println(card.LastName)
		doc := bson.D{
			{Key: "firstname", Value: search["firstname"]},
			{Key: "lastname", Value: search["lastname"]},
			{Key: "seasonsplayed", Value: search["seasonsplayed"]},
			{Key: "season", Value: card.Season},
			{Key: "manufacturer", Value: card.Manufacturer},
			{Key: "set", Value: card.Set},
			{Key: "insert", Value: card.Insert},
			{Key: "parallel", Value: card.Parallel},
			{Key: "cardnumber", Value: card.CardNumber},
			{Key: "notes", Value: card.Notes}}
		result, err := db.InsertOne(context.TODO(), doc)
		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
		if err != nil {
			return &pb.Response{Success: 0}, nil
		}
		return &pb.Response{Success: 1}, err
	}

	doc := bson.D{
		{Key: "firstname", Value: card.FirstName},
		{Key: "lastname", Value: card.LastName},
		{Key: "seasonsplayed", Value: card.SeasonsPlayed},
		{Key: "season", Value: card.Season},
		{Key: "manufacturer", Value: card.Manufacturer},
		{Key: "set", Value: card.Set},
		{Key: "insert", Value: card.Insert},
		{Key: "parallel", Value: card.Parallel},
		{Key: "cardnumber", Value: card.CardNumber},
		{Key: "notes", Value: card.Notes}}
	result, err := db.InsertOne(context.TODO(), doc)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

	if err != nil {
		return &pb.Response{Success: 0}, nil
	}
	return &pb.Response{Success: 1}, err
}