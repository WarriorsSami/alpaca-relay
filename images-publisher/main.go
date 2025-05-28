package main

import (
	"alpaca-relay/generated/proto"
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewBrokerClient(conn)

	// Declare an exchange
	_, err = client.DeclareExchange(context.Background(), &proto.DeclareExchangeRequest{
		Exchange: "images",
		Type:     proto.ExchangeType_TOPIC,
	})

	for i := 0; i < 10; i++ {
		msg := &proto.PublishRequest{
			Exchange:   "images",
			RoutingKey: "images.*",
			Payload:    "Message " + strconv.Itoa(i),
		}
		_, err := client.Publish(context.Background(), msg)
		if err != nil {
			log.Printf("could not publish message: %v", err)
		}
		time.Sleep(1 * time.Second)
	}
}
