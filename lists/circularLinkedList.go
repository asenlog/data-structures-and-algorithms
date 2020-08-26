package main

// cllNode = circular Linked List Node
// That is in order for each file under the same
// package name, to declare its own Node
type cllNode struct {
	val string
	next *cllNode
}

type circularLinkedList struct {
	head *cllNode
	tail *cllNode
}

func main() {

}