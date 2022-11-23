package main

import "fmt"

type node struct {
	data int
	next *node
}

type linklist struct {
	head *node
	tail *node
}

func (l *linklist) push(n int) {
	newnode := &node{n, nil}
	if l.head == nil {
		l.head = newnode
	} else {
		l.tail.next = newnode
	}
	l.tail = newnode
}

func (l *linklist) display() {
	p := l.head
	for p != nil {
		fmt.Print(p.data, "-->")
		p = p.next
	}
}

func (l *linklist) len() int {
	p := l.head
	count := 0
	for p != nil {
		p = p.next
		count++
	}
	return count

}

func (l *linklist) add_first(n int) {
	newnode := &node{n, nil}
	if l.head == nil {
		l.head = newnode
		l.tail = newnode
	} else {
		newnode.next = l.head
		l.head = newnode
	}
}

func (l *linklist) add_last(n int) {
	newnode := &node{n, nil}
	if l.head == nil {
		l.head = newnode
		l.tail = newnode
	} else {

		l.tail.next = newnode
	}
}

func (l *linklist) add_any(n int, pos int) {
	newnode := &node{n, nil}
	p := l.head
	i := 0
	for i < pos-1 {
		p = p.next
		i++
	}
	newnode.next = p.next
	p.next = newnode
}

func (l *linklist) delete_first() {
	if l.head == nil {
		fmt.Print("it is empty")
		return
	}
	//l.head.data = l.head.data
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
		return
	}

}

func (l *linklist) delete_end() {
	if l.head == nil {
		fmt.Print("it is empty")
		return
	}
	p := l.head
	i := 1
	for i < l.len()-1 {
		p = p.next
		i++
	}
	l.tail = p
	p = p.next
	//l.head.data = l.head.data
	l.tail.next = nil
	return

}

func (l *linklist) delete_any(pos int) {
	if l.head == nil {
		fmt.Print("it is empty")
		return
	}
	p := l.head
	i := 1
	for i < pos {
		p = p.next
		i++
	}
	p.data = p.next.data
	p.next = p.next.next
	return

}

func main() {
	l := linklist{}
	l.push(10)
	l.push(20)
	l.push(30)
	l.push(40)
	l.display()
	a := l.len()
	fmt.Println()
	fmt.Print(a)
	fmt.Println()
	l.add_first(8)
	l.add_first(5)
	l.display()
	b := l.len()
	fmt.Println()
	fmt.Print(b)
	fmt.Println()
	l.add_last(55)
	l.display()
	c := l.len()
	fmt.Println()
	fmt.Print(c)
	fmt.Println()
	l.add_any(25, 4)
	l.display()
	d := l.len()
	fmt.Println()
	fmt.Print(d)
	fmt.Println()
	l.delete_first()
	l.display()
	e := l.len()
	fmt.Println()
	fmt.Print(e)
	fmt.Println()
	l.delete_end()
	l.display()
	f := l.len()
	fmt.Println()
	fmt.Print(f)
	fmt.Println()
	l.delete_any(5)
	l.display()
	g := l.len()
	fmt.Println()
	fmt.Print(g)

}
