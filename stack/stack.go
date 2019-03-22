/*
 * Array based generics stack.
 */
package stack

type (
	Stack []interface{}
)

// Instantiation of stack
func New() *Stack {
	return &Stack{}
}

// return length of stack
func (this *Stack) Len() int {
	return len(*this)
}

// Push an element to stack
func (this *Stack) Push(ele interface{}) {
	*this = append(*this, ele)
}

// Pop up last element
func (this *Stack) Pop() interface{} {
	element := (*this)[1]
	*this = (*this)[1:]
	return element
}
