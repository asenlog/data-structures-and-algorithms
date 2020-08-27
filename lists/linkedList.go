package main

import "fmt"

/**
 * Implement a Singly Linked List without a tail
 */

// llNode = Linked List Node
// That is in order for each file under the same
// package name, to declare its own Node
type llNode struct {
	val string
	next *llNode
}

// linkedList the Linked List
type linkedList struct {
	head *llNode
	tail *llNode
}

// O(n) time & space || If we were to use a tail we would have O(1) time & space
// Add a node to the tail of the list
func (ll *linkedList) append(val string) {
	if ll.head == nil {
		ll.head = &llNode{
			val:  val,
			next: nil,
		}
		return
	}

	node := ll.head
	for node.next != nil {
		node = node.next
	}

	node.next = &llNode{
		val:  val,
		next: nil,
	}
	return
}

// O(1) time & space
// Add Node to the Head of the list
func (ll *linkedList) prepend(val string) {
	if ll.head == nil {
		ll.head = &llNode{
			val:  val,
			next: nil,
		}
		return
	}

	newHead :=  &llNode{
		val:  val,
		next: nil,
	}
	newHead.next = ll.head
	ll.head = newHead
	return
}

// O(n) time & space
// Delete the first node with the desired data
func (ll *linkedList) delete(val string) error {
	if ll.head == nil {
		return nil
	}

	// cover the case where the value
	// we need is in the Head of the list
	node := ll.head
	if node.val == val {
		node = node.next
		return nil
	}

	for node.next != nil {
		if node.val == val {
			node.next = node.next.next
			return nil
		}
		node = node.next
	}

	return fmt.Errorf("value not found in the list")
}

// O(n) time & space
// Reverse the Linked List
func (ll *linkedList) reverse() {
	if ll.head == nil {
		return
	}

	// prev is the new reversed linked list that will be created.
	var prev *llNode
	current := ll.head
	for current != nil {
		next := current.next
		// Next 2 to lines do the trick
		// of swapping the nodes
		current.next = prev
		prev = current

		current = next
	}
	ll.head = prev
	return
}

// O(n) time & space
// Delete the first node with the desired data
func (ll *linkedList) print() {
	if ll.head == nil {
		return
	}

	node := ll.head
	for node != nil {
		fmt.Print(node.val + " -> ")
		node = node.next
	}
	fmt.Println()
}

// Test the Data Structure by
// building and running the file.
func main() {
	ll := &linkedList{}
	ll.append("1")
	ll.append("2")
	ll.append("3")
	ll.print()
	ll.reverse()
	ll.print()
}