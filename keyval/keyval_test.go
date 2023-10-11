package keyval

import (
	"fmt"
	"testing"
)

func TestKV(t *testing.T) {
	kv, err1 := NewKVlist("test")
	if err1 != nil {
		t.Errorf("%v", err1.Error())
	}
	fmt.Println(kv)
	err2 := kv.Save()
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	kv, err := Load("test")
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	kv.KVpair["key1"] = []string{"val1", "val2"}

	err3 := kv.Save()
	if err3 != nil {

		t.Errorf("%v", err3.Error())
	}
	kv2, _ := Load("test")
	if _, err := kv.UpdateByIndex("key1", []int{1, -2}, "val7"); err != nil {
		fmt.Println(err.Error())
	}
	kv2.Remove("key1", "val3")
	kv2.Save()
	fmt.Println(kv)
	Delete(kv2)
}
