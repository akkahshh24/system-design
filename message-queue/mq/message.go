package mq

type Message struct {
	ID      string `json:"id"`
	Topic   string `json:"topic"`
	Payload string `json:"payload"`
	// Additional metadata can be added here
}
