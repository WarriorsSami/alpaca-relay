package main

import (
	"alpaca-relay/generated/proto"
	"context"
	"errors"
	"github.com/google/uuid"
	"log"
	"net"
	"regexp"
	"sync"

	"google.golang.org/grpc"
)

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

type Message struct {
	Id      uuid.UUID
	Payload string
}

func NewMessage(payload string) *Message {
	return &Message{
		Id:      uuid.New(),
		Payload: payload,
	}
}

func (m *Message) ToProto() *proto.Message {
	return &proto.Message{
		Id:      m.Id.String(),
		Payload: m.Payload,
	}
}

type Queue struct {
	queueType               proto.QueueType
	subscribers             []chan *Message // Buffers for subscribers
	nextSubscriberToRecvIdx int32
	unackedMessages         map[uuid.UUID]*Message // For temporarily tracking unacknowledged messages
}

func NewQueue(queueType proto.QueueType) *Queue {
	return &Queue{
		queueType:               queueType,
		subscribers:             make([]chan *Message, 0),
		nextSubscriberToRecvIdx: 0,
		unackedMessages:         make(map[uuid.UUID]*Message),
	}
}

func (q *Queue) AddSubscriber() chan *Message {
	subscriber := make(chan *Message, 100) // Buffered channel for subscribers
	q.subscribers = append(q.subscribers, subscriber)

	return subscriber
}

func (q *Queue) SendMessage(msg *Message) {
	// Send the message to all subscribers in a round-robin fashion
	for i := 0; i < len(q.subscribers); i++ {
		subscriber := q.subscribers[q.nextSubscriberToRecvIdx%int32(len(q.subscribers))]
		select {
		case subscriber <- msg:
			// Successfully sent message to subscriber
			q.nextSubscriberToRecvIdx = (q.nextSubscriberToRecvIdx + 1) % int32(len(q.subscribers))
			break // Exit after sending to one subscriber
		default:
			// Handle full channel or slow consumer
			log.Printf("Subscriber channel is full, skipping message delivery to subscriber %d", i)
		}
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
	msg := NewMessage(req.Payload)

	for _, queue := range queues {
		queue.SendMessage(msg)
	}

	return &proto.PublishResponse{Success: true}, nil
}

func (b *BrokerServer) Subscribe(req *proto.SubscribeRequest, stream proto.Broker_SubscribeServer) error {
	b.mu.RLock()
	exchange, exists := b.exchanges[req.Exchange]
	if !exists {
		return errors.New("exchange not found")
	}

	queue, exists := exchange.queues[req.Queue]
	if !exists {
		return errors.New("queue not found")
	}

	subscriber := queue.AddSubscriber()
	b.mu.RUnlock()

	for msg := range subscriber {
		// Mark the message as unacknowledged
		queue.unackedMessages[msg.Id] = msg

		if err := stream.Send(msg.ToProto()); err != nil {
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

func (b *BrokerServer) AckMessage(ctx context.Context, req *proto.AckMessageRequest) (*proto.AckMessageResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	exchange, exists := b.exchanges[req.Exchange]
	if !exists {
		return &proto.AckMessageResponse{Success: false}, errors.New("exchange not found")
	}

	queue, exists := exchange.queues[req.Queue]
	if !exists {
		return &proto.AckMessageResponse{Success: false}, errors.New("queue not found")
	}

	// Acknowledge the message by removing it from unackedMessages
	if _, exists := queue.unackedMessages[uuid.MustParse(req.MessageId)]; exists {
		delete(queue.unackedMessages, uuid.MustParse(req.MessageId))
		return &proto.AckMessageResponse{Success: true}, nil
	}

	return &proto.AckMessageResponse{Success: false}, errors.New("message not found or already acknowledged")
}

func (b *BrokerServer) NackMessage(ctx context.Context, req *proto.NackMessageRequest) (*proto.NackMessageResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	exchange, exists := b.exchanges[req.Exchange]
	if !exists {
		return &proto.NackMessageResponse{Success: false}, errors.New("exchange not found")
	}

	queue, exists := exchange.queues[req.Queue]
	if !exists {
		return &proto.NackMessageResponse{Success: false}, errors.New("queue not found")
	}

	// Nack the message by removing it from unackedMessages and potentially re-queuing it
	if msg, exists := queue.unackedMessages[uuid.MustParse(req.MessageId)]; exists {
		delete(queue.unackedMessages, uuid.MustParse(req.MessageId))

		// Optionally, you could re-queue the message or handle it differently based on your requirements
		if req.Requeue {
			// Re-queue the message to the specified queue
			queue.SendMessage(msg)
			log.Printf("Message %s NACKed and re-queued", req.MessageId)
		} else {
			// If not re-queuing, you could log or handle the NACK differently
			log.Printf("Message %s NACKed and not re-queued", req.MessageId)
		}

		return &proto.NackMessageResponse{Success: true}, nil
	}

	return &proto.NackMessageResponse{Success: false}, errors.New("message not found or already acknowledged")
}

func (b *BrokerServer) Disconnect() error {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Clean up all exchanges and queues
	for _, exchange := range b.exchanges {
		for _, queue := range exchange.queues {
			for _, subscriber := range queue.subscribers {
				close(subscriber) // Close each subscriber channel
			}
			queue.subscribers = nil // Clear subscribers
		}
	}
	b.exchanges = make(map[string]*Exchange) // Clear exchanges

	return nil
}

const (
	connectionAddress = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", connectionAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	brokerServer := NewBrokerServer()

	defer func(brokerServer *BrokerServer) {
		err := brokerServer.Disconnect()
		if err != nil {
			log.Printf("error disconnecting broker server: %v", err)
		} else {
			log.Println("Broker server disconnected successfully")
		}
	}(brokerServer)

	proto.RegisterBrokerServer(grpcServer, brokerServer)

	log.Println("Broker server is running on ", connectionAddress)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
