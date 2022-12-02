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

// -------------------------------------------------------------------------------------------------------

// package main

// import "fmt"

// // DLLNode ... newNode of doubly linked list
// type DLLNode struct {
// 	data int
// 	prev *DLLNode // points to previous node in link list
// 	next *DLLNode // points to next node in link list
// }

// // DLL ...
// type DLL struct {
// 	size int
// 	head *DLLNode
// 	tail *DLLNode
// }

// // NewNode ... it return a new node
// func NewNode(data int) *DLLNode {
// 	return &DLLNode{
// 		data: data,
// 		prev: nil,
// 		next: nil,
// 	}
// }

// // NewDLL ... returns a doubly linked list
// func NewDLL() DLL {
// 	return DLL{
// 		size: 0,
// 		head: nil,
// 		tail: nil,
// 	}
// }

// //CheckIfEmpty ... check if empty
// func (dll *DLL) CheckIfEmpty() bool {
// 	if dll.size == 0 {
// 		return true
// 	}
// 	return false
// }

// // CheckIfEmptyAndAdd ... check if doubly link list is empty
// func (dll *DLL) CheckIfEmptyAndAdd(newNode *DLLNode) bool {
// 	// check if list is empty
// 	if dll.size == 0 {
// 		// insert first node in doubly linked list
// 		dll.head = newNode
// 		dll.tail = newNode
// 		dll.size++
// 		return true
// 	}
// 	return false
// }

// // InsertBeginning ... insert in the beginning of doubly linked list
// func (dll *DLL) InsertBeginning(data int) {
// 	newNode := &DLLNode{
// 		data: data,
// 		prev: nil,
// 		next: nil,
// 	}
// 	if !(dll.CheckIfEmptyAndAdd(newNode)) {
// 		head := dll.head
// 		// update newnode links - prev and next
// 		newNode.next = head
// 		newNode.prev = nil

// 		// update head node
// 		head.prev = newNode

// 		// update dll start and length
// 		dll.head = newNode
// 		dll.size++
// 		return
// 	}
// 	return
// }

// // InsertEnd ... inserts a node in the end of doubly linked list
// func (dll *DLL) InsertEnd(data int) {
// 	newNode := &DLLNode{
// 		data: data,
// 		prev: nil,
// 		next: nil,
// 	}
// 	if !(dll.CheckIfEmptyAndAdd(newNode)) {
// 		head := dll.head
// 		for i := 0; i < dll.size; i++ {
// 			if head.next == nil {
// 				// update newnode links - prev and next
// 				newNode.prev = head
// 				newNode.next = nil

// 				//update head node
// 				head.next = newNode

// 				// update dll end and size
// 				dll.tail = newNode
// 				dll.size++
// 				break
// 			}
// 			head = head.next
// 		}
// 	}
// 	return
// }

// // InsertMiddle ... insert between two nodes in doubly linkedlist
// func (dll *DLL) InsertMiddle(data int, loc int) {
// 	newNode := &DLLNode{
// 		data: data,
// 		prev: nil,
// 		next: nil,
// 	}
// 	if !(dll.CheckIfEmptyAndAdd(newNode)) {
// 		head := dll.head
// 		for i := 1; i < dll.size; i++ {
// 			if i == loc {
// 				// update newnode links - prev and next
// 				newNode.prev = head.prev
// 				newNode.next = head

// 				// update head node
// 				head.prev.next = newNode
// 				head.prev = newNode

// 				//keep traversing till we find the location
// 				dll.size++
// 				return
// 			}
// 			head = head.next
// 		}
// 	}
// 	return
// }

// // DeleteFirst ... Delete last element
// func (dll *DLL) DeleteFirst() int {
// 	if !(dll.CheckIfEmpty()) {
// 		head := dll.head
// 		if head.prev == nil {
// 			deletedNode := head.data

// 			// update doubly linked list
// 			dll.head = head.next
// 			dll.head.prev = nil
// 			dll.size--

// 			return deletedNode
// 		}
// 	}
// 	return -1
// }

// // DeleteLast ... deletes last element from doubly linked list
// func (dll *DLL) DeleteLast() int {
// 	if !(dll.CheckIfEmpty()) {
// 		// delete from last
// 		head := dll.head
// 		for {
// 			if head.next == nil {
// 				break
// 			}
// 			head = head.next
// 		}

// 		// update doubly linked list
// 		dll.tail = head.prev
// 		dll.tail.next = nil
// 		dll.size--
// 		return head.data
// 	}
// 	return -1
// }

// // DeleteMiddle .. delete from any location
// func (dll *DLL) DeleteMiddle(pos int) int {
// 	if !(dll.CheckIfEmpty()) {
// 		//delete from any position
// 		head := dll.head
// 		for i := 1; i <= pos; i++ {
// 			if head.next == nil && pos > i {
// 				//list is lesser than given position
// 				return -1
// 			} else if i == pos {
// 				// delete from this location
// 				head.prev.next = head.next
// 				head.next.prev = head.prev
// 				dll.size--
// 				return head.data
// 			}
// 			head = head.next
// 		}
// 	}
// 	return -1
// }

// // Display ... Display the doubly linked list
// func (dll *DLL) Display() {
// 	head := dll.head
// 	for i := 0; i < dll.size; i++ {
// 		fmt.Println("-----newNode no------", i+1)
// 		fmt.Println("data", head.data)
// 		fmt.Println("prev", head.prev)
// 		fmt.Println("next", head.next)
// 		fmt.Println("------------------")
// 		head = head.next
// 	}
// }

// func main() {
// 	dll := DLL{}
// 	dll.InsertBeginning(100)
// 	dll.InsertBeginning(90)
// 	dll.InsertBeginning(70)
// 	dll.InsertBeginning(60)
// 	dll.InsertBeginning(50)
// 	dll.InsertBeginning(40)
// 	dll.InsertBeginning(30)

// 	dll.InsertBeginning(19) // Update dll

// 	dll.Display()
// }
