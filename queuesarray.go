package main

import "fmt"

type qarray struct {
	Q    []int
	size int
}

func (q *qarray) len() int {
	return q.size
}

func (q *qarray) isempty() bool {
	length := len(q.Q)
	return length == 0
}

func (q *qarray) enqueue(e int) {
	q.Q = append(q.Q, e)
	q.size++
}
func (q *qarray) dequeue() int {
	if q.isempty() {
		fmt.Println("Queue is empty")
		return 0
	} else {
		e := q.Q[0]
		q.Q = q.Q[1:]
		q.size--
		return e
	}
}

func (q *qarray) print() {
	for i := 0; i < q.len(); i++ {
		fmt.Println(q.Q[i])
	}
}
func (q *qarray) front() int {
	return q.Q[0]
}
func main() {
	s := qarray{}
	s.enqueue(1)
	s.enqueue(2)
	s.enqueue(3)
	s.enqueue(4)
	// el := s.dequeue()
	// fmt.Println(el)
	fmt.Println(s.size)
	s.print()
	f := s.front()
	fmt.Println(f)
	a := s.isempty()
	fmt.Println(a)

}
