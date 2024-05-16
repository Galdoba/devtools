package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stk := NewStack[int]()
	stk.Push(7)
	stk.Push(5)
	stk.Push(3)
	stk.Push(11)
	fmt.Println(stk)
	fmt.Println(stk.Pop())
	fmt.Println(stk.Pop())
	fmt.Println(stk.Pop())
	fmt.Println(stk.Pop())
	fmt.Println(stk.Pop())
}
