package pkg

import (
	"container/list"
)

type Queue struct {
	v *list.List
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

// Enqueue adds value to the end of the queue
func (q *Queue) Enqueue(val interface{}) {
	q.v.PushBack(val)
}

// Dequeue removes value from the front of the queue and returns it
func (q *Queue) Dequeue() interface{} {
	e := q.v.Front() // first node
	if e != nil {
		q.v.Remove(e) // remove node
		return e
	}
	return nil
}
