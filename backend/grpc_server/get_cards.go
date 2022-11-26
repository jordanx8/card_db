package rpc_server

import (
	"fmt"

	m "github.com/jordanx8/card_db/mongodb"
	pb "github.com/jordanx8/card_db/protos"
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
}

func (s *cardServiceServer) GetCards(query *pb.Query, stream pb.CardService_GetCardsServer) error {

	client, err := m.GetMongoClient()
	if err != nil {
		fmt.Println(err)
	}
	collection := client.Database("card_db").Collection(query.GetTableName())
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
