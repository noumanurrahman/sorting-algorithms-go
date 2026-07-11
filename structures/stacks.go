package structures

import (
	"errors"
)

type Stack[T any] struct {
	items []T
}

func (stack *Stack[T]) Push(item T) {
	stack.items = append(stack.items, item)
}
func (stack *Stack[T]) Size() int {
	return len(stack.items)
}

func (stack *Stack[T]) Pop() (T, error) {
	if len(stack.items) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	item := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return item, nil
}

func (stack *Stack[T]) Peek() (T, error) {
	if len(stack.items) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	return stack.items[len(stack.items)-1], nil
}

// IsBalanced is just an example function to demonstrate Stack
func IsBalanced(str string) bool {
	stack := new(Stack[string])
	for charInt := range str {
		char := string(str[charInt])
		if char == "(" {
			stack.Push(char)
		} else if char == ")" {
			_, err := stack.Pop()

			if err != nil {
				return false
			}
		}
	}
	if stack.Size() == 0 {
		return true
	}
	return false
}
