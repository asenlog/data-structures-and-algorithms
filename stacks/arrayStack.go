package main

import "fmt"

/**
 * Implement a stacks using a Slice (an array if you will)
 */

type stack []string

// O(1) time
func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

// O(1) time
func (s *stack) Push(val string) {
	*s = append(*s, val)
}

// O(1) time
func (s *stack) Pop() (string, error) {
	// If stacks is empty just return error
	if s.IsEmpty() {   	// or if len(*s) == 0 {}
		return "", fmt.Errorf("empty stacks")
	}

	// Get the index of the top most element.
	index := len(*s) - 1
	// Index into the slice and obtain the element.
	element := (*s)[index]
	// Remove it from the stacks by slicing it off.
	*s = (*s)[:index]

	return element, nil
}


// Test the Data Structure by
// building and running the file.
func main() {
	st := stack{}
	st.Push("1")
	st.Push("2")
	st.Push("3")
	fmt.Printf("%v\n", st)
	val, _ := st.Pop()
	fmt.Printf("Popped value is: %s\n", val)
	fmt.Printf("%v", st)
}



