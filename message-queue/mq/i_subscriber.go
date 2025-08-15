package mq

// Public interface for the message queue subscriber
type ISubscriber interface {
	// Consume messages from the queue
	Consume() error
}
