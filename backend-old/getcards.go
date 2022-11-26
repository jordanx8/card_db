package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"

	pb "github.com/jordanx8/card_db/backend/protos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c Card) ConvertToGRPC() *pb.Card {
	return &pb.Card{
		FirstName:     c.FirstName,
		LastName:      c.LastName,
		SeasonsPlayed: c.SeasonsPlayed,
		Season:        c.Season,
		Manufacturer:  c.Manufacturer,
		Set:           c.Set,
		Insert:        c.Insert,
		Parallel:      c.Parallel,
		CardNumber:    c.CardNumber,
		Notes:         c.Notes,
	}
}

func (s CardService) GetCards(ctx context.Context, empty *pb.Empty) (*pb.CardArray, error) {
	fmt.Println("GetCards()")
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
	findOptions := options.Find()
	cursor, err := db.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		panic(err)
	}

	var results []Card
	for cursor.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem Card
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)

	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].LastName != results[j].LastName {
			return results[i].LastName < results[j].LastName
		}
		if results[i].FirstName != results[j].FirstName {
			return results[i].FirstName < results[j].FirstName
		}
		if results[i].Season != results[j].Season {
			return results[i].Season < results[j].Season
		}
		if results[i].Set != results[j].Set {
			return results[i].Set < results[j].Set
		}
		if results[i].Insert != results[j].Insert {
			return results[i].Insert < results[j].Insert
		}
		if results[i].Parallel == "Base" {
			return true
		}
		return results[i].CardNumber < results[j].CardNumber
	})

	var cards []*pb.Card
	for _, v := range results {
		cards = append(cards, v.ConvertToGRPC())
	}

	//Close the cursor once finished
	cursor.Close(context.TODO())

	return &pb.CardArray{Cards: cards}, err
}
