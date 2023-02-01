package ebnf

import (
	"fmt"
	"strings"
)

/*
aaa[bb]aaa[bb]aaa
aaa aaa[bb]aaa
aaaaaa[bb]aaa
aaaaaa aaa
aaaaaaaaa
aaaaaaaaa

*/

func OmitComments(data string) string {
	dataNew := debracketGreedy(data, "(*", "*)")
	for dataNew != data {
		data = dataNew
		dataNew = debracketGreedy(data, "(*", "*)")
	}
	return data
}

func MapTerminalStrings(data string) map[int]string {
	tsm := make(map[int]string)
	currentTS := -1
	feed := strings.Split(data, "")
	for i, f := range feed {
		if 
		tail := strings.Join(feed[i:])
	}
	bod, pre, suf := trimClosingsGreedy(data, `"`, `"`)

	return tsm
}

func debracketGreedy(str, open, close string) string {
	//111(*222*)333
	feed := strings.Split(str, "")
	opened := false
	closed := false
	result := ""
	buf := ""
	for _, f := range feed {
		//fmt.Println(result)
		result += f
		if closed {
			continue
		}
		buf += f
		if !opened && strings.HasSuffix(result, open) {
			opened = true
			//	fmt.Println("open")
			buf = open + ".."
			result += ".."
			continue
		}
		if opened && !closed && strings.HasSuffix(result, close) {
			//	fmt.Println("close")
			closed = true
		}
	}
	result = strings.Join(strings.Split(result, buf), "")
	//fmt.Println("END:", result)
	if !closed {
		//	fmt.Println("STR")
		return str
	}
	return result
}

func TerminalString(data string) string {
	str, _, _ := trimClosingsGreedy(data, "'", "'")
	if str != data {
		return str
	}
	str, _, _ = trimClosingsGreedy(data, `"`, `"`)
	if str != data {
		return str
	}
	return ""
}

func haveTerminalString(data string) bool {
	quotes := []string{"'", `"`}
	for _, q := range quotes {
		parts := strings.Split(data, q)
		if len(parts) > 2 {
			return true
		}
	}
	return false
}

/*
Usage					Notation	name	priority
----------------		---------	------	----
definition				=			LIT		4
concatenation			,			JOIN	5
termination				;
alternation				|			OR		3
optional				[ ... ]		OPT		1
repetition				{ ... }		WHILE	2
grouping				( ... )
terminal string			" ... "
terminal string			' ... '
comment					(* ... *)
special sequence		? ... ?
exception				-			NOT		0
/////////////////////////////
любая строка для обработки должна начинаться с одного из Notation
если оно не найдено - ошибка некоректный синтаксис или глиф не определен
*/
func nextToken(data string) (token string, RHS string, rest string, err error) {
	//token, RHS, rest := string, string, string
	data = OmitComments(data)
	switch {
	default:
		token = "UNRECONIZED"
		RHS = data
	case isDefinition(data):
		token = "LIT"
		//RHS, rest, err = grabDefinition(data)
	case isTerminalString1(data):
		token = "terminal_string_1"
		RHS, rest, err = grabTS1(data)
	case isTerminalString2(data):
		token = "terminal_string_2"
		RHS, rest, err = grabTS2(data)
	}
	return
}

func grabTS1(data string) (RHS string, rest string, err error) {
	dt := strings.Split(data, "'")
	if len(dt) < 2 {
		err = fmt.Errorf("TS1 is not closed: '%v'", data)
		RHS = data
	}
	RHS = dt[1]
	rest = strings.Join(dt[2:], "'")
	return
}

func grabTS2(data string) (RHS string, rest string, err error) {
	dt := strings.Split(data, `"`)
	if len(dt) < 2 {
		err = fmt.Errorf("TS2 is not closed: '%v'", data)
		RHS = data
	}
	RHS = dt[1]
	rest = strings.Join(dt[2:], `"`)
	return
}

//DEFINITIONS:
func isTerminalString1(data string) bool {
	return strings.HasPrefix(data, "'")
}
func isTerminalString2(data string) bool {
	return strings.HasPrefix(data, `"`)
}
func isDefinition(data string) bool {
	return strings.HasPrefix(data, "=")
}
func isConcatenation(data string) bool {
	return strings.HasPrefix(data, ",")
}
func isTermination(data string) bool {
	return strings.HasPrefix(data, ";")
}
func isAlternation(data string) bool {
	return strings.HasPrefix(data, "|")
}
func isOptional(data string) bool {
	return strings.HasPrefix(data, "[")
}
func isRepetition(data string) bool {
	return strings.HasPrefix(data, "{")
}
func isGrouping(data string) bool {
	return strings.HasPrefix(data, "(")
}
func isComment(data string) bool {
	return strings.HasPrefix(data, "(*")
}
func isSpecialSequance(data string) bool {
	return strings.HasPrefix(data, "?")
}
func isException(data string) bool {
	return strings.HasPrefix(data, "-")
}

//(token, lhs, rhs, tail string, error)
// func nextToken(data string) (string, string, string, string, error) {
// 	feed := strings.Split(data, "")

// 	for _, f := range feed {
// 		switch f {
// 		case "'": //termStr
// 			bod, pre, suf := trimClosingsGreedy(data, "'", "'")
// 			return "TS1", pre, bod, suf, nil
// 		case `"`:
// 			bod, pre, suf := trimClosingsGreedy(data, `"`, `"`)
// 			return "TS2", pre, bod, suf, nil
// 		case "[":
// 			bod, pre, suf := trimClosingsGreedy(data, "[", "]")
// 			return "OPT", pre, bod, suf, nil
// 		case "{":
// 			bod, pre, suf := trimClosingsGreedy(data, "{", "}")
// 			return "WHILE", pre, bod, suf, nil
// 		case "(":
// 			bod, pre, suf := trimClosingsGreedy(data, "(", ")")
// 			return "grouping", pre, bod, suf, nil
// 		case "?":
// 			bod, pre, suf := trimClosingsGreedy(data, "?", "?")
// 			return "special_seq", pre, bod, suf, nil
// 		}
// 	}
// 	return "", "", "", "", fmt.Errorf("can't find token")
// }

/*

[{'abc'}]"de"

*/

type tokenList struct {
	t []token
}

type token struct {
	name string
	rhs  string
}

//func parseTokens(data string) tokenList {

//}

func terminated(rbnf string) bool {
	return strings.HasSuffix(rbnf, ";")
}

func haveClosings(rbnf string) int {
	if strings.Contains(rbnf, "[") && strings.Contains(rbnf, "]") {
		return 1
	}

	return 0
}

////////////////////////////////
func trimClosingsLazy(str, open, close string) (string, string, string) {
	if !strings.Contains(str, open) || !strings.Contains(str, close) {
		return str, "", ""
	}
	pre, str1 := trimPrefixBefore(str, open)
	str2, suf := trimPrefixBefore(str1, close)
	return str2, pre, suf
}

func trimClosingsGreedy(str, open, close string) (string, string, string) {
	if !strings.Contains(str, open) || !strings.Contains(str, close) {
		return str, "", ""
	}
	pre, str1 := trimPrefixBefore(str, open)
	suf, str2 := trimSuffixBefore(str1, close)
	return str2, pre, suf
}

func trimPrefixBefore(str, trig string) (string, string) {
	arr1 := strings.Split(str, trig)
	return arr1[0], strings.Join(arr1[1:], trig)
}

func trimSuffixBefore(str, trig string) (string, string) {
	arr := strings.Split(str, trig)
	return arr[len(arr)-1], strings.Join(arr[:len(arr)-1], trig)
}
