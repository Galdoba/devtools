package keyval

import (
	"fmt"
	"os"
	"testing"
)

func TestKV(t *testing.T) {
	kv, err1 := Load("test/test", os.O_CREATE)
	if err1 != nil {
		t.Errorf("%v", err1.Error())
	}
	fmt.Println("BLANK")
	fmt.Println(kv)
	fmt.Println("ADD ANY 1")
	err1 = kv.Add("k1", "v1", false)
	if err1 != nil {
		t.Errorf("%v", err1.Error())
	}
	fmt.Println(kv)

	fmt.Println("ADD ANY 2")
	err1 = kv.Add("k1", "v1", false)
	if err1 != nil {
		t.Errorf("%v", err1.Error())
	}
	fmt.Println(kv)

	fmt.Println("ADD Unique 1")
	err1 = kv.Add("k1", "v2", true)
	if err1 != nil {
		t.Errorf("%v", err1.Error())
	}
	fmt.Println(kv)

	fmt.Println("ADD Unique 2")
	err1 = kv.Add("k1", "v2", true)
	if err1 != nil {
		t.Errorf("%v", err1.Error())
	}
	fmt.Println(kv)
	fmt.Println("Set of 2")
	err1 = kv.Set("k1", "v3", "true")
	if err1 != nil {
		t.Errorf("%v", err1.Error())
	}
	fmt.Println(kv)

	fmt.Println("Set of 3")
	err1 = kv.Set("k1", "v4", "v5", "v6")
	if err1 != nil {
		t.Errorf("%v", err1.Error())
	}
	fmt.Println(kv)

	fmt.Println("Set of 1")
	err1 = kv.Set("k1", "v7")
	if err1 != nil {
		t.Errorf("%v", err1.Error())
	}
	fmt.Println(kv)
	fmt.Println("Delete")
	err1 = Delete(kv)
	if err1 != nil {
		t.Errorf("%v", err1.Error())
	}
}
