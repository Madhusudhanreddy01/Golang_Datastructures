package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
	Prev *node
}

type doublylinked struct {
	head *node
	tail *node
}

func (l *doublylinked) push(n int) {
	newnode := &node{n, nil, nil}
	if l.head == nil {
		l.head = newnode
		l.tail = newnode
	} else {
		l.tail.next = newnode
		newnode.Prev = l.tail
		l.tail = newnode
	}

}

func (l *doublylinked) add_first(n int) {
	newnode := &node{n, nil, nil}
	if l.head == nil {
		l.head = newnode
		l.tail = newnode
	} else {
		l.head.Prev = newnode
		newnode.next = l.head
		l.head = newnode
	}
}

func (l *doublylinked) add_any(n int, pos int) {
	newnode := &node{n, nil, nil}
	p := l.head
	i := 1
	for i < pos-1 {
		p = p.next
		i = i + 1
	}
	newnode.next = p.next
	p.next.Prev = newnode
	p.next = newnode
	newnode.Prev = p

}

func (l *doublylinked) delete_first() {
	if l.head == nil {
		fmt.Println("it is a empty node")
		return
	}
	// e := l.head.data
	l.head = l.head.next
	l.head.Prev = nil
	if l.head == nil {
		l.tail = nil
	}

}

func (l *doublylinked) delete_end() {
	if l.head == nil {
		fmt.Println("it is a empty node")
		return
	}
	l.tail = l.tail.Prev
	l.tail.next = nil
}

func (l *doublylinked) delete_any(pos int) {
	p := l.head
	i := 1
	for i < pos {
		p = p.next
		i = i + 1
	}
	p.next = p.next.next
	p.next.Prev = p
}

func (l *doublylinked) display() {
	p := l.head
	for p != nil {
		fmt.Print(p.data, "-->")
		p = p.next

	}
}

func (l *doublylinked) len() int {
	p := l.head
	count := 0
	for p != nil {
		p = p.next
		count++
	}
	return count

}

func (l *doublylinked) displayprev() {
	p := l.tail
	for p != nil {
		fmt.Print(p.data, "<--")
		p = p.Prev
	}

}

func main() {
	l := doublylinked{}
	l.push(10)
	l.push(20)
	l.push(30)
	l.push(40)
	l.display()
	a := l.len()
	fmt.Println()
	fmt.Print(a)
	fmt.Println()
	l.displayprev()
	fmt.Println()
	l.add_first(8)
	l.add_first(5)
	l.display()
	b := l.len()
	fmt.Println()
	fmt.Print(b)
	fmt.Println()
	l.displayprev()
	fmt.Println()
	l.add_any(15, 4)
	l.add_any(25, 6)
	l.display()
	c := l.len()
	fmt.Println()
	fmt.Print(c)
	fmt.Println()
	l.displayprev()
	fmt.Println()
	l.delete_first()
	l.display()
	d := l.len()
	fmt.Println()
	fmt.Print(d)
	fmt.Println()
	l.displayprev()
	fmt.Println()
	l.delete_end()
	l.display()
	e := l.len()
	fmt.Println()
	fmt.Print(e)
	fmt.Println()
	l.displayprev()
	fmt.Println()
	l.delete_any(4)
	l.display()
	f := l.len()
	fmt.Println()
	fmt.Print(f)
	fmt.Println()
	l.displayprev()

}
