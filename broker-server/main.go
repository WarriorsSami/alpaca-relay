package main

import (
	"alpaca-relay/generated/proto"
	"context"
	"errors"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

// TODO: extract queues and exchanges scaffolding logic to dedicated gRPC methods
// TODO: persist queues and exchanges in a database and also events in a database or in memory (based on queue type)
// TODO: implement acknowledgments for messages
// TODO: implement dead-letter queues and message retries

type BrokerServer struct {
	proto.UnimplementedBrokerServer
	mu        sync.RWMutex
	exchanges map[string]*Exchange
}

type Exchange struct {
	queues map[string]*Queue
}

type Queue struct {
	subscribers []chan string
}

func NewBrokerServer() *BrokerServer {
	queues := make(map[string]*Queue)
	queues["queue1"] = &Queue{
		subscribers: make([]chan string, 0),
	}

	exchanges := make(map[string]*Exchange)
	exchanges["default"] = &Exchange{
		queues: queues,
	}

	return &BrokerServer{
		exchanges: exchanges,
	}
}

func (b *BrokerServer) Publish(ctx context.Context, req *proto.PublishRequest) (*proto.PublishResponse, error) {
	b.mu.RLock()
	exchange, exists := b.exchanges[req.Exchange]
	b.mu.RUnlock()
	if !exists {
		return &proto.PublishResponse{Success: false}, errors.New("exchange not found")
	}

	queue, exists := exchange.queues[req.RoutingKey]
	if !exists {
		return &proto.PublishResponse{Success: false}, errors.New("queue not found")
	}

	for _, subscriber := range queue.subscribers {
		select {
		case subscriber <- req.Payload:
		default:
			// Handle full channel or slow consumer
		}
	}

	return &proto.PublishResponse{Success: true}, nil
}

func (b *BrokerServer) Subscribe(req *proto.SubscribeRequest, stream proto.Broker_SubscribeServer) error {
	b.mu.RLock()
	exchange, exists := b.exchanges["default"]
	b.mu.RUnlock()
	if !exists {
		return errors.New("exchange not found")
	}

	queue, exists := exchange.queues[req.Queue]
	if !exists {
		return errors.New("queue not found")
	}

	msgChan := make(chan string, 100)
	queue.subscribers = append(queue.subscribers, msgChan)

	for msg := range msgChan {
		if err := stream.Send(&proto.Message{Payload: msg}); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterBrokerServer(grpcServer, NewBrokerServer())

	log.Println("Broker server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
