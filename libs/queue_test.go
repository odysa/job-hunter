package libs

import (
	"testing"
)

type QueueTest []struct {
	queue Queue
	ans   int
}

var testData = QueueTest{
	{Queue{}, 0},
	{Queue{1}, 1},
	{Queue{1, 2}, 2},
	{Queue{1, 2, 3}, 3},
	{Queue{1, 2, 3, 4}, 4},
	{Queue{1, 2, 3, 4, 5}, 5},
	{Queue{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, 13},
}

const itemLimit = 10000

func TestQueue_IsEmpty(t *testing.T) {
	queue := Queue{1, 2}
	if queue.IsEmpty() {
		t.Errorf("queue should not be empty")
	}
	queue = Queue{}
	if !queue.IsEmpty() {
		t.Errorf("queue should be empty")
	}
}
func TestQueue_Length(t *testing.T) {
	for _, test := range testData {
		if test.queue.Length() != test.ans {
			t.Errorf("length error, queue length should be %d, but got %d", test.queue.Length(), test.ans)
		}
	}
	queue := Queue{}
	for i := 0; i < itemLimit; i++ {
		queue.Push(i)
		if queue.Length() != i+1 {
			t.Errorf("Length should be %d,but got %d", i+1, queue.Length())
		}
	}
}
func TestQueue_Pop(t *testing.T) {
	for _, test := range testData {
		queue := test.queue
		for _, item := range test.queue {
			pop := queue.Pop()
			if item != pop {
				t.Errorf("Pop error, should be %v, but got %v", item, pop)
			}
		}
	}
}
func TestQueue_Push(t *testing.T) {
	queue := Queue{}
	for i := 0; i < itemLimit; i++ {
		queue.Push(i)
		if queue[queue.Length()-1] != i {
			t.Errorf("Push error, last item should be %v, but got %v", i, queue[queue.Length()-1])
		}
	}
}
func TestQueue_Top(t *testing.T) {
	queue := Queue{}
	for i := 0; i < itemLimit; i++ {
		queue.Push(i)
		if queue.Top() != i {
			t.Errorf("Top error, it should be %v, but got %v", i, queue.Top())
		}
		queue.Pop()
	}
}
