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
	stack.Push(1)
	stack.Push("P")
	if stack.Len() != 3 {
		t.Errorf("Push() method failed. Stack length should be 2.")
	}

	// test pop up last element
	e := stack.Pop()
	if e.(string) != "P" {
		t.Errorf("Pop() method failed. Element should be \"P\"")
	}

	e = stack.Pop()
	if e.(int) != 1 {
		t.Errorf("Pop() method failed. It should be 1.")
	}

	if stack.Len() != 1 {
		t.Errorf("Pop() method failed. Length of stack should 1 now.")
	}

	// test Top() method
	e = stack.Top()
	if e.(float64) != 3.1415 {
		t.Errorf("Top() method failed. It should be 3.1415")
	}

	if stack.Len() != 0 {
		t.Errorf("Stack length should be 0.")
	}
}
