# AlpacaRelay Message Broker

AlpacaRelay is a toy message broker inspired by the design of RabbitMQ (especially the AMQP 0-9-1 protocol implementation) and built using Golang, gRPC and Sqlite. It is designed to be simple and easy to use, with a focus on learning and experimentation.

## MVP Features
- [x] Support for exchanges, queues, and bindings
- [x] Basic publish/subscribe messaging
- [x] gRPC interface for communication
- [x] Support for declaring exchanges and queues
- [x] Round-robin message distribution among consumers bound to a queue
- [x] Support for multiple exchange types (direct, fanout, topic)
- [x] Support for acknowledgements, negative acknowledgements and message re-queuing

## Roadmap
- [ ] Durable exchanges, queues and messages using Sqlite and sqlc
- [ ] Prefetching messages for consumers
- [ ] Dead-letter exchanges and queues
- [ ] Message TTL (Time-To-Live)
- [ ] Message priorities and ordering (using FIFO or priority queues)
- [ ] Support for other message protocols (e.g., MQTT, STOMP)
- [ ] Monitoring and management interface (e.g., Prometheus, Grafana)
- [ ] Authentication and authorization mechanisms
- [ ] Support for message headers and properties
- [ ] Support for transactions and message batching
- [ ] Plugin system for extending functionality
- [ ] Support for multiple message formats (e.g., JSON, Protobuf, CloudEvents)
- [ ] Support for message compression and encryption
- [ ] Support for delayed message delivery
- [ ] Clustering support for high availability and scalability