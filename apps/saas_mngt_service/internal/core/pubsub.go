//go:generate sh -c "mockgen -source=$GOFILE -destination=mock/$(basename $GOFILE .go)_test.go -package=mock"
package core

import "context"

type Message struct {
	// Idempotency ID to ensure message is processed only once.
	ID string
	// Topic is the name of the topic to which the message belongs.
	Topic string
	// Payload is the actual data to be sent in the message.
	Payload []byte
	// PartitionKey is used to determine the partition for the message when using kafka.
	PartitionKey string
}

type Publisher interface {
	// Publish sends a message to the specified topic.
	Publish(ctx context.Context, messages ...*Message) error
}

type Subscriber interface {
	// Subscribe listens for messages on the specified topic and processes them.
	Subscribe(ctx context.Context, topic string, handler func(msg *Message) error) error
}
