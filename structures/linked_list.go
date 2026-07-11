package structures

import "fmt"

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func (n *Node[T]) New(val T) *Node[T] {
	n.Value = val
	n.Next = nil
	return n
}

func (n *Node[T]) SetNext(node *Node[T]) {
	n.Next = node
}

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{Head: nil, Tail: nil}
}

func (l *LinkedList[T]) Append(val T) {
	newNode := &Node[T]{Value: val, Next: nil}
	if l.Head == nil {
		l.Head = newNode
		l.Size++
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.SetNext(newNode)
	l.Size++
}

func (l *LinkedList[T]) Prepend(val T) {
	node := &Node[T]{Value: val, Next: nil}
	node.SetNext(l.Head)
	if l.Head == nil {
		l.Tail = node
	}
	l.Head = node
	l.Size++
}

func (l *LinkedList[T]) Display() {
	current := l.Head
	for current != nil {
		fmt.Printf("%v -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}
