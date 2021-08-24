package main

import (
	"bytes"
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/aeone1/Go/tree/master/go-usermessage-grpc/usermsg"
	"google.golang.org/grpc"
)

const (
	port = "5051"
)

type UserMessageServer struct {
	pb.UnimplementedUserMessageServer
}

func (s *UserMessageServer) CreateMessage(ctx context.Context, in *pb.NewMessage) (*pb.Message, error) {
	log.Printf("Recieved: %v", in.GetUserId())
	message_id := []byte("Ju53nsdiu53nJH")
	return &pb.Message{
		Id: message_id,
		UserId: in.GetUserId(),
		Timestamp: ,
	}, nil
}

func main(){
	
}
