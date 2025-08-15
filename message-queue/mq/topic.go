package mq

import (
	"sync"
)

type Topic struct {
	name        string
	messages    []Message
	subscribers []*SubscriberOffset
	mutex       sync.Mutex
	// Additional metadata can be added here
	// like partitions, replication factor, retention time, etc.
}

func NewTopic(name string) *Topic {
	return &Topic{
		name:        name,
		messages:    make([]Message, 0),
		subscribers: make([]*SubscriberOffset, 0),
		mutex:       sync.Mutex{},
	}
}

func (t *Topic) Name() string {
	return t.name
}

func (t *Topic) AddSubscriber(subscriberOffset *SubscriberOffset) {
	t.subscribers = append(t.subscribers, subscriberOffset)
}

// func (t *Topic) AddMessage(message Message) {
// 	// Add sanity checks

// 	t.Mutex.Lock()
// 	defer t.Mutex.Unlock()

// 	// Add the message to the topic's messages
// 	t.Messages = append(t.Messages, message)
// }

// func (t *Topic) AddSubscriber(subscriber Subscriber) {
// 	// Add sanity checks

// 	t.Mutex.Lock()
// 	defer t.Mutex.Unlock()

// 	// Add the subscriber to the topic's subscribers
// 	t.Subscribers = append(t.Subscribers, subscriber)
// }
