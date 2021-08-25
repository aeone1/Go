package main

import (
	"context"
	"log"
	"net"

	pb "github.com/aeone1/Go/tree/master/go-usermessage-grpc/usermsg"
	"google.golang.org/grpc"
	"github.com/google/uuid"
)

const (
	port = ":5051"
)

func NewUserMessageServer() *UserMessageServer {
	return &UserMessageServer{
		message_list: &pb.MessageList{},
	}
}
type UserMessageServer struct {
	pb.UnimplementedUserMessageServer
	message_list *pb.MessageList
}

func (server *UserMessageServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserMessageServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	return s.Serve(lis)
}

func (s *UserMessageServer) CreateNewMessage(ctx context.Context, in *pb.NewMessage) (*pb.Message, error) {
	log.Printf("Recieved: %v, %v", in.GetUserId(), in)
	//message_id := []byte(RandStringBytesMaskImprSrcSB(15).Output)
	message_id := uuid.New().String()
	created_message_data := &pb.MessageData{
		Id: uuid.New().String(),
		Data: in.GetMessageData().Data,
		Content: in.GetMessageData().Content,
	}
	created_message := &pb.Message{
		Id: message_id,
		UserId: in.GetUserId(),
		Ts: in.GetTs(),
		MessageData: created_message_data,
	}
	s.message_list.Messages = append(s.message_list.Messages, created_message)
	return created_message, nil
}

func (s *UserMessageServer) GetMessages(ctx context.Context, in *pb.GetMessageParams) (*pb.MessageList, error) {
	return s.message_list, nil
}

func main() {
	user_msg_server := NewUserMessageServer()
	if err := user_msg_server.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
