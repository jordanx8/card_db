package rpc_server

import (
	"context"
	"fmt"
	"log"

	m "github.com/jordanx8/card_db/mongodb"
	pb "github.com/jordanx8/card_db/protos"
	"go.mongodb.org/mongo-driver/bson"
)

type Player struct {
	firstName     string
	lastName      string
	seasonsPlayed string
	seasons       []string
}

func (s *cardServiceServer) GetPlayers(empty *pb.Empty, stream pb.CardService_GetPlayersServer) error {

	client, err := m.GetMongoClient()
	if err != nil {
		fmt.Println(err)
	}
	collection := client.Database("card_db").Collection("players")

	cur, err := collection.Find(context.TODO(), bson.D{{}})
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

	return nil
}

// func (s *routeGuideServer) ListFeatures(rect *pb.Rectangle, stream pb.RouteGuide_ListFeaturesServer) error {
// 	for _, feature := range s.savedFeatures {
// 	  if inRange(feature.Location, rect) {
// 		if err := stream.Send(feature); err != nil {
// 		  return err
// 		}
// 	  }
// 	}
// 	return nil
//   }
