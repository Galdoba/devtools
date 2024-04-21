package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	st := NewStack[string]()
	es1 := NewElement("str1")
	es2 := NewElement("str2")
	es3 := NewElement("str3")
	// ei4 := NewElement[int]()
	st.Push(es1)
	st.Push(es2)
	st.Push(es3)
	fmt.Println(st.Pop())
	fmt.Println(st)
}
