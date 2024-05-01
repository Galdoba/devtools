package csvp

import (
	"fmt"
	"testing"
)

func TestFromString(t *testing.T) {

	text := `1997,Ford,E350,"ac, abs, moon",3000.00` + "\n" + `1999,Chevy,"Venture ""Extended Edition""", ,4900.00` + "\n" + `1996,Jeep,Grand Cherokee,"MUST SELL! air, moon roof, loaded",4799.00`

	fmt.Println(text)

	c, err := FromString(text)

	if err != nil {
		t.Errorf("err:%v", err.Error())
	}
	fmt.Println("===")
	fmt.Println(c)
	e := NewEntry("aaa", "bbb", "ccc", "ddd", "eee", "iii", "uuu")
	c.PrependEntry(e)
	fmt.Println("===")
	fmt.Println(c)
	c.SwitchEntries(0, 1)
	fmt.Println("     +++++   ")
	fmt.Println(c)
	c.AppendColumn("111", "2+2", "444", "----")
	fmt.Println("===***")
	fmt.Println(c)
	c.SwitchColumns(0, 2)
	fmt.Println("===***")
	fmt.Println(c)
	c.AppendEntry(NewEntry("aaa", "rrr"))
	fmt.Println("===*|||")
	c.SetFieldValue(2, 2, "===========")
	c.SetFieldValue(2, 0, "===========")
	fmt.Println(c)
}
