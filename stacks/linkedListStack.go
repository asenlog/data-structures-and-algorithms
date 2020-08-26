package main

import "fmt"

type llsNode struct {
	val string
	next *llsNode
}

// llStack = Linked List Stack
type llStack struct {
	head *llsNode
	size int
}

// O(1) time & space -  Push to the top (head of Linked List) Stack
func (lls *llStack) Push(val string) {
	if lls.head == nil {
		lls.head = &llsNode{
			val:  val,
			next: nil,
		}
		lls.size += 1
		return
	}

	newHead :=  &llsNode{
		val:  val,
		next: nil,
	}
	newHead.next = lls.head
	lls.head = newHead
	lls.size += 1
	return
}

// O(1) time & space - Pop from the Stack (pop from the head of Linked List)
func (lls *llStack) Pop() (string, error){
	if lls.head == nil {
		return "", fmt.Errorf("stack is empty")
	}

	val := lls.head.val
	lls.head = lls.head.next
	lls.size -= 1

	return val, nil
}

// Test the Data Structure by
// building and running the file.
func main() {
	lls := llStack{}
	lls.Push("1")
	lls.Push("2")
	lls.Push("3")
	val, _ := lls.Pop()
	fmt.Printf("Popped value is: %s\n", val)
	val, _ = lls.Pop()
	fmt.Printf("Popped value is: %s\n", val)
}
