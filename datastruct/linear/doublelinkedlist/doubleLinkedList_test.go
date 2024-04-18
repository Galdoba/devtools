package doublelinkedlist

import (
	"fmt"
	"testing"
)

func TestAddToHead(t *testing.T) {
	ll := NewList[int]()
	ll.AddToHead(1)
	if ll.headNode.property != 1 {
		t.Errorf("expect 1, but have %v", ll.headNode.property)
	}
	ll.AddToHead(3)
	if ll.headNode.property != 3 {
		t.Errorf("expect 3 but have %v", ll.headNode.property)
	}
}

func TestIterateList(t *testing.T) {
	lst := NewList[int]()
	lst.AddToHead(1)
	lst.AddToHead(2)
	lst.AddToHead(3)
	lst.IterateList(action)

}

func action(property int) error {
	fmt.Println(property)
	return nil
}

func TestLastNode(t *testing.T) {
	fmt.Println("Last Node")
	lst := NewList[int]()
	lst.AddToHead(1)
	lst.AddToHead(5)
	lst.AddToHead(9)
	lst.AddToEnd(4)
	lst.IterateList(action)
}

func TestAddAfter(t *testing.T) {
	fmt.Println("Add After")
	lst := NewList[int]()
	lst.AddToHead(1)
	lst.AddToHead(3)
	lst.AddToEnd(5)
	lst.AddAfter(1, 7)
	n := lst.NodeBetweenValues(1, 5)
	lst.IterateList(action)
	fmt.Println(n)
}
