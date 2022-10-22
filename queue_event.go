package asyncqueue

import (
	"fmt"
	"time"
)

type QueueEvent struct {
	timeStamp time.Time
	eventType QueueEventType
}

func (e QueueEvent) String() string {
	return fmt.Sprintf("event type [%v] at %v", e.eventType, e.timeStamp)
}

type QueueEventType string

const (
	NewJobAdded QueueEventType = "newJobAdded"
	QueueStart  QueueEventType = "queueStart"
	QueueEnd    QueueEventType = "queueEnd"
	BeforeJob   QueueEventType = "beforeJob"
	AfterJob    QueueEventType = "afterJob"
	QueueClosed QueueEventType = "queueClosed"
)
