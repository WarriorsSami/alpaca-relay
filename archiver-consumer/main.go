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

	stream, err := client.Subscribe(context.Background(), &proto.SubscribeRequest{Queue: "queue1"})
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
