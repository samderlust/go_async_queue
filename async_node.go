package asyncqueue

type AsyncJob func()

type AsyncNode struct {
	Job  AsyncJob
	Next *AsyncNode
}
