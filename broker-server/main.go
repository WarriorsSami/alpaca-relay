package main

import (
	"alpaca-relay/generated/proto"
	"context"
	"errors"
	"log"
	"net"
	"regexp"
	"sync"

	"google.golang.org/grpc"
)

// TODO: implement acknowledgments for messages
// TODO: implement dead-letter queues and message retries
// TODO: persist queues and exchanges in a database and also events in a database or in memory (based on queue type) (e.g. Sqlite using sqlc)
// TODO: support multiple message formats (e.g., JSON, Protobuf, CloudEvents, etc.)
// TODO: add tests for the broker server

type BrokerServer struct {
	proto.UnimplementedBrokerServer
	mu        sync.RWMutex
	exchanges map[string]*Exchange
}

func NewBrokerServer() *BrokerServer {
	return &BrokerServer{
		exchanges: make(map[string]*Exchange),
	}
}

type Exchange struct {
	exchangeType proto.ExchangeType
	queues       map[string]*Queue
}

func NewExchange(exchangeType proto.ExchangeType) *Exchange {
	return &Exchange{
		exchangeType: exchangeType,
		queues:       make(map[string]*Queue),
	}
}

func (e *Exchange) GetQueues(routingKey string) []*Queue {
	queues := make([]*Queue, 0)
	switch e.exchangeType {
	case proto.ExchangeType_DIRECT:
		if queue, exists := e.queues[routingKey]; exists {
			queues = append(queues, queue)
		}
	case proto.ExchangeType_FANOUT:
		for _, queue := range e.queues {
			queues = append(queues, queue)
		}
	case proto.ExchangeType_TOPIC:
		// Match routing key with queue names (e.g. "images.archive" matches "images.archive" and "images.*")
		for queueName, queue := range e.queues {
			// Use regex for more complex routing key matching
			regex := "^" + routingKey + "$"
			if matched, _ := regexp.MatchString(regex, queueName); matched {
				queues = append(queues, queue)
			}
		}
	}

	return queues
}

type Queue struct {
	queueType   proto.QueueType
	subscribers []chan string
}

func NewQueue(queueType proto.QueueType) *Queue {
	return &Queue{
		queueType:   queueType,
		subscribers: make([]chan string, 0),
	}
}

func (b *BrokerServer) Publish(ctx context.Context, req *proto.PublishRequest) (*proto.PublishResponse, error) {
	b.mu.RLock()
	exchange, exists := b.exchanges[req.Exchange]
	b.mu.RUnlock()
	if !exists {
		return &proto.PublishResponse{Success: false}, errors.New("exchange not found")
	}

	queues := exchange.GetQueues(req.RoutingKey)

	for _, queue := range queues {
		for _, subscriber := range queue.subscribers {
			select {
			case subscriber <- req.Payload:
			default:
				// Handle full channel or slow consumer
			}
		}
	}

	return &proto.PublishResponse{Success: true}, nil
}

func (b *BrokerServer) Subscribe(req *proto.SubscribeRequest, stream proto.Broker_SubscribeServer) error {
	b.mu.RLock()
	exchange, exists := b.exchanges[req.Exchange]
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

func (b *BrokerServer) DeclareExchange(ctx context.Context, req *proto.DeclareExchangeRequest) (*proto.DeclareExchangeResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, exists := b.exchanges[req.Exchange]; exists {
		return &proto.DeclareExchangeResponse{Success: true}, nil
	}

	b.exchanges[req.Exchange] = NewExchange(req.Type)
	return &proto.DeclareExchangeResponse{Success: true}, nil
}

func (b *BrokerServer) DeclareQueue(ctx context.Context, req *proto.DeclareQueueRequest) (*proto.DeclareQueueResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	exchange, exists := b.exchanges[req.Exchange]
	if !exists {
		return &proto.DeclareQueueResponse{Success: false}, errors.New("exchange not found")
	}

	if _, exists := exchange.queues[req.Queue]; exists {
		return &proto.DeclareQueueResponse{Success: true}, nil
	}

	exchange.queues[req.Queue] = NewQueue(req.Type)
	return &proto.DeclareQueueResponse{Success: true}, nil
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
