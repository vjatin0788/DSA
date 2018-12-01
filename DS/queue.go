package DS

import (
	"fmt"
)

var queueKey int64

type QueueImp interface {
	EnQueue(data interface{}) bool
	DeQueue(data interface{}) bool
}

type Queue struct {
	QueueBody []interface{}
	Front     int32
	Rear      int32
	Size      int32
	Capacity  int32
}

func (q *Queue) EnQueue(data interface{}) bool {
	if data != nil {
		if q.Size < q.Capacity {
			q.Rear = (q.Rear + 1) % q.Capacity
			q.QueueBody[q.Rear] = data
			q.Size++
			return true
		} else {
			fmt.Println("OverFlow")
			return false
		}

	}
	return false
}

func (q *Queue) EnQueueFront(data interface{}) bool {
	if data != nil {
		if q.Size < q.Capacity {
			q.Rear = (q.Rear + 1) % q.Capacity
			q.QueueBody[q.Rear] = data
			q.Size++
			return true
		} else {
			fmt.Println("OverFlow")
			return false
		}

	}
	return false
}

func (q *Queue) DeQueue() interface{} {
	if q.Size > 0 {
		q.Front = (q.Front + 1) % q.Capacity
		data := q.QueueBody[q.Front%q.Capacity]
		q.QueueBody[q.Front%q.Capacity] = nil
		q.Size--
		return data
	} else {
		fmt.Println("UnderFlow")
	}
	return nil
}

func (q *Queue) IsEmpty() bool {
	if q.Size == 0 {
		return true
	}
	return false
}

func (q *Queue) ResetQueue() {
	for i := q.Front; i < q.Size; i++ {
		q.QueueBody[i] = nil
	}
	q.Front = -1
	q.Rear = -1
	q.Size = 0
}

func InitQueue(capacity int32) *Queue {

	queueImp := Queue{
		Front:     -1,
		Rear:      -1,
		Size:      0,
		Capacity:  capacity,
		QueueBody: make([]interface{}, capacity),
	}
	return &queueImp
}

func InitHeapQueue(capacity int64) *HeapNode {
	arr := make([]int64, capacity)
	key := make([]int64, capacity)
	queueKey = capacity
	return &HeapNode{
		Capacity: capacity,
		Arr:      arr,
		Key:      key,
	}
}

func (heap *HeapNode) Enqueue(val int64) {
	heap.InsertKey(queueKey, val)
	queueKey--
}
func (heap *HeapNode) Dequeue() int64 {
	_, val := heap.DeleteKey()

	return val
}
