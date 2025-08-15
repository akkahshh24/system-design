package mq

import (
	"sync/atomic"
)

type SubscriberOffset struct {
	subscriber ISubscriber
	offset     atomic.Int64
}

func NewSubscriberOffset(subscriber ISubscriber) *SubscriberOffset {
	return &SubscriberOffset{
		subscriber: subscriber,
		offset:     atomic.Int64{},
	}
}
