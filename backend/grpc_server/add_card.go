package rpc_server

import (
	"context"
	"errors"
	"log"

	m "github.com/jordanx8/card_db/mongodb"
	card_db "github.com/jordanx8/card_db/protos"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *CardServiceServer) AddCard(c context.Context, p *card_db.CardRequest) (*card_db.Response, error) {
	log.Println("Running AddCard() w/ request:")
	log.Println(p)
	client, err := m.GetMongoClient()
	if err != nil {
		log.Fatal(err)
		return &card_db.Response{Response: err.Error()}, err
	}

	cardToAdd := bson.D{
		{Key: "playerName", Value: p.GetPlayerName()},
		{Key: "season", Value: p.GetSeason()},
		{Key: "manufacturer", Value: p.GetManufacturer()},
		{Key: "set", Value: p.GetSet()},
		{Key: "insert", Value: p.GetInsert()},
		{Key: "parallel", Value: p.GetParallel()},
		{Key: "cardNumber", Value: p.GetCardNumber()},
		{Key: "notes", Value: p.GetNotes()},
		{Key: "imageLink", Value: p.GetImageLink()},
		{Key: "team", Value: p.GetTeam()},
	}

	collection := client.Database("card_db").Collection(p.GetTableName())

	//check if card exists already
	count, err := collection.CountDocuments(
		context.TODO(),
		bson.D{
			{Key: "playerName", Value: p.GetPlayerName()},
			{Key: "season", Value: p.GetSeason()},
			{Key: "manufacturer", Value: p.GetManufacturer()},
			{Key: "set", Value: p.GetSet()},
			{Key: "insert", Value: p.GetInsert()},
			{Key: "parallel", Value: p.GetParallel()},
			{Key: "cardNumber", Value: p.GetCardNumber()},
			{Key: "team", Value: p.GetTeam()},
		},
	)
	if count > 0 {
		return &card_db.Response{Response: "Card already exists"}, errors.New("Card already exists.")
	}

	_, err = collection.InsertOne(context.TODO(), cardToAdd)
	if err != nil {
		log.Fatal(err)
		return &card_db.Response{Response: err.Error()}, err
	}
	return &card_db.Response{Response: "Success"}, err
}
