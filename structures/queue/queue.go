package queue

import "fmt"

type item struct {
	data int
	next *item
}
type queue struct {
	count int
	first *item
	last *item
}

func main() {
	q := InitQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Print()
	q.Dequeue()
	q.Print()
	q.Dequeue()
	q.Print()
	q.Dequeue()
	q.Print()
}

func (q *queue) Enqueue(data int) {
	i := initItem(data)
	if q.count == 0 {
		q.first = i
		q.last = i
	} else {
		q.last.next = i
		q.last = i
	}
	q.count = q.count + 1
}

func (q *queue) Dequeue() (int) {
	if q.count == 0 {
		panic("Queue is empty")
	}
	i := q.first
	q.first = i.next
	q.count = q.count - 1
	return i.data
}

func (q *queue) Peek() (*item) {
	if q.count == 0 {
		return nil
	}
	return q.first
}

// Output each item in the queue
func (q *queue) Print() {
	if q.count == 0 {
		out("queue is empty\n")
		return
	}
	i := q.first
	for count := 0; i != nil && count < q.count; count++ {
		out("%d ", i.data)
		i = i.next
	}
	out("\n")
}

func InitQueue() (*queue) {
	q := new(queue)
	q.count = 0
	q.first = nil
	q.last = nil
	return q
}
func initItem(data int) (*item) {
	i := new(item)
	i.data = data
	i.next = nil
	return i
}

func out(format string, v...interface{}) {
	fmt.Printf(format, v...)
}