package rpc_server

import (
	"context"
	"errors"
	"log"

	m "github.com/jordanx8/card_db/mongodb"
	card_db "github.com/jordanx8/card_db/protos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *CardServiceServer) AddSeason(c context.Context, r *card_db.SeasonRequest) (*card_db.Response, error) {

	client, err := m.GetMongoClient()
	if err != nil {
		log.Fatal(err)
		return &card_db.Response{Response: err.Error()}, err
	}

	playerSearch := bson.D{
		{Key: "firstName", Value: r.GetFirstName()},
		{Key: "lastName", Value: r.GetLastName()},
	}

	addToArray := bson.D{{Key: "$addToSet", Value: bson.D{{Key: "seasons", Value: r.GetSeason()}}}}

	collection := client.Database("card_db").Collection("players")

	_, err = collection.UpdateOne(
		context.TODO(),
		playerSearch,
		addToArray,
	)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &card_db.Response{Response: "Player does not exist"}, errors.New("Player does not exist.")
		}
		log.Fatal(err)
		return &card_db.Response{Response: err.Error()}, err
	}

	return &card_db.Response{Response: "Success"}, err
}