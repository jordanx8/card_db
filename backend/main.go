package main

import (
	"log"
	"net"

	grpc_server "github.com/jordanx8/card_db/grpc_server"
	pb "github.com/jordanx8/card_db/protos"
	"google.golang.org/grpc"
)

// TODO: player scraper from basketball reference. grab all https://www.basketball-reference.com/teams/NOH/players.html Pels

func main() {

	log.Println("starting up")
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	server := grpc_server.CardServiceServer{}
	pb.RegisterCardServiceServer(s, &server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
