package rpc_server

import (
	pb "github.com/jordanx8/card_db/protos"
)

type CardServiceServer struct {
	pb.UnimplementedCardServiceServer
}
