package ebnf

import (
	"fmt"
	"testing"
)

func parseConstructionInput() []string {
	return []string{
		`111(*)222(*)333`,
		`111(*)222(*333`,

		`111(*222*)333(*444*)555`,
		//"111'222'333",
		//`111"222"333`,
		//`111222333`,
	}
}

/*
`111
)222(*333`,
*/

func TestOmitComments(t *testing.T) {
	fmt.Println("OMIT COMMENTS")
	for _, input := range parseConstructionInput() {
		fmt.Println("RUN:", input)
		fmt.Println(OmitComments(input), "--")
	}

}

func TestParseConstruction(t *testing.T) {
	return
	for _, input := range parseConstructionInput() {
		fmt.Println("RUN:", input)
		fmt.Println(nextToken(input))
	}

}
