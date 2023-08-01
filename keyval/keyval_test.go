package keyval

import (
	"fmt"
	"testing"
)

func TestKV(t *testing.T) {
	kv, er := New("Foo")
	fmt.Println(kv)
	fmt.Println(er)
}
