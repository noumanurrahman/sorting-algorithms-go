package examples

import "sorting-algorithms/structures"

func LinkedList() {
	ll := structures.NewLinkedList[string]()
	ll.Append("Hello!")
	ll.Prepend("Hey!")
	ll.Append("How're you?!")
	ll.Display()
	ll.Prepend("Hi!")
	ll.Display()
}
