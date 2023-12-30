package rpc_server

import (
	"context"
	"fmt"
	"log"

	m "github.com/jordanx8/card_db/mongodb"
	pb "github.com/jordanx8/card_db/protos"
	"go.mongodb.org/mongo-driver/bson"
)

type Cards struct {
	playerName   string
	season       string
	manufacturer string
	set          string
	insert       string
	parallel     string
	cardNumber   string
	notes        []string
	imageLink    string
	team         string
}

func (s *CardServiceServer) GetCards(query *pb.Query, stream pb.CardService_GetCardsServer) error {

	client, err := m.GetMongoClient()
	if err != nil {
		fmt.Println(err)
	}
	collection := m.GetDatabase(client).Collection(query.GetTableName())

	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem *pb.Card
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
