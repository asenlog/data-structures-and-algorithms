package main

import "fmt"

/**
 * Implement a Doubly Linked List
 */

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

// O(1) time & space
// Add a node to the tail of the list
func (dll *doublyLinkedList) append(val string) {
	if dll.head == nil {
		dll.head = &dllNode{
			val:  val,
			next: nil,
			prev: nil,
		}
		dll.tail = dll.head
		return
	}

	// The last node that will be inserted
	// prev: is the previous node which at the moment is the tail of the list
	lastNode := &dllNode{
		val:  val,
		next: nil,
		prev: dll.tail,
	}

	// before we set the new Node as the last node
	// we need to link it to the the last node
	dll.tail.next = lastNode
	// set the tail to point to the last node
	dll.tail = lastNode
	return
}

// O(n) time & space
// Add a node to the head of the list
func (dll *doublyLinkedList) prepend(val string) {
	if dll.head == nil {
		dll.head = &dllNode{
			val:  val,
			next: nil,
			prev: nil,
		}
		dll.tail = dll.head
		return
	}

	headNode := &dllNode{
		val:  val,
		next: dll.head,
		prev: nil,
	}

	dll.head.prev = headNode
	dll.head = headNode

	return
}

// O(n) time & space
// Delete the specified node of the list
func (dll *doublyLinkedList) delete(val string) {
	if dll.head == nil {
		return
	}

	// FIX: check tail and head only when the size of the list is > 2
	if dll.head.val == val {
		dll.head = dll.head.next
		dll.head.prev = nil
		return
	}

	if dll.tail.val == val {
		prePreLastNode := dll.tail.prev
		dll.tail = prePreLastNode
		prePreLastNode.next = nil
		return
	}

	node := dll.head
	for node.val != val {
		node = node.next
	}

	preNode := node.prev
	nextNode := node.next
	preNode.next = nextNode
	nextNode.prev = preNode
	return
}

// O(n) time & space
// Delete the first node with the desired data
func (dll *doublyLinkedList) print() {
	if dll.head == nil {
		return
	}

	node := dll.head
	for node != nil {
		fmt.Print(node.val + " -> ")
		node = node.next
	}
	fmt.Println()
}

func main() {
	dll := &doublyLinkedList{}
	dll.append("2")
	dll.append("3")
	dll.print()
	dll.prepend("1")
	dll.print()
	dll.delete("2")
	dll.print()
	dll.delete("1")
	dll.print()
}