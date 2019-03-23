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
	length := len(*this)
	if length > 0 {
		lastIndex := length - 1
		element := (*this)[lastIndex]
		*this = (*this)[0:lastIndex]
		return element
	}

	return nil
}

// Top - return the first element
func (this *Stack) Top() interface{} {
	length := len(*this)
	if length > 0 {
		element := (*this)[0]
		*this = (*this)[1:]
		return element
	}

	return nil
}
