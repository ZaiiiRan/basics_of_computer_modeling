package queue

import (
	"container/list"
)

type Queue struct {
	queue *list.List
}

func NewQueue() *Queue {
	return &Queue{
		queue: list.New(),
	}
}

func (q *Queue) Enqueue(element any) {
	q.queue.PushBack(element)
}

func (q *Queue) Dequeue() any {
	front := q.queue.Front()
	if front != nil {
		q.queue.Remove(front)
		return front.Value
	}
	return nil
}

func (q *Queue) Size() int {
	return q.queue.Len()
}

func (q *Queue) Front() any {
	front := q.queue.Front()
	if front != nil {
		return front.Value
	}
	return nil
}
