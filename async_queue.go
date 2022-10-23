package asyncqueue

import "time"

type asyncQueue struct {
	size      int
	first     *AsyncNode
	last      *AsyncNode
	autoRun   bool
	isRunning bool
	listener  func(event QueueEvent)
}

type AsyncQueue interface {
	AddJob(job AsyncJob)
	Start()
	AddListener(func(event QueueEvent))
	Size() int
	enqueue(node AsyncNode)
	dequeue()
	emitEvent(eventType queueEventType)
}

// Add new job into the queue
func (a *asyncQueue) AddJob(job AsyncJob) {
	node := AsyncNode{
		Job: job,
	}

	a.enqueue(node)

	if a.autoRun {
		a.Start()
	}
}

// Start implements AsyncQueueI
func (a *asyncQueue) Start() {
	if a.size == 0 || a.isRunning {
		return
	}

	a.isRunning = true
	a.emitEvent(QueueEventTypes.QueueStart)

	for a.size > 0 {
		a.dequeue()
	}

	a.isRunning = false
	a.emitEvent(QueueEventTypes.QueueEnd)
}

// dequeue implements AsyncQueueI
func (a *asyncQueue) dequeue() {
	if a.first == nil {
		return
	}

	currentNode := a.first

	a.emitEvent(QueueEventTypes.BeforeJob)

	currentNode.Job()

	if a.size == 1 {
		a.first = nil
		a.last = nil
	} else {
		a.first = currentNode.Next
		currentNode.Next = nil
	}

	a.size--
	a.emitEvent(QueueEventTypes.AfterJob)
}

// add a new node into the queue
func (a *asyncQueue) enqueue(node AsyncNode) {
	if a.first == nil {
		a.first = &node
		a.last = &node
	} else {
		a.last.Next = &node
		a.last = &node
	}

	a.size++
	a.emitEvent(QueueEventTypes.NewJobAdded)
}

func (a *asyncQueue) emitEvent(eventType queueEventType) {
	if a.listener != nil {
		event := QueueEvent{
			timeStamp: time.Now(),
			eventType: eventType,
		}
		a.listener(event)
	}
}
func (a *asyncQueue) AddListener(listener func(event QueueEvent)) {
	a.listener = listener
}

func (a *asyncQueue) Size() int {
	return a.size
}

func NewQueue() AsyncQueue {
	var q AsyncQueue = &asyncQueue{}
	return q
}
func NewAutoQueue() AsyncQueue {
	var q AsyncQueue = &asyncQueue{
		autoRun: true,
	}
	return q
}
