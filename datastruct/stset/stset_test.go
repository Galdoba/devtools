package stset

import (
	"fmt"
	"testing"
)

func TestNewNonStrict(t *testing.T) {
	ss := NewStrict()
	fmt.Println(ss.Add("linux-IN", "windows-IN"))
	fmt.Println(ss.Add("linux-OUT", "windows-OUT"))
	v, b := ss.GetValue("linux-OUT")
	fmt.Println(v, b)
	v, b = ss.GetValue("windows-OUT")
	fmt.Println(v, b)
}
