// package main

// import "fmt"

// type node struct {
// 	data int
// 	next *node
// }

// type linklist struct {
// 	head *node
// 	tail *node
// }

// func (l *linklist) push(n int) {
// 	newnode := &node{n, nil}
// 	if l.head == nil {
// 		l.head = newnode
// 	} else {
// 		l.tail.next = newnode
// 	}
// 	l.tail = newnode
// }

// func (l *linklist) display() {
// 	p := l.head
// 	for p != nil {
// 		fmt.Print(p.data, "-->")
// 		p = p.next
// 	}
// }

// func (l *linklist) len() int {
// 	p := l.head
// 	count := 0
// 	for p != nil {
// 		p = p.next
// 		count++
// 	}
// 	return count

// }

// func (l *linklist) add_first(n int) {
// 	newnode := &node{n, nil}
// 	if l.head == nil {
// 		l.head = newnode
// 		l.tail = newnode
// 	} else {
// 		newnode.next = l.head
// 		l.head = newnode
// 	}
// }

// func (l *linklist) add_last(n int) {
// 	newnode := &node{n, nil}
// 	if l.head == nil {
// 		l.head = newnode
// 		l.tail = newnode
// 	} else {

// 		l.tail.next = newnode
// 	}
// }

// func (l *linklist) add_any(n int, pos int) {
// 	newnode := &node{n, nil}
// 	p := l.head
// 	i := 0
// 	for i < pos-1 {
// 		p = p.next
// 		i++
// 	}
// 	newnode.next = p.next
// 	p.next = newnode
// }

// func (l *linklist) delete_first() {
// 	if l.head == nil {
// 		fmt.Print("it is empty")
// 		return
// 	}
// 	//l.head.data = l.head.data
// 	l.head = l.head.next
// 	if l.head == nil {
// 		l.tail = nil
// 		return
// 	}

// }

// func (l *linklist) delete_end() {
// 	if l.head == nil {
// 		fmt.Print("it is empty")
// 		return
// 	}
// 	p := l.head
// 	i := 1
// 	for i < l.len()-1 {
// 		p = p.next
// 		i++
// 	}
// 	l.tail = p
// 	p = p.next
// 	//l.head.data = l.head.data
// 	l.tail.next = nil
// 	return

// }

// func (l *linklist) delete_any(pos int) {
// 	if l.head == nil {
// 		fmt.Print("it is empty")
// 		return
// 	}
// 	p := l.head
// 	i := 1
// 	for i < pos {
// 		p = p.next
// 		i++
// 	}
// 	p.data = p.next.data
// 	p.next = p.next.next
// 	return

// }

// func main() {
// 	l := linklist{}
// 	l.push(10)
// 	l.push(20)
// 	l.push(30)
// 	l.push(40)
// 	l.display()
// 	a := l.len()
// 	fmt.Println()
// 	fmt.Print(a)
// 	fmt.Println()
// 	l.add_first(8)
// 	l.add_first(5)
// 	l.display()
// 	b := l.len()
// 	fmt.Println()
// 	fmt.Print(b)
// 	fmt.Println()
// 	l.add_last(55)
// 	l.display()
// 	c := l.len()
// 	fmt.Println()
// 	fmt.Print(c)
// 	fmt.Println()
// 	l.add_any(25, 4)
// 	l.display()
// 	d := l.len()
// 	fmt.Println()
// 	fmt.Print(d)
// 	fmt.Println()
// 	l.delete_first()
// 	l.display()
// 	e := l.len()
// 	fmt.Println()
// 	fmt.Print(e)
// 	fmt.Println()
// 	l.delete_end()
// 	l.display()
// 	f := l.len()
// 	fmt.Println()
// 	fmt.Print(f)
// 	fmt.Println()
// 	l.delete_any(5)
// 	l.display()
// 	g := l.len()
// 	fmt.Println()
// 	fmt.Print(g)

// }

// -------------------------------------------------------------------------------------------------------------

package main

import "fmt"

type ListNode struct { // defines a ListNode in a linked list
	data interface{} // the datum
	next *ListNode   // pointer to the next ListNode
}

type LinkedList struct {
	head *ListNode
	size int
}

func (ll *LinkedList) insertAtBeginning(data interface{}) {
	newNode := &ListNode{
		data: data,
	}
	if ll.head == nil {
		ll.head = newNode
	} else {
		newNode.next = ll.head
		ll.head = newNode
	}
	ll.size++
	return
}

func (ll *LinkedList) insertAtEnd(data interface{}) {
	newNode := &ListNode{
		data: data,
	}
	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	ll.size++
	return
}

// Insert adds an item at position i
func (ll *LinkedList) insert(position int, data interface{}) error {
	// This condition to check whether the position given is valid or not.
	if position < 1 || position > ll.size+1 {
		return fmt.Errorf("insert: Index out of bounds")
	}
	newNode := &ListNode{data, nil}

	var prev, current *ListNode
	prev = nil
	current = ll.head

	for position > 1 {
		prev = current
		current = current.next
		position = position - 1
	}

	if prev != nil {
		prev.next = newNode
		newNode.next = current
	} else {
		newNode.next = current
		ll.head = newNode
	}

	ll.size++
	return nil
}

func (ll *LinkedList) deleteFront() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("deleteFront: List is empty")
	}
	data := ll.head.data
	ll.head = ll.head.next
	ll.size--
	return data, nil
}

func (ll *LinkedList) deleteLast() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("deleteLast: List is empty")
	}
	var prev *ListNode
	current := ll.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	if prev != nil {
		prev.next = nil
	} else {
		ll.head = nil
	}
	ll.size--
	return current.data, nil
}

// delete removes an element at position i
func (ll *LinkedList) delete(position int) (interface{}, error) {
	// This condition to check whether the position given is valid or not.
	if position < 1 || position > ll.size+1 {
		return nil, fmt.Errorf("insert: Index out of bounds")
	}
	// Complete this method
	var prev, current *ListNode
	prev = nil
	current = ll.head

	pos := 0

	if position == 1 {
		ll.head = ll.head.next
	} else {
		for pos != position-1 {
			pos = pos + 1
			prev = current
			current = current.next
		}

		if current != nil {
			prev.next = current.next
		}
	}
	ll.size--
	return current.data, nil
}

func (ll *LinkedList) display() error {
	if ll.head == nil {
		return fmt.Errorf("display: List is empty")
	}
	current := ll.head
	for current != nil {
		fmt.Printf("%v -> ", current.data)
		current = current.next
	}
	fmt.Println()
	return nil
}

// with extra size field
func (ll *LinkedList) length() int {
	return ll.size
}

// length returns the linked list size
func (ll *LinkedList) length2() int {
	size := 0
	current := ll.head
	for current != nil {
		size++
		current = current.next
	}
	return size
}

func main() {
	ll := LinkedList{}
	fmt.Printf("insertAtBeginning: A\n")
	ll.insertAtBeginning("A")
	fmt.Printf("insertAtBeginning: B\n")
	ll.insertAtBeginning("B")
	fmt.Printf("insertAtEnd: C\n")
	ll.insertAtEnd("C")
	fmt.Printf("insert: D\n")
	ll.insert(2, "D")
	fmt.Printf("length: %d\n", ll.length())

	err := ll.display()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("delete\n")
	_, err = ll.delete(4)
	if err != nil {
		fmt.Printf("deleteError: %s\n", err.Error())
	}
	fmt.Printf("length: %d\n", ll.length())
	err = ll.display()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("deleteFront\n")
	_, err = ll.deleteFront()
	if err != nil {
		fmt.Printf("deleteFront Error: %s\n", err.Error())
	}
	fmt.Printf("length: %d\n", ll.length())
	err = ll.display()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("deleteLast\n")
	_, err = ll.deleteLast()
	if err != nil {
		fmt.Printf("deleteLast Error: %s\n", err.Error())
	}
	fmt.Printf("length: %d\n", ll.length())
	err = ll.display()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("deleteLast\n")
	_, err = ll.deleteLast()
	if err != nil {
		fmt.Printf("deleteLast Error: %s\n", err.Error())
	}
	fmt.Printf("length: %d\n", ll.length())
	err = ll.display()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("deleteLast\n")
	_, err = ll.deleteLast()
	if err != nil {
		fmt.Printf("deleteLast Error: %s\n", err.Error())
	}
	fmt.Printf("length: %d\n", ll.length())
	err = ll.display()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("length: %d\n", ll.length())
}
