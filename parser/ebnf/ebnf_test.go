package ebnf

import (
	"fmt"
	"strings"
	"testing"
)

func parseConstructionInput() []string {
	return []string{
		`111(*222*)333`,
		`111(*)222(*333`,
		`111(*222*)333(*444*)555`,
		`111(*22*)*6*)333(*444*)555`,
		"111'222'333",
		`111"222"333`,
		"'222'333",
		`"222"333`,
		//`111222333`,
	}
}

//(*
//*)
/*
`111
)222(*333`,
*/

func TestNextToken(t *testing.T) {
	data1 := `[{'abc'}]def`
	fmt.Println(nextToken(data1))
	data2 := `{'abc'}def`
	fmt.Println(nextToken(data2))
	data3 := `'abc'def`
	fmt.Println(nextToken(data3))
	data4 := `'def'`
	fmt.Println(nextToken(data4))

	for i, dt := range strings.Split("baaaab", "b") {
		fmt.Println(i, dt)
	}
}

func TestOmitComments(t *testing.T) {
	return
	fmt.Println("OMIT COMMENTS")
	for _, input := range parseConstructionInput() {
		fmt.Println("RUN:", input)
		fmt.Println(OmitComments(input), "--")
	}

}

func TestClosings(t *testing.T) {
	return
	fmt.Println("CLOSINGS")
	for _, input := range parseConstructionInput() {
		fmt.Println("RUN:", input)
		fmt.Println(TerminalString(input))
	}

}
