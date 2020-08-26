package main

// dllNode = doubly Linked List Node
// That is in order for each file under the same
// package name, to declare its own Node
type dllNode struct {
	val string
	next *dllNode
	prev *dllNode
}

type doublyLinkedList struct {
	head *dllNode
	tail *dllNode
}


func main() {

}