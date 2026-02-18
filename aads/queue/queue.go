package main

import "fmt"

type IntQueue struct {
	buf  []int
	head int
	tail int
	size int
}

func NewIntQueue(cap int) IntQueue {
	if cap < 1 {
		cap = 1
	}
	return IntQueue{buf: make([]int, cap)}
}

func (q *IntQueue) Push(v int) {
	if q.size == len(q.buf) {
		q.grow()
	}
	q.buf[q.tail] = v
	q.tail = (q.tail + 1) % len(q.buf)
	q.size++
}

func (q *IntQueue) Pop() (int, bool) {
	if q.size == 0 {
		return 0, false
	}
	v := q.buf[q.head]
	q.head = (q.head + 1) % len(q.buf)
	q.size--
	return v, true
}

func (q *IntQueue) grow() {
	newBuf := make([]int, len(q.buf)*2)
	for i := 0; i < q.size; i++ {
		newBuf[i] = q.buf[(q.head+i)%len(q.buf)]
	}
	q.buf = newBuf
	q.head = 0
	q.tail = q.size
}

func (q *IntQueue) Print() {
	for i := 0; i < q.size; i++ {
		fmt.Println(q.buf[(q.head+i)%len(q.buf)])
	}
}

func main() {
	queue := NewIntQueue(2)
	queue.Push(0)
	queue.Push(1)
	queue.Push(2)
	queue.Print()
	fmt.Println("-")
	queue.Pop()
	queue.Print()
}
