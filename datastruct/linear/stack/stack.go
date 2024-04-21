package stack

import (
	"fmt"
)

type Element[T comparable] struct {
	value T
}

func (e *Element[T]) String() string {
	s := fmt.Sprintf("%v", e.value)
	return s
}

func NewStack[T comparable]() *Stack[T] {
	st := &Stack[T]{}
	st.New()
	return st
}

func NewElement[T comparable](value T) *Element[T] {
	e := Element[T]{value}
	return &e
}

func (st *Stack[T]) New() {

	st.elements = make([]*Element[T], 0)
}

type Stack[T comparable] struct {
	elements     []*Element[T]
	elementCount int
}

func (st *Stack[T]) Push(e *Element[T]) {
	st.elements = append(st.elements[:st.elementCount], e)
	st.elementCount++
}

func (st *Stack[T]) Pop() *Element[T] {
	if st.elementCount == 0 {
		return nil
	}
	var lenght int = len(st.elements)
	var elem *Element[T] = st.elements[lenght-1]
	// st.elementCount = st.elementCount - 1
	if lenght > 1 {
		st.elements = st.elements[:lenght-1]
	} else {
		st.elements = st.elements[0:]
	}
	st.elementCount = len(st.elements)
	return elem
}
