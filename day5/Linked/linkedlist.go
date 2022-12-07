package Linked

import "fmt"

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (n *Node) GetData() string {
	return n.data
}

func (l *LinkedList) GetHead() *Node {
	return l.head
}

func (l *LinkedList) GetTail() *Node {
	temp2 := l.head
	for temp2.next != nil {
		temp2 = temp2.next
	}
	return temp2
}

func (l *LinkedList) Append(appendMe *LinkedList) {
	if appendMe.head == nil {
		return
	}
	if l.head == nil {
		l.head = appendMe.head
		l.length = appendMe.length
		return
	}
	temp2 := l.head
	for temp2.next != nil {
		temp2 = temp2.next
	}
	temp2.next = appendMe.head
	l.length = l.length + appendMe.length
}

func (l *LinkedList) InsertAtHead(data string) {
	temp1 := &Node{data, nil}
	if l.head == nil {
		l.head = temp1
	} else {
		temp2 := l.head
		l.head = temp1
		temp1.next = temp2
	}
	l.length += 1
}

func (l *LinkedList) InsertAtTail(data string) {
	temp1 := &Node{data, nil}
	if l.head == nil {
		l.head = temp1
	} else {
		temp2 := l.head
		for temp2.next != nil {
			temp2 = temp2.next
		}
		temp2.next = temp1
	}
	l.length += 1
}

func (l *LinkedList) Insert(n int, data string) {
	if n == 0 {
		l.InsertAtHead(data)
	} else if n == l.length-1 {
		l.InsertAtTail(data)
	} else {
		temp1 := &Node{data, nil}
		temp2 := l.head
		for i := 0; i < n-1; i++ {
			temp2 = temp2.next
		}
		temp1.next = temp2.next
		temp2.next = temp1
	}
	l.length += 1
}

func (l *LinkedList) DeleteAtHead() {
	if l.length == 0 {
		return
	}
	temp := l.head
	l.head = temp.next
	l.length -= 1
}

func (l *LinkedList) DeleteAtTail() {
	if l.length == 0 {
		return
	}
	if l.head.next == nil {
		l.head = nil
		l.length -= 1
		return
	}
	temp1 := l.head
	var temp2 *Node
	for temp1.next != nil {
		temp2 = temp1
		temp1 = temp1.next
	}
	temp2.next = nil
	l.length -= 1
}

func (l *LinkedList) Delete(n int) {
	if n == 0 {
		l.DeleteAtHead()
	} else if n == l.length-1 {
		l.DeleteAtTail()
	} else {
		temp1 := l.head
		for i := 0; i < n-1; i++ {
			temp1 = temp1.next
		}
		temp2 := temp1.next
		temp1.next = temp2.next
	}
	l.length -= 1
}

func (l *LinkedList) Reverse() {
	var curr, prev, next *Node
	curr = l.head
	prev = nil
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}

func (l *LinkedList) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.data)
		list = list.next
	}
	fmt.Println()
}

/*
type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Insert(key interface{}) {
	list := &Node{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l
}

func (L *List) GetTail() *Node {
	return L.tail
}

func (L *List) RemoveTail() {
	if L.tail == L.head {
		L.tail = nil
		L.head = nil
	} else {
		L.tail.prev.next = nil
		L.tail = L.tail.prev
	}
}
func (L *List) RemoveHead() {
	if L.tail == L.head {
		L.tail = nil
		L.head = nil
	} else {
		L.head.next = nil
		L.tail = L.tail.prev
	}
}
func (L *List) GetHead() *Node {
	return L.head
}

func (L *List) GetNext(list *Node) *Node {
	return list.next
}

func (L *List) GetPrevious(list *Node) *Node {
	return list.prev
}

func (N *Node) GetKey() interface{} {
	return N.key
}

func (L *List) MoveTo(N, after *Node) {
	N.prev.next = N.next
	N.next = after.next
	N.prev = after
	after.next = N
}

func (L *List) Delete(n *Node) {
	if L.head == n {
		L.head = L.head.next
		return
	}
	var prev *Node = nil
	curr := L.head
	for curr != nil && curr != n {
		prev = curr
		curr = curr.next
	}
	if curr == nil {
		return
	}
	prev.next = curr.next
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func Display(list *Node) {
	for list != nil {
		fmt.Printf("%v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func ShowBackwards(list *Node) {
	for list != nil {
		fmt.Printf("%v <-", list.key)
		list = list.prev
	}
	fmt.Println()
}

func (l *List) Reverse() {
	curr := l.head
	var prev *Node
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
	Display(l.head)
}
*/
