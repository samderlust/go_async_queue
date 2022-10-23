# Async Queue - ensure list of task execute in order

This dart package ensure your pack of task execute in order, one after the other. The queue can be initiated in global scope then used in local module or functions

## Features

- (Normal Queue) Add multiple jobs into queue before firing
- (Auto Queue) Firing job as soon as any job is added to the queue
- (Both) Option to add queue listener that emits events that happen in the queue

### 1. Normal Queue

```
	q := NewQueue()

	q.AddJob(func() { asyncJob(func() { res = append(res, 1) }, 4) })
	q.AddJob(func() { asyncJob(func() { res = append(res, 2) }, 2) })
	q.AddJob(func() { asyncJob(func() { res = append(res, 3) }, 1) })

	q.Start()
```

### 2. Auto Star Queue

- execute job without explicitly call [start()]

```
    q := NewAutoQueue()

	q.AddJob(func() { res = append(res, 1) })

```

### Add Queue Listener

```
	q := NewQueue()
	res := []QueueEvent{}

	q.AddListener(func(event QueueEvent) { fmt.Println(event)})
```
