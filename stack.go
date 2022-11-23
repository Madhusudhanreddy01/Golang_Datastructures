// package main

// import "fmt"

// type node struct {
// 	data int
// 	next *node
// }

// type stacklist struct {
// 	head *node
// }

// func (s *stacklist) push(n int) {
// 	newnode := &node{n, nil}
// 	if s.head == nil {
// 		s.head = newnode
// 	} else {
// 		newnode.next = s.head
// 		s.head = newnode
// 	}
// }

// func (s *stacklist) display() {
// 	p := s.head
// 	for p != nil {
// 		fmt.Print(p.data, "-->")
// 		p = p.next
// 	}
// }

// func (s *stacklist) pop() {
// 	if s.head == nil {
// 		fmt.Print("it is a empty stack")
// 		return
// 	}
// 	// s.head.data = s.head.data
// 	s.head = s.head.next
// }

// func (s *stacklist) removedelement() {
// 	if s.head == nil {
// 		fmt.Print("it is a empty stack")
// 		return
// 	}
// 	fmt.Print(s.head.data)
// }

// func (s *stacklist) len() int {
// 	p := s.head
// 	count := 0
// 	for p != nil {
// 		p = p.next
// 		count++
// 	}
// 	return count
// }

// func main() {
// 	l := stacklist{}
// 	l.push(10)
// 	l.push(20)
// 	l.push(30)
// 	l.display()
// 	fmt.Println()
// 	fmt.Println("removed element")
// 	l.removedelement()
// 	fmt.Println()

// 	l.pop()
// 	l.display()
// 	fmt.Println()

// 	l.push(30)
// 	l.push(40)
// 	l.display()
// 	fmt.Println()

// 	l.removedelement()

// 	a := l.len()
// 	fmt.Println()
// 	fmt.Print(a)

// }

package main

import (
	"fmt"
)

type Stack struct {
	data []int
}

func (S Stack) len() int {
	return len(S.data)
}

func (S Stack) isEmpty() bool {
	return S.len() == 0
}

func (S *Stack) push(element int) {
	S.data = append(S.data, element)
}

func (S *Stack) pop() int {
	if S.isEmpty() {
		fmt.Println("----Stack is Empty -----")
		return -1
	}

	res := S.top()

	S.data = S.data[0 : S.len()-1]

	return res
}

func (S Stack) top() int {
	if S.isEmpty() {
		fmt.Println("----Stack is Empty -----")
		return -1
	}
	return S.data[S.len()-1]
}

func main() {
	fmt.Println("--- Slice Stack---")
	s := Stack{}
	fmt.Println(s.len())
	s.push(10)
	s.push(20)
	s.push(30)
	fmt.Println(s.data)
	fmt.Println(s.len())
	fmt.Println(s.top())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.data)
	fmt.Println(s.len())
	fmt.Println(s.top())
}
