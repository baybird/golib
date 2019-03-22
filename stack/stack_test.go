package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	// test instantiation
	stack := New()

	if stack.Len() != 0 {
		t.Errorf("Instantiation of stack failed.")
	}

	// println(cap(*stack))

	// test push method
	stack.Push(3.1415)
	stack.Push("P")
	if stack.Len() != 2 {
		t.Errorf("Push() method failed. Stack length should be 2.")
	}

	// test pop up last element
	e := stack.Pop()
	if e.(string) != "P" || stack.Len() != 1 {
		t.Errorf("Pop() method failed.")
	}
}
