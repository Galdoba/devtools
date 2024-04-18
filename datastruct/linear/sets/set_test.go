package sets

import (
	"fmt"
	"testing"
)

func TestSets(t *testing.T) {
	s1 := New[string]()
	s4 := New[string]()

	s2 := New[int]()
	s3 := New[int]()

	s1.AddElement("a")
	s1.AddElement("b")
	s1.AddElement("c")
	s1.DeleteElement("b")
	s4.AddElement("qqq")
	s4.AddElement("c")
	fmt.Println(s1)

	s2.AddElement(21)
	s2.AddElement(22)
	s2.AddElement(23)
	s2.AddElement(24)
	s3.AddElement(31)
	s3.AddElement(22)
	s3.AddElement(33)

	intersect := s4.Intersect(s1)
	union := s4.Union(s1)
	fmt.Println(intersect)
	fmt.Println(union)

}
