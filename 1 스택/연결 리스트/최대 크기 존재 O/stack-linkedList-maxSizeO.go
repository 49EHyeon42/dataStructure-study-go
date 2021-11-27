package main

import (
	"errors"
)

var (
	errisEmpty = errors.New("It's empty.")
	errisFull  = errors.New("It's full.")
)

func newNode(value int, link *node) node {
	return node{value: value, link: link}
}

type node struct {
	value int
	link  *node
}

func NewStack(maxSize int) stack {
	return stack{topNode: nil, currentSize: 0, maxSize: maxSize}
}

type stack struct {
	topNode              *node
	currentSize, maxSize int
}

func (s stack) isEmpty() bool {
	return s.topNode == nil
}

func (s stack) isFull() bool {
	return s.currentSize == s.maxSize
}

func (s stack) Size() int {
	return s.currentSize
}

func (s *stack) Push(value int) error {
	if s.isFull() {
		return errisFull
	}
	newNode := newNode(value, s.topNode)
	s.topNode, s.currentSize = &newNode, s.currentSize+1
	return nil
}

func (s *stack) Pop() (int, error) {
	if s.isEmpty() {
		return 0, errisEmpty
	}
	value := s.topNode.value
	s.topNode, s.currentSize = s.topNode.link, s.currentSize-1
	return value, nil
}

func (s stack) Peek() (int, error) {
	if s.isEmpty() {
		return 0, errisEmpty
	}
	return s.topNode.value, nil
}

func main() {
}
