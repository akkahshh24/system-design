package messagequeue

import "message-queue/mq"

func main() {
	// create a new message queue
	queue := mq.NewMessageQueue("localhost:9092")

	// Create topics
	topic1, err := queue.CreateTopic("topic1")
	if err != nil {
		panic(err)
	}

	topic2, err := queue.CreateTopic("topic2")
	if err != nil {
		panic(err)
	}

	// Create subscribers
	subscriber1 := NewDemoSubscriber("subscriber1")
	subscriber2 := NewDemoSubscriber("subscriber2")

	// Subscribe to topics
	queue.Subscribe(topic1, subscriber1)
	queue.Subscribe(topic2, subscriber2)

	// Publish messages
	err = queue.Publish(topic1, "Hello from topic1")
	if err != nil {
		panic(err)
	}
	err = queue.Publish(topic2, "Hello from topic2")
	if err != nil {
		panic(err)
	}
}
