package asyncqueue

import (
	"fmt"
	"time"
)

type QueueEvent struct {
	timeStamp time.Time
	eventType queueEventType
}

func (e QueueEvent) String() string {
	return fmt.Sprintf("event type [%v] at %v", e.eventType, e.timeStamp)
}

type queueEventType string

var QueueEventTypes = struct {
	NewJobAdded queueEventType
	QueueStart  queueEventType
	QueueEnd    queueEventType
	BeforeJob   queueEventType
	AfterJob    queueEventType
	QueueClosed queueEventType
}{
	NewJobAdded: "newJobAdded",
	QueueStart:  "queueStart",
	QueueEnd:    "queueEnd",
	BeforeJob:   "beforeJob",
	AfterJob:    "afterJob",
	QueueClosed: "queueClosed",
}
