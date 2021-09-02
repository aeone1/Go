package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/aeone1/Go/tree/master/go-usermessage-grpc/usermsg"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

/*
TODO: Error handling func and error messages (grpcError)
*/

func NewUserMessageServer() *UserMessageServer {
	return &UserMessageServer{}
}

type UserMessageServer struct {
	pb.UnimplementedUserMessageServer
	conn *pgx.Conn
}

func (server *UserMessageServer) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GO_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserMessageServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	return s.Serve(lis)
}

func (server *UserMessageServer) CreateNewMessage(
	ctx context.Context,
	in *pb.NewMessage,
) (*pb.Message, error) {
	log.Printf("\nRecieved: %v\n", in)

	createSql := `
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	CREATE TABLE IF NOT EXISTS Messages(
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		user_id NUMERIC NOT NULL,
		ts TIMESTAMP,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP 
	);
	CREATE TABLE IF NOT EXISTS Message_data(
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		data TEXT,
		content jsonb,
		message_id UUID NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP 
	);`

	_, err := server.conn.Exec(context.Background(), createSql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Table creation failed: %v", err)
		os.Exit(1)
	}

	created_message_data := &pb.MessageData{
		Data:    in.GetMessageData().Data,
		Content: in.GetMessageData().Content,
	}
	created_message := &pb.Message{
		UserId:      in.GetUserId(),
		Ts:          in.GetTs(),
		MessageData: created_message_data,
	}

	//Create a transaction
	tx, err := server.conn.Begin(context.Background())
	if err != nil {
		log.Fatalf("conn.Begin failed: %v", err)
	}

	insertMessageSql := `insert into Messages(user_id, ts)
	values($1, $2) returning id`

	err = tx.QueryRow(
		context.Background(),
		insertMessageSql,
		created_message.UserId,
		created_message.Ts.AsTime(),
	).Scan(&created_message.Id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow Insert Message failed: %v\n", err)
		os.Exit(1)
	}

	insertMessageDataSql := `insert into Message_data(data, content, message_id)
	values($1, $2, $3) returning id`

	err = tx.QueryRow(
		context.Background(),
		insertMessageDataSql,
		created_message_data.Data,
		created_message.MessageData.Content,
		created_message.Id,
	).Scan(&created_message_data.Id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow Insert Message_data failed: %v\n", err)
		os.Exit(1)
	}

	tx.Commit(context.Background())
	return created_message, nil
}

func (server *UserMessageServer) GetMessages(
	ctx context.Context,
	in *pb.GetMessageParams,
) (*pb.MessageList, error) {
	message_list := &pb.MessageList{}
	selectAllMessageSql := ""
	var rows pgx.Rows
	var err error
	if in.StartTs != nil && in.EndTs != nil {
		selectAllMessageSql = `select m.id, m.user_id, m.ts, md.id as data_id, md.data, md.content 
		from Messages m 
		left join Message_data md on md.message_id = m.id
		where m.ts between $1 and $2`
		rows, err = server.conn.Query(context.Background(), selectAllMessageSql, in.StartTs.AsTime(), in.EndTs.AsTime())
		log.Printf("\nin.StartTs: %v, in.EndTs: %v", in.StartTs, in.EndTs)
	} else {
		selectAllMessageSql := `select m.id, m.user_id, m.ts, md.id as data_id, md.data, md.content 
		from Messages m 
		left join Message_data md on md.message_id = m.id`
		rows, err = server.conn.Query(context.Background(), selectAllMessageSql)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		message := &pb.Message{}
		message_data := &pb.MessageData{}
		var ts time.Time
		err = rows.Scan(
			&message.Id,
			&message.UserId,
			&ts,
			&message_data.Id,
			&message_data.Data,
			&message_data.Content,
		)
		if err != nil {
			return nil, err
		}
		message.MessageData = message_data
		message.Ts = timestamppb.New(ts)
		message_list.Messages = append(message_list.Messages, message)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return message_list, nil
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	database_url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	log.Printf("database_url: %v", database_url)
	conn, err := pgx.Connect(context.Background(), database_url)
	if err != nil {
		log.Fatalf("Unable to establish connection: %v", err)
	}
	defer conn.Close(context.Background())
	user_msg_server := NewUserMessageServer()
	user_msg_server.conn = conn
	if err := user_msg_server.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
