package mq

type TopicHandler struct {
	topic               *Topic
	subscriptionWorkers map[string]SubscriptionWorker
}

func NewTopicHandler(topic *Topic) *TopicHandler {
	return &TopicHandler{
		topic:               topic,
		subscriptionWorkers: make(map[string]SubscriptionWorker),
	}
}
