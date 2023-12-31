package rpc_server

import (
	"context"
	"errors"
	"log"

	m "github.com/jordanx8/card_db/mongodb"
	card_db "github.com/jordanx8/card_db/protos"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *CardServiceServer) AddPlayer(c context.Context, p *card_db.PlayerRequest) (*card_db.Response, error) {
	log.Println("Running AddPlayer() w/ request:")
	log.Println(p)
	client, err := m.GetMongoClient()
	if err != nil {
		log.Fatal(err)
		return &card_db.Response{Response: err.Error()}, err
	}

	playerToAdd := bson.D{
		{Key: "firstName", Value: p.GetFirstName()},
		{Key: "lastName", Value: p.GetLastName()},
		{Key: "seasonsPlayed", Value: p.GetSeasonsPlayed()},
		{Key: "seasons", Value: []string{}},
	}

	collection := m.GetDatabase(client).Collection("Players")

	//check if player exists already
	count, err := collection.CountDocuments(
		context.TODO(),
		bson.D{{Key: "firstName", Value: p.GetFirstName()}, {Key: "lastName", Value: p.GetLastName()}},
	)
	if count > 0 {
		return &card_db.Response{Response: "Player already exists"}, errors.New("Player already exists.")
	}

	_, err = collection.InsertOne(context.TODO(), playerToAdd)
	if err != nil {
		log.Fatal(err)
		return &card_db.Response{Response: err.Error()}, err
	}
	return &card_db.Response{Response: "Success"}, err
}
