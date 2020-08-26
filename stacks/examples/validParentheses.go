/**
LeetCode: https://leetcode.com/problems/valid-parentheses/
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.

Note that an empty string is also considered valid.

Example 1:
Input: "()"
Output: true

Example 2:
Input: "()[]{}"
Output: true

Example 3:
Input: "(]"
Output: false
 */

package examples

import "fmt"

/**
* Implement a stacks using a Slice (an Array if you will)
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
	if s.IsEmpty() {
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

func getReversedBracket(s string) string {
	switch s {
	case ")":
		return "("
	case "}":
		return "{"
	case "]":
		return "["
	default:
		return ""
	}
}

func isLeftBracket(s string) bool {
	switch s {
	case "(", "{", "[":
		fmt.Println("is left")
		return true
	default:
		return false
	}
}

func isValid(s string) bool {
	st := stack{}
	for _, c := range s {
		if ok := isLeftBracket(string(c)); ok {
			fmt.Println("in push")
			st.Push(string(c))
			continue;
		}

		// get the reversed
		rev := getReversedBracket(string(c))
		val, err := st.Pop()
		// we also check here if the stacks is empty
		if err != nil {
			return false
		}

		if rev == "" || rev != val {
			return false
		}
	}

	return st.IsEmpty()
}
