package mq

import (
	"errors"
)

type MessageQueue struct {
	addr          string
	topicHandlers map[string]*TopicHandler
}

func NewMessageQueue(addr string) *MessageQueue {
	return &MessageQueue{
		addr:          addr,
		topicHandlers: make(map[string]*TopicHandler),
	}
}

func (q *MessageQueue) CreateTopic(topicName string) (*Topic, error) {
	if _, exists := q.topicHandlers[topicName]; exists {
		return nil, errors.New("topic already exists")
	}

	topic := NewTopic(topicName)
	q.topicHandlers[topicName] = NewTopicHandler(topic)
	return topic, nil
}

func (q *MessageQueue) Subscribe(topic *Topic, subscriber ISubscriber) {
	subscriberOffset := NewSubscriberOffset(subscriber)
	topic.AddSubscriber(subscriberOffset)
}

func (q *MessageQueue) Publish(topic *Topic, message *Message) error {