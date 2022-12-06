package rpc_server

import (
	"context"
	"errors"
	"log"

	m "github.com/jordanx8/card_db/mongodb"
	card_db "github.com/jordanx8/card_db/protos"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *CardServiceServer) AddPlayer(c context.Context, p *card_db.Player) (*card_db.RequestResponse, error) {

	player := Player{
		firstName:     p.GetFirstName(),
		lastName:      p.GetLastName(),
		seasonsPlayed: p.GetSeasonsPlayed(),
		seasons:       p.GetSeasons(),
	}

	client, err := m.GetMongoClient()
	if err != nil {
		log.Fatal(err)
		return &card_db.RequestResponse{Success: false}, err
	}

	doc := bson.D{
		{Key: "firstName", Value: player.firstName},
		{Key: "lastName", Value: player.lastName},
		{Key: "seasonsPlayed", Value: player.seasonsPlayed},
		{Key: "seasons", Value: player.seasons},
	}

	collection := client.Database("card_db").Collection("players")

	//check if player exists already
	count, err := collection.CountDocuments(
		context.TODO(),
		bson.D{{Key: "firstName", Value: player.firstName}, {Key: "lastName", Value: player.lastName}},
	)
	if count > 0 {
		return &card_db.RequestResponse{Success: false}, errors.New("Player already exists.")
	}

	_, err = collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
		return &card_db.RequestResponse{Success: false}, err
	}
	return &card_db.RequestResponse{Success: true}, err
}
