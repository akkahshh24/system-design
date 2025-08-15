package mq

type Publisher struct{}

func NewPublisher(addr string, topicName string) *Publisher {
	return &Publisher{}
}
