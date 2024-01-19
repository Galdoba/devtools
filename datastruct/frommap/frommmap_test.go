package frommap

import (
	"fmt"
	"testing"
)

func TestKeysSorted(t *testing.T) {
	mss := make(map[string]string)
	mss["1"] = "1"
	mss["10"] = "10"
	mss["2"] = "2"
	mss["A"] = "A"
	mss["a"] = "a"
	mss["B"] = "B"
	mss["b"] = "b"
	mss["Г"] = "Г"
	mss["г"] = "г"
	mss["Д"] = "Д"
	mss["д"] = "д"
	keys := Keys_MSS_Sorted(mss, false)
	for i, k := range keys {
		t.Errorf("position:%v key:%v\n", i, k)
	}
	keysR := Keys_MSS_Sorted(mss, true)
	for i, k := range keysR {
		t.Errorf("position:%v key:%v\n", i, k)
	}
	keysUsed := KeysUsed(keys)
	keysUsed["A"] = true
	fmt.Println(keysUsed)
}
