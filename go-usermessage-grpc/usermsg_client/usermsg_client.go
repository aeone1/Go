package main

import (
	"context"
	"log"
	"time"

	pb "github.com/aeone1/Go/tree/master/go-usermessage-grpc/usermsg"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	address = "localhost:50515"
)

type Message struct {
	user_id uint64
	ts *timestamp.Timestamp
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock()) //with block is a blocking call to dail
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserMessageClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	message := Message{ 1314674, timestamppb.New(time.Now())}
	message_data_content := make(map[string]string)
	message_data_content["H"] = "Hello"
	message_data_content["W"] = "WO~RLD"
	message_data := &pb.NewMessageData{
		Data: "hello world",
		Content: message_data_content,
	}
	//third_message := pb.Message{ UserId: 243534, }

	r, err := c.CreateNewMessage(ctx, &pb.NewMessage{
		UserId: message.user_id, 
		Ts: message.ts, 
		MessageData: message_data,
	})
	if err != nil {
		log.Fatalf("could not create usermessage: %v", err)
	}
	log.Printf(`UserMessage:
	Id: %v
	UserId: %v
	Ts: %v
	MessageData: %v`, r.GetId(), r.GetUserId(), r.GetTs(), r.GetMessageData())
	params := &pb.GetMessageParams{}
	rec, err := c.GetMessages(ctx, params)
	if err != nil {
		log.Fatalf("could not retrieve user messages: %v", err)
	}
	log.Print("\nUSER MESSAGES LIST: \n")
	log.Printf("rec.GetMessages: %v\n", rec.GetMessages())
}