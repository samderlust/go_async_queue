package asyncqueue

import (
	"reflect"
	"testing"
	"time"
)

func TestOrder(t *testing.T) {
	q := NewQueue()
	res := []int{}

	q.AddJob(func() { asyncJob(func() { res = append(res, 1) }, 4) })
	q.AddJob(func() { asyncJob(func() { res = append(res, 2) }, 2) })
	q.AddJob(func() { asyncJob(func() { res = append(res, 3) }, 1) })

	if reflect.DeepEqual(res, []int{1, 2, 3}) {
		t.Errorf("slices not match")
	}

	if q.Size() != 3 {
		t.Errorf("Size not match 3")
	}

	q.Start()

	if q.Size() != 0 {
		t.Errorf("Size not match 0 ")
	}

}

func TestQueueEventEmit(t *testing.T) {
	q := NewQueue()
	res := []QueueEvent{}

	q.AddListener(func(event QueueEvent) { res = append(res, event) })

	q.AddJob(func() {})

	if res[len(res)-1].eventType != QueueEventTypes.NewJobAdded {
		t.Errorf("Expect QueueEventTypes.NewJobAdded")
	}

	q.AddJob(func() {})

	q.Start()

	if len(res) == 5 {
		t.Errorf("Size not match 5 ")
	}

}

func TestAutoQueue(t *testing.T) {
	q := NewAutoQueue()
	res := []int{}

	q.AddJob(func() { res = append(res, 1) })

	if res[len(res)-1] != 1 {
		t.Errorf("expected 1")
	}
	q.AddJob(func() { res = append(res, 2) })

	if res[len(res)-1] != 2 {
		t.Errorf("expected 1")
	}

}

func asyncJob(job func(), delay time.Duration) {
	time.Sleep(delay * time.Second)
	job()
}
