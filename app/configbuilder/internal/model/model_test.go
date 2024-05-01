package model

import (
	"fmt"
	"testing"
)

func TestModel(t *testing.T) {
	m := NewModel("yaml")
	f := NewField("yaml").WithSource("Source 1").WithDataType("map[string]int").WithDesignation("IntMap").WithComment("comment").ToggleOmitempty().WithComment("this is a comment").WithValue("default", "5").WithValue("aspare", "bar")
	m.Fields = append(m.Fields, f)
	fmt.Println(f.String())

	fmt.Println(m.String())
	// m2, err := FromString(m.String())
	// if err != nil {
	// 	t.Errorf("error met: %v", err.Error())
	// }
	// fmt.Println(m)
	// fmt.Println("===")
	// fmt.Println(m2)

}
