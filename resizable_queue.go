package resizable_array_queue

import "errors"

type Queue struct {
	size    int
	head    int
	tail    int
	storage []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		size:    0,
		head:    0,
		tail:    0,
		storage: make([]interface{}, 2),
	}
}

func (q *Queue) Enqueue(element interface{}) error {
	if q.size != 0 && q.size == len(q.storage) {
		q.resize(q.size * 2)
	}
	q.storage[q.tail] = element
	q.tail = q.tail + 1
	q.size = q.size + 1
	if q.tail == len(q.storage) {
		q.tail = 0 //wrap around
	}
	return nil
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.size == 0 {
		return "", errors.New("Trying to dequeue empty queue")
	}
	value := q.storage[q.head]
	q.storage[q.head] = nil
	q.head = q.head + 1
	q.size = q.size - 1
	if q.size != 0 && q.size <= (len(q.storage)/4) {
		q.resize(len(q.storage) / 2)
	}
	if q.head == len(q.storage) {
		q.head = 0 //wrap around
	}
	return value, nil
}

func (q *Queue) resize(size int) {
	newStorage := make([]interface{}, size)
	for i := 0; i < q.size; i++ {
		newStorage[i] = q.storage[q.head+i%len(q.storage)]
	}
	q.head = 0
	q.tail = q.size
	q.storage = newStorage
}
