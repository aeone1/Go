package main

import (
	"context"
	"log"
	"net"
	"os"
	"io/ioutil"

	pb "github.com/aeone1/Go/tree/master/go-usermessage-grpc/usermsg"
	"google.golang.org/grpc"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
)

/*
TODO: Error handling func and error messages (grpcError)
*/

const (
	port = ":50515"
	fileName = "user_message.json"
)

func NewUserMessageServer() *UserMessageServer {
	return &UserMessageServer{
	}
}
type UserMessageServer struct {
	pb.UnimplementedUserMessageServer
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

func WriteToFile (m *pb.MessageList, fileName string) (string, error) {
	jsonBytes, err := protojson.Marshal(m)
	if err != nil {
		return "Json Marshaling failed: %v", err
	}
	if err := ioutil.WriteFile(fileName, jsonBytes, 0664); err != nil {
		return "Failed write to file: %v", err
	}
	return "", nil
}

func (s *UserMessageServer) CreateNewMessage(
	ctx context.Context, 
	in *pb.NewMessage,
	) (*pb.Message, error) {
	log.Printf("Recieved: %v", in)
	// check file exist by reading it
	readBytes, err := ioutil.ReadFile(fileName)
	message_list := &pb.MessageList{}
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
	if err != nil {
		if os.IsNotExist(err) {
			log.Print("File not found. Creating a new file")
			message_list.Messages = append(message_list.Messages, created_message)
			if errMsg, err := WriteToFile(message_list, fileName); err != nil {
				log.Fatalf(errMsg, err)
			}
			return created_message, nil
		}
		log.Fatalf("Error reading file: %v", err)
	}

	if err := protojson.Unmarshal(readBytes, message_list); err != nil {
		log.Fatalf("Failed to pass message list: %v", err)
	}
	message_list.Messages = append(message_list.Messages, created_message)
	if errMsg, err := WriteToFile(message_list, fileName); err != nil {
		log.Fatalf(errMsg, err)
	}
	return created_message, nil
}

func (s *UserMessageServer) GetMessages(
	ctx context.Context, 
	in *pb.GetMessageParams,
	) (*pb.MessageList, error) {
	jsonBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Failed to read from file: %v", err)
	}
	message_list := &pb.MessageList{}
	if err := protojson.Unmarshal(jsonBytes, message_list); err != nil {
		log.Fatalf("Failed to Unmarshal: %v", err)
	}
	return message_list, nil
}

func main() {
	user_msg_server := NewUserMessageServer()
	if err := user_msg_server.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
