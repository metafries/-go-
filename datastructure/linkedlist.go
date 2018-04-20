package main

import "fmt"

type SLLNode struct {
	next  *SLLNode
	value int
}

func (node *SLLNode) SetValue(v int) {
	node.value = v
}

func (node *SLLNode) GetValue() int {
	return node.value
}

func NewSLLNode() *SLLNode {
	return new(SLLNode)
}

type SinglyLinkedList struct {
	head *SLLNode
	tail *SLLNode
}

func newSinglyLinkedList() *SinglyLinkedList {
	return new(SinglyLinkedList)
}

func (list *SinglyLinkedList) Add(v int) {
	newNode := &SLLNode{value: v}
	if list.head == nil {
		list.head = newNode
	} else if list.tail == list.head {
		list.head.next = newNode
	} else if list.tail != nil {
		list.tail.next = newNode
	}
	list.tail = newNode
}

func (list *SinglyLinkedList) String() string {
	var str string
	for node := list.head; node != nil; node = node.next {
		str += fmt.Sprintf(" {%d} ", node.GetValue())
	}
	return str
}

func main() {
	list := newSinglyLinkedList()
	list.Add(12)
	list.Add(99)
	list.Add(37)
	fmt.Println("Singly Linked List", list)
}
