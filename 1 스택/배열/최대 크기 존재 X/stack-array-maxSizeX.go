package main

import (
	"errors"
)

var (
	errisEmpty = errors.New("It's empty.")
)

type stack []int

func NewStack() stack {
	return make(stack, 0)
}

func (s stack) isEmpty() bool {
	return s.Size() == 0
}

func (s stack) Size() int {
	return len(s)
}

func (s *stack) Push(value int) {
	*s = append(*s, value)
}

func (s *stack) Pop() (int, error) {
	if s.isEmpty() {
		return 0, errisEmpty
	}
	value := (*s)[s.Size()-1]
	*s = (*s)[:s.Size()-1]
	return value, nil
}

func (s stack) Peek() (int, error) {
	if s.isEmpty() {
		return 0, errisEmpty
	}
	return s[s.Size()-1], nil
}

func main() {
}
