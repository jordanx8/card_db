package rpc_server

import (
	"context"
	"fmt"
	"log"

	m "github.com/jordanx8/card_db/mongodb"
	pb "github.com/jordanx8/card_db/protos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Player struct {
	firstName     string
	lastName      string
	seasonsPlayed string
	seasons       []string
}

func (s *CardServiceServer) GetPlayers(empty *pb.Empty, stream pb.CardService_GetPlayersServer) error {

	client, err := m.GetMongoClient()
	if err != nil {
		fmt.Println(err)
	}
	collection := m.GetDatabase(client).Collection("players")

	opts := options.Find().SetSort(bson.D{{Key: "lastName", Value: 1}})

	cur, err := collection.Find(context.TODO(), bson.D{{}}, opts)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem *pb.Player
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		if err := stream.Send(elem); err != nil {
			return err
		}
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	//Close the cursor once finished
	cur.Close(context.TODO())

	return err
}
