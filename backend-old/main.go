package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/jordanx8/card_db/backend/protos"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type CardService struct {
	MongoClient *mongo.Client
	pb.UnimplementedCardServiceServer
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("forgot to add mongo pw")
	}
	// ctx := context.Background()
	// b, err := ioutil.ReadFile("credentials.json")
	// if err != nil {
	// 	log.Fatalf("Unable to read client secret file: %v", err)
	// }

	// // If modifying these scopes, delete your previously saved token.json.
	// config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	// if err != nil {
	// 	log.Fatalf("Unable to parse client secret file to config: %v", err)
	// }
	// client := getClient(config)

	// srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	// if err != nil {
	// 	log.Fatalf("Unable to retrieve Sheets client: %v", err)
	// }

	// spreadsheetId := "10-eMLMp7EGFdQd9rZW1Bb5TzB7AN00uDvg__40_M05c"
	// readRange := "Pelicans!A2:E"
	// resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	// cards := scrapeCards(resp, "Pelicans")
	// readRange = "LSU!A2:E"
	// resp, err = srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	// cards2 := scrapeCards(resp, "LSU")

	// seedToMongo(cards, "Pelicans")
	// seedToMongo(cards2, "LSU")
	// fmt.Println("Seeding done.")

	ctx := context.Background()
	log.Println("starting up")
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	clientOptions := options.Client().ApplyURI("mongodb://localhost")
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	s := grpc.NewServer()
	server := CardService{MongoClient: client}
	pb.RegisterCardServiceServer(s, &server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// results := getCards("Pelicans")
	// fmt.Print(results[0])
}
