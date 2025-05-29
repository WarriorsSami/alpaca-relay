package main

import (
	"alpaca-relay/generated/proto"
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"google.golang.org/grpc"
)

const (
	exchangeName      = "images"
	queueName         = "cropper"
	connectionAddress = "localhost:50051"
)

func main() {
	conn, err := grpc.NewClient(connectionAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewBrokerClient(conn)

	// Declare an exchange and a queue
	_, err = client.DeclareExchange(context.Background(), &proto.DeclareExchangeRequest{
		Exchange: exchangeName,
		Type:     proto.ExchangeType_TOPIC,
	})

	if err != nil {
		log.Fatalf("could not declare exchange: %v", err)
	}

	_, err = client.DeclareQueue(context.Background(), &proto.DeclareQueueRequest{
		Queue:    queueName,
		Exchange: exchangeName,
		Type:     proto.QueueType_NORMAL,
	})

	stream, err := client.Subscribe(context.Background(), &proto.SubscribeRequest{
		Exchange: exchangeName,
		Queue:    queueName,
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

		// Ack the message
		log.Printf("Acknowledging message: %s", msg.Payload)
		_, err = client.AckMessage(context.Background(), &proto.AckMessageRequest{
			MessageId: msg.Id,
			Queue:     queueName,
			Exchange:  exchangeName,
		})
		if err != nil {
			log.Printf("Error acknowledging message: %v", err)
		}
	}
}
