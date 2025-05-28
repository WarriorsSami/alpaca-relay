package main

import (
	"alpaca-relay/generated/proto"
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewBrokerClient(conn)

	// Declare an exchange and a queue
	_, err = client.DeclareExchange(context.Background(), &proto.DeclareExchangeRequest{
		Exchange: "images",
		Type:     proto.ExchangeType_TOPIC,
	})

	if err != nil {
		log.Fatalf("could not declare exchange: %v", err)
	}

	_, err = client.DeclareQueue(context.Background(), &proto.DeclareQueueRequest{
		Queue:    "images.archive",
		Exchange: "images",
		Type:     proto.QueueType_NORMAL,
	})

	stream, err := client.Subscribe(context.Background(), &proto.SubscribeRequest{
		Exchange: "images",
		Queue:    "images.archive",
	})
	if err != nil {
		log.Fatalf("could not subscribe: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Fatalf("error receiving message: %v", err)
		}
		log.Printf("Received message: %s", msg.Payload)
	}
}
