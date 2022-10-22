package asyncqueue

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestOrder(t *testing.T) {
	q := NewQueue()
	res := []int{}

	q.AddListener(func(event QueueEvent) { fmt.Println(event) })

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

func asyncJob(job func(), delay time.Duration) {
	time.Sleep(delay * time.Second)
	job()
}
