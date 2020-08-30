package main

// cllNode = circular Linked List Node
// That is in order for each file under the same
// package name, to declare its own Node
type cllNode struct {
	val string
	next *cllNode
	prev *cllNode
}

type circularLinkedList struct {
	head *cllNode
	tail *cllNode
}

func (cll *circularLinkedList) append(val string) {

}

func (cll *circularLinkedList) prepend(val string) {

}

func (cll *circularLinkedList) delete(val string) {

}

// O(n) time | O(1) space
// Check and return if the list is circular
func (cll *circularLinkedList) isCircular() bool {
	if cll.head == nil {
		return false
	}

	// slow pointer
	sp := cll.head
	// fast pointer
	fp := cll.head

	for fp != nil && fp.next != nil {
		sp = sp.next
		fp = fp.next.next

		if sp == fp {
			return true
		}
	}

	// If we get to a node where fast doesn't have a next node or doesn't exist itself,
	// the list has an end and isn't circular
	return false
}

func main() {
	cll := circularLinkedList{}
	cll.append("1")
	cll.append("2")
	cll.append("3")
	cll.isCircular()

}