package main

import (
	"errors"
)

var (
	errisEmpty = errors.New("It's empty.")
	errisFull  = errors.New("It's full.")
)

func NewStack(maxSize int) stack {
	return stack{make([]int, maxSize), 0, maxSize}
}

type stack struct {
	array               []int
	currentIdx, maxSize int
}

func (s stack) isEmpty() bool {
	return s.currentIdx == 0
}

func (s stack) isFull() bool {
	return s.currentIdx == s.maxSize
}

func (s stack) Size() int {
	return len(s.array)
}

func (s *stack) Push(value int) error {
	if s.isFull() {
		return errisFull
	}
	s.array[s.currentIdx] = value
	s.currentIdx += 1
	return nil
}

func (s *stack) Pop() (int, error) {
	if s.isEmpty() {
		return 0, errisEmpty
	}
	s.currentIdx -= 1
	value := s.array[s.currentIdx]
	return value, nil
}

func (s stack) Peek() (int, error) {
	if s.isEmpty() {
		return 0, errisEmpty
	}
	return s.array[s.currentIdx-1], nil
}

func main() {
}
