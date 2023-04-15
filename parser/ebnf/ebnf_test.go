package ebnf

// import (
// 	"fmt"
// 	"testing"
// )

// func parseConstructionInput() []string {
// 	return []string{
// 		`111?222?333`,
// 		`111(*222*)333`,
// 		`111(*)222(*333`,
// 		`111(*222*)333(*444*)555`,
// 		`111(*22*)*6*)333(*444*)555`,
// 		//`111222333`,
// 	}
// }

// //(*
// //*)
// /*
// `111
// )222(*333`,
// */

// func TestNextToken(t *testing.T) {
// 	/*
// 		{aaaa{bbb}ccc}ddd
// 	*/
// 	str := `?'aaa""a'b?b'cc*)cc`
// 	//fmt.Println(grabGridy(str, `"`))
// 	//op, cl := "(*", "*)"
// 	fmt.Println(catchToken(str))

// }

// func TestOmitComments(t *testing.T) {
// 	fmt.Println("OMIT COMMENTS")
// 	for _, input := range parseConstructionInput() {
// 		fmt.Print("RUN:", input)
// 		fmt.Println(" ==> ", OmitComments(input))
// 	}

// }

// func TestClosings(t *testing.T) {
// 	return
// 	fmt.Println("CLOSINGS")
// 	for _, input := range parseConstructionInput() {
// 		fmt.Println("RUN:", input)
// 		fmt.Println(TerminalString(input))
// 	}

// }
