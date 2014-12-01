package resizable_array_queue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue("a")
	q.Enqueue("a")
	q.Enqueue("b")
	assert.Equal(t, 4, len(q.storage))
	q.Enqueue("b")
	q.Enqueue("c")
	assert.Equal(t, 8, len(q.storage))

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

	assert.Equal(t, 4, len(q.storage))

	fmt.Println(q.Dequeue())
	assert.Equal(t, 2, len(q.storage))

	fmt.Println(q.Dequeue())
	assert.Equal(t, 2, len(q.storage))

	q.Enqueue(2)
	assert.Equal(t, 2, len(q.storage))
}
