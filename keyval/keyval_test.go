package keyval

import (
	"fmt"
	"testing"
)

func TestKV(t *testing.T) {

	kv2, err := LoadCollection("Foo")
	if err != nil {
		fmt.Println(err.Error())
		kv2, _ = NewCollection("Foo")
	}
	fmt.Println(kv2.Get("foo2"))
	kv2.Set("FOO", "BAR")
	kv2.Set("FOO1", "BAR2")
	kv2.Set("FOO3", "BAR4")
	kv2.Set("FOO5", "BAR6")
	err = SaveCollection(kv2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(kv2.Get("FOO5"))
	fmt.Println(kv2.Get("FOO7"))
	// err = DeleteCollection(kv2)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

}
