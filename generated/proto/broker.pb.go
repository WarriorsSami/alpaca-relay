// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/broker.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExchangeType int32

const (
	ExchangeType_DIRECT ExchangeType = 0
	ExchangeType_FANOUT ExchangeType = 1
	ExchangeType_TOPIC  ExchangeType = 2
)

// Enum value maps for ExchangeType.
var (
	ExchangeType_name = map[int32]string{
		0: "DIRECT",
		1: "FANOUT",
		2: "TOPIC",
	}
	ExchangeType_value = map[string]int32{
		"DIRECT": 0,
		"FANOUT": 1,
		"TOPIC":  2,
	}
)

func (x ExchangeType) Enum() *ExchangeType {
	p := new(ExchangeType)
	*p = x
	return p
}

func (x ExchangeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExchangeType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_broker_proto_enumTypes[0].Descriptor()
}

func (ExchangeType) Type() protoreflect.EnumType {
	return &file_proto_broker_proto_enumTypes[0]
}

func (x ExchangeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExchangeType.Descriptor instead.
func (ExchangeType) EnumDescriptor() ([]byte, []int) {
	return file_proto_broker_proto_rawDescGZIP(), []int{0}
}

type QueueType int32

const (
	QueueType_NORMAL  QueueType = 0
	QueueType_DURABLE QueueType = 1
)

// Enum value maps for QueueType.
var (
	QueueType_name = map[int32]string{
		0: "NORMAL",
		1: "DURABLE",
	}
	QueueType_value = map[string]int32{
		"NORMAL":  0,
		"DURABLE": 1,
	}
)

func (x QueueType) Enum() *QueueType {
	p := new(QueueType)
	*p = x
	return p
}

func (x QueueType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueueType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_broker_proto_enumTypes[1].Descriptor()
}

func (QueueType) Type() protoreflect.EnumType {
	return &file_proto_broker_proto_enumTypes[1]
}

func (x QueueType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueueType.Descriptor instead.
func (QueueType) EnumDescriptor() ([]byte, []int) {
	return file_proto_broker_proto_rawDescGZIP(), []int{1}
}

type PublishRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Exchange      string                 `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	RoutingKey    string                 `protobuf:"bytes,2,opt,name=routing_key,json=routingKey,proto3" json:"routing_key,omitempty"`
	Payload       string                 `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	mi := &file_proto_broker_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_broker_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_proto_broker_proto_rawDescGZIP(), []int{0}
}

func (x *PublishRequest) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *PublishRequest) GetRoutingKey() string {
	if x != nil {
		return x.RoutingKey
	}
	return ""
}

func (x *PublishRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type PublishResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublishResponse) Reset() {
	*x = PublishResponse{}
	mi := &file_proto_broker_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishResponse) ProtoMessage() {}

func (x *PublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_broker_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishResponse.ProtoReflect.Descriptor instead.
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return file_proto_broker_proto_rawDescGZIP(), []int{1}
}

func (x *PublishResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type SubscribeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Exchange      string                 `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Queue         string                 `protobuf:"bytes,2,opt,name=queue,proto3" json:"queue,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	mi := &file_proto_broker_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_broker_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_proto_broker_proto_rawDescGZIP(), []int{2}
}

func (x *SubscribeRequest) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *SubscribeRequest) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Payload       string                 `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_proto_broker_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_broker_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_broker_proto_rawDescGZIP(), []int{3}
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type DeclareExchangeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Exchange      string                 `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Type          ExchangeType           `protobuf:"varint,2,opt,name=type,proto3,enum=broker.ExchangeType" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeclareExchangeRequest) Reset() {
	*x = DeclareExchangeRequest{}
	mi := &file_proto_broker_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeclareExchangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeclareExchangeRequest) ProtoMessage() {}

func (x *DeclareExchangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_broker_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeclareExchangeRequest.ProtoReflect.Descriptor instead.
func (*DeclareExchangeRequest) Descriptor() ([]byte, []int) {
	return file_proto_broker_proto_rawDescGZIP(), []int{4}
}

func (x *DeclareExchangeRequest) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *DeclareExchangeRequest) GetType() ExchangeType {
	if x != nil {
		return x.Type
	}
	return ExchangeType_DIRECT
}

type DeclareExchangeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeclareExchangeResponse) Reset() {
	*x = DeclareExchangeResponse{}
	mi := &file_proto_broker_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeclareExchangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeclareExchangeResponse) ProtoMessage() {}

func (x *DeclareExchangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_broker_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeclareExchangeResponse.ProtoReflect.Descriptor instead.
func (*DeclareExchangeResponse) Descriptor() ([]byte, []int) {
	return file_proto_broker_proto_rawDescGZIP(), []int{5}
}

func (x *DeclareExchangeResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeclareQueueRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Queue         string                 `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
	Exchange      string                 `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Type          QueueType              `protobuf:"varint,3,opt,name=type,proto3,enum=broker.QueueType" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeclareQueueRequest) Reset() {
	*x = DeclareQueueRequest{}
	mi := &file_proto_broker_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeclareQueueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeclareQueueRequest) ProtoMessage() {}

func (x *DeclareQueueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_broker_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeclareQueueRequest.ProtoReflect.Descriptor instead.
func (*DeclareQueueRequest) Descriptor() ([]byte, []int) {
	return file_proto_broker_proto_rawDescGZIP(), []int{6}
}

func (x *DeclareQueueRequest) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

func (x *DeclareQueueRequest) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *DeclareQueueRequest) GetType() QueueType {
	if x != nil {
		return x.Type
	}
	return QueueType_NORMAL
}

type DeclareQueueResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeclareQueueResponse) Reset() {
	*x = DeclareQueueResponse{}
	mi := &file_proto_broker_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeclareQueueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeclareQueueResponse) ProtoMessage() {}

func (x *DeclareQueueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_broker_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeclareQueueResponse.ProtoReflect.Descriptor instead.
func (*DeclareQueueResponse) Descriptor() ([]byte, []int) {
	return file_proto_broker_proto_rawDescGZIP(), []int{7}
}

func (x *DeclareQueueResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type AckMessageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MessageId     string                 `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Queue         string                 `protobuf:"bytes,2,opt,name=queue,proto3" json:"queue,omitempty"`
	Exchange      string                 `protobuf:"bytes,3,opt,name=exchange,proto3" json:"exchange,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AckMessageRequest) Reset() {
	*x = AckMessageRequest{}
	mi := &file_proto_broker_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AckMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckMessageRequest) ProtoMessage() {}

func (x *AckMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_broker_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckMessageRequest.ProtoReflect.Descriptor instead.
func (*AckMessageRequest) Descriptor() ([]byte, []int) {
	return file_proto_broker_proto_rawDescGZIP(), []int{8}
}

func (x *AckMessageRequest) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *AckMessageRequest) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

func (x *AckMessageRequest) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

type AckMessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AckMessageResponse) Reset() {
	*x = AckMessageResponse{}
	mi := &file_proto_broker_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AckMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckMessageResponse) ProtoMessage() {}

func (x *AckMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_broker_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckMessageResponse.ProtoReflect.Descriptor instead.
func (*AckMessageResponse) Descriptor() ([]byte, []int) {
	return file_proto_broker_proto_rawDescGZIP(), []int{9}
}

func (x *AckMessageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type NackMessageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MessageId     string                 `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Queue         string                 `protobuf:"bytes,2,opt,name=queue,proto3" json:"queue,omitempty"`
	Exchange      string                 `protobuf:"bytes,3,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Requeue       bool                   `protobuf:"varint,4,opt,name=requeue,proto3" json:"requeue,omitempty"` // If true, the message will be re-queued
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NackMessageRequest) Reset() {
	*x = NackMessageRequest{}
	mi := &file_proto_broker_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NackMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NackMessageRequest) ProtoMessage() {}

func (x *NackMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_broker_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NackMessageRequest.ProtoReflect.Descriptor instead.
func (*NackMessageRequest) Descriptor() ([]byte, []int) {
	return file_proto_broker_proto_rawDescGZIP(), []int{10}
}

func (x *NackMessageRequest) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *NackMessageRequest) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

func (x *NackMessageRequest) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *NackMessageRequest) GetRequeue() bool {
	if x != nil {
		return x.Requeue
	}
	return false
}

type NackMessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NackMessageResponse) Reset() {
	*x = NackMessageResponse{}
	mi := &file_proto_broker_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NackMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NackMessageResponse) ProtoMessage() {}

func (x *NackMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_broker_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NackMessageResponse.ProtoReflect.Descriptor instead.
func (*NackMessageResponse) Descriptor() ([]byte, []int) {
	return file_proto_broker_proto_rawDescGZIP(), []int{11}
}

func (x *NackMessageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_broker_proto protoreflect.FileDescriptor

const file_proto_broker_proto_rawDesc = "" +
	"\n" +
	"\x12proto/broker.proto\x12\x06broker\"g\n" +
	"\x0ePublishRequest\x12\x1a\n" +
	"\bexchange\x18\x01 \x01(\tR\bexchange\x12\x1f\n" +
	"\vrouting_key\x18\x02 \x01(\tR\n" +
	"routingKey\x12\x18\n" +
	"\apayload\x18\x03 \x01(\tR\apayload\"+\n" +
	"\x0fPublishResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"D\n" +
	"\x10SubscribeRequest\x12\x1a\n" +
	"\bexchange\x18\x01 \x01(\tR\bexchange\x12\x14\n" +
	"\x05queue\x18\x02 \x01(\tR\x05queue\"3\n" +
	"\aMessage\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x18\n" +
	"\apayload\x18\x02 \x01(\tR\apayload\"^\n" +
	"\x16DeclareExchangeRequest\x12\x1a\n" +
	"\bexchange\x18\x01 \x01(\tR\bexchange\x12(\n" +
	"\x04type\x18\x02 \x01(\x0e2\x14.broker.ExchangeTypeR\x04type\"3\n" +
	"\x17DeclareExchangeResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"n\n" +
	"\x13DeclareQueueRequest\x12\x14\n" +
	"\x05queue\x18\x01 \x01(\tR\x05queue\x12\x1a\n" +
	"\bexchange\x18\x02 \x01(\tR\bexchange\x12%\n" +
	"\x04type\x18\x03 \x01(\x0e2\x11.broker.QueueTypeR\x04type\"0\n" +
	"\x14DeclareQueueResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"d\n" +
	"\x11AckMessageRequest\x12\x1d\n" +
	"\n" +
	"message_id\x18\x01 \x01(\tR\tmessageId\x12\x14\n" +
	"\x05queue\x18\x02 \x01(\tR\x05queue\x12\x1a\n" +
	"\bexchange\x18\x03 \x01(\tR\bexchange\".\n" +
	"\x12AckMessageResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"\x7f\n" +
	"\x12NackMessageRequest\x12\x1d\n" +
	"\n" +
	"message_id\x18\x01 \x01(\tR\tmessageId\x12\x14\n" +
	"\x05queue\x18\x02 \x01(\tR\x05queue\x12\x1a\n" +
	"\bexchange\x18\x03 \x01(\tR\bexchange\x12\x18\n" +
	"\arequeue\x18\x04 \x01(\bR\arequeue\"/\n" +
	"\x13NackMessageResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess*1\n" +
	"\fExchangeType\x12\n" +
	"\n" +
	"\x06DIRECT\x10\x00\x12\n" +
	"\n" +
	"\x06FANOUT\x10\x01\x12\t\n" +
	"\x05TOPIC\x10\x02*$\n" +
	"\tQueueType\x12\n" +
	"\n" +
	"\x06NORMAL\x10\x00\x12\v\n" +
	"\aDURABLE\x10\x012\xaa\x03\n" +
	"\x06Broker\x12:\n" +
	"\aPublish\x12\x16.broker.PublishRequest\x1a\x17.broker.PublishResponse\x128\n" +
	"\tSubscribe\x12\x18.broker.SubscribeRequest\x1a\x0f.broker.Message0\x01\x12R\n" +
	"\x0fDeclareExchange\x12\x1e.broker.DeclareExchangeRequest\x1a\x1f.broker.DeclareExchangeResponse\x12I\n" +
	"\fDeclareQueue\x12\x1b.broker.DeclareQueueRequest\x1a\x1c.broker.DeclareQueueResponse\x12C\n" +
	"\n" +
	"AckMessage\x12\x19.broker.AckMessageRequest\x1a\x1a.broker.AckMessageResponse\x12F\n" +
	"\vNackMessage\x12\x1a.broker.NackMessageRequest\x1a\x1b.broker.NackMessageResponseB\x11Z\x0fgenerated/protob\x06proto3"

var (
	file_proto_broker_proto_rawDescOnce sync.Once
	file_proto_broker_proto_rawDescData []byte
)

func file_proto_broker_proto_rawDescGZIP() []byte {
	file_proto_broker_proto_rawDescOnce.Do(func() {
		file_proto_broker_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_broker_proto_rawDesc), len(file_proto_broker_proto_rawDesc)))
	})
	return file_proto_broker_proto_rawDescData
}

var file_proto_broker_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_broker_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_broker_proto_goTypes = []any{
	(ExchangeType)(0),               // 0: broker.ExchangeType
	(QueueType)(0),                  // 1: broker.QueueType
	(*PublishRequest)(nil),          // 2: broker.PublishRequest
	(*PublishResponse)(nil),         // 3: broker.PublishResponse
	(*SubscribeRequest)(nil),        // 4: broker.SubscribeRequest
	(*Message)(nil),                 // 5: broker.Message
	(*DeclareExchangeRequest)(nil),  // 6: broker.DeclareExchangeRequest
	(*DeclareExchangeResponse)(nil), // 7: broker.DeclareExchangeResponse
	(*DeclareQueueRequest)(nil),     // 8: broker.DeclareQueueRequest
	(*DeclareQueueResponse)(nil),    // 9: broker.DeclareQueueResponse
	(*AckMessageRequest)(nil),       // 10: broker.AckMessageRequest
	(*AckMessageResponse)(nil),      // 11: broker.AckMessageResponse
	(*NackMessageRequest)(nil),      // 12: broker.NackMessageRequest
	(*NackMessageResponse)(nil),     // 13: broker.NackMessageResponse
}
var file_proto_broker_proto_depIdxs = []int32{
	0,  // 0: broker.DeclareExchangeRequest.type:type_name -> broker.ExchangeType
	1,  // 1: broker.DeclareQueueRequest.type:type_name -> broker.QueueType
	2,  // 2: broker.Broker.Publish:input_type -> broker.PublishRequest
	4,  // 3: broker.Broker.Subscribe:input_type -> broker.SubscribeRequest
	6,  // 4: broker.Broker.DeclareExchange:input_type -> broker.DeclareExchangeRequest
	8,  // 5: broker.Broker.DeclareQueue:input_type -> broker.DeclareQueueRequest
	10, // 6: broker.Broker.AckMessage:input_type -> broker.AckMessageRequest
	12, // 7: broker.Broker.NackMessage:input_type -> broker.NackMessageRequest
	3,  // 8: broker.Broker.Publish:output_type -> broker.PublishResponse
	5,  // 9: broker.Broker.Subscribe:output_type -> broker.Message
	7,  // 10: broker.Broker.DeclareExchange:output_type -> broker.DeclareExchangeResponse
	9,  // 11: broker.Broker.DeclareQueue:output_type -> broker.DeclareQueueResponse
	11, // 12: broker.Broker.AckMessage:output_type -> broker.AckMessageResponse
	13, // 13: broker.Broker.NackMessage:output_type -> broker.NackMessageResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_broker_proto_init() }
func file_proto_broker_proto_init() {
	if File_proto_broker_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_broker_proto_rawDesc), len(file_proto_broker_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_broker_proto_goTypes,
		DependencyIndexes: file_proto_broker_proto_depIdxs,
		EnumInfos:         file_proto_broker_proto_enumTypes,
		MessageInfos:      file_proto_broker_proto_msgTypes,
	}.Build()
	File_proto_broker_proto = out.File
	file_proto_broker_proto_goTypes = nil
	file_proto_broker_proto_depIdxs = nil
}
