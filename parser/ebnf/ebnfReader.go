package ebnf

// import (
// 	"fmt"
// 	"strings"
// )

// /*
// aaa[bb]aaa[bb]aaa
// aaa aaa[bb]aaa
// aaaaaa[bb]aaa
// aaaaaa aaa
// aaaaaaaaa
// aaaaaaaaa

// */

// func OmitComments(data string) string {
// 	dataNew := debracketGreedy(data, "(*", "*)")
// 	for dataNew != data {
// 		data = dataNew
// 		dataNew = debracketGreedy(data, "(*", "*)")
// 	}
// 	return data
// }

// func debracketGreedy(str, open, close string) string {
// 	//111(*222*)333
// 	feed := strings.Split(str, "")
// 	opened := false
// 	closed := false
// 	result := ""
// 	buf := ""
// 	for _, f := range feed {
// 		//fmt.Println(result)
// 		result += f
// 		if closed {
// 			continue
// 		}
// 		buf += f
// 		if !opened && strings.HasSuffix(result, open) {
// 			opened = true
// 			//	fmt.Println("open")
// 			buf = open + ".."
// 			result += ".."
// 			continue
// 		}
// 		if opened && !closed && strings.HasSuffix(result, close) {
// 			//	fmt.Println("close")
// 			closed = true
// 		}
// 	}
// 	result = strings.Join(strings.Split(result, buf), "")
// 	//fmt.Println("END:", result)
// 	if !closed {
// 		//	fmt.Println("STR")
// 		return str
// 	}
// 	return result
// }

// func TerminalString(data string) string {
// 	str, _, _ := trimClosingsGreedy(data, "'", "'")
// 	if str != data {
// 		return str
// 	}
// 	str, _, _ = trimClosingsGreedy(data, `"`, `"`)
// 	if str != data {
// 		return str
// 	}
// 	return ""
// }

// func haveTerminalString(data string) bool {
// 	quotes := []string{"'", `"`}
// 	for _, q := range quotes {
// 		parts := strings.Split(data, q)
// 		if len(parts) > 2 {
// 			return true
// 		}
// 	}
// 	return false
// }

// /*
// Usage					Notation	name	priority
// ----------------		---------	------	----
// 0   definition				=			LIT		4
// 1   concatenation			,			JOIN	5
// 2   termination				;
// 3   alternation				|			OR		3
// 4   optional				[ ... ]		OPT		1
// 5   repetition				{ ... }		WHILE	2
// 6   grouping				( ... )
// 7   terminal string			" ... "
// 8   terminal string			' ... '
// 9   comment					(* ... *)	_OMIT_
// 10  special sequence		? ... ? 	_OMIT_
// 11  exception				-			NOT		0
// /////////////////////////////
// любая строка для обработки должна начинаться с одного из Notation
// если оно не найдено - ошибка некоректный синтаксис или глиф не определен
// 0 ищем определение ( = ) 																	возвращаем Trigger, LHS, RHS
// 1 ищем следующий простой триггер ( = , ; |) 												возвращаем Trigger, RHS, rest
// 2 ищем следующий закрытый триггер ( [ .. ] { .. } ( .. ) " .. " ' .. ' (* .. *) ? .. ? ) 	возвращаем Trigger, RHS, rest
// 3 ищем отрицание ( - ) 																		возвращаем Trigger, RHS, rest
// */

// type token struct {
// 	ttype int
// 	RHS   string
// 	LHS   string
// 	tail  string
// 	err   error
// }

// const (
// 	undefined       = iota
// 	definition      //  =
// 	concatenation   //  ,
// 	termination     //  ;
// 	alternation     //  |
// 	optional        //  [ ... ]
// 	repetition      //  { ... }
// 	grouping        //  ( ... )
// 	terminalstring1 //  " ... "
// 	terminalstring2 //  ' ... '
// 	comment         //  (* ... *)	- omit or ERROR?
// 	specialsequence //  ? ... ?		- omit or ERROR?
// 	exception       //  -
// )

// var invalidSequances = []string{"(*)", "(:)", "(/)"}

// func haveInvalidSeq(data string) string {
// 	for _, val := range invalidSequances {
// 		if strings.Contains(data, val) {
// 			return val
// 		}
// 	}
// 	return ""
// }

// func catchToken2(data string) token {
// 	feed := strings.Split(data, "")
// 	t := token{}
// 	t.ttype = undefined
// 	buf := ""
// 	state := stateBasic
// 	for i, f := range feed {
// 		buf += f
// 		switch {
// 		default:
// 		case strings.HasSuffix(buf, ","):
// 			switch state {
// 			case stateTerminal, stateComment, stateSpecial:
// 				continue
// 			}
// 			t.ttype = concatenation
// 			t.LHS = strings.TrimSuffix(buf, ",")
// 			t.RHS = restAsString(feed, i)
// 			return t
// 		case strings.HasSuffix(buf, "="):
// 		}
// 	}
// 	return t
// }

// func restAsString(sl []string, i int) string {
// 	if len(sl) <= i {
// 		return ""
// 	}
// 	return strings.Join(sl[:i], "")
// }

// var tokenOrder = []string{",", "=", "|", "*)", ")", "]", "}", "-", "`", "*", `"`, "?", "(*", "(", "[", "{", ";"}

// func catchToken(data string) token {
// 	//TODO: перестроить так чтобы он ловил терминальные символы в порядке из стандарта (таблица 1 стр 7)
// 	feed := strings.Split(data, "")
// 	t := token{}
// 	t.ttype = undefined
// 	buf := ""
// 	for i, f := range feed {
// 		buf += f
// 		switch {
// 		default:
// 			t.err = fmt.Errorf("token not defined")
// 		case strings.HasSuffix(buf, ","):
// 			t.ttype = concatenation
// 			t.LHS = formatLHS(buf, ",")
// 			t.RHS = strings.Join(feed[i:], "")
// 		case strings.HasSuffix(buf, "="):
// 			t.ttype = definition
// 			t.LHS = formatLHS(buf, "=")
// 			t.RHS = strings.Join(feed[i:], "")
// 		case strings.HasSuffix(buf, "|"):
// 			t.ttype = alternation
// 			t.LHS = formatLHS(buf, "|")
// 			t.RHS = strings.Join(feed[i:], "")
// 		case strings.HasSuffix(buf, "*)"):
// 			t.err = fmt.Errorf("unimplemented")
// 		case strings.HasSuffix(buf, ")"):
// 			t.err = fmt.Errorf("unimplemented")
// 		case strings.HasSuffix(buf, "]"):
// 			t.err = fmt.Errorf("unimplemented")
// 		case strings.HasSuffix(buf, "}"):
// 			t.err = fmt.Errorf("unimplemented")
// 		case strings.HasSuffix(buf, "-"):
// 			t.err = fmt.Errorf("unimplemented")
// 		case strings.HasSuffix(buf, "`"):
// 			t.err = fmt.Errorf("unimplemented")
// 		case strings.HasSuffix(buf, "*"):
// 			t.err = fmt.Errorf("unimplemented")
// 		case strings.HasSuffix(buf, `"`):
// 			t.err = fmt.Errorf("unimplemented")
// 		case strings.HasSuffix(buf, "?"):
// 			t.err = fmt.Errorf("unimplemented")
// 		case strings.HasSuffix(buf, "(*"):
// 			t.err = fmt.Errorf("unimplemented")
// 		case strings.HasSuffix(buf, "("):
// 			t.err = fmt.Errorf("unimplemented")
// 		case strings.HasSuffix(buf, "["):
// 			t.err = fmt.Errorf("unimplemented")
// 		case strings.HasSuffix(buf, "{"):
// 			t.err = fmt.Errorf("unimplemented")
// 		case strings.HasSuffix(buf, ";"):
// 			t.ttype = termination
// 			t.LHS = formatLHS(buf, ";")
// 			t.RHS = ""
// 			/*
// 		case strings.HasSuffix(buf, "["):
// 			t.ttype = optional
// 			t.LHS = formatLHS(buf, "[")
// 			t.RHS, t.tail, t.err = grabBrackets(strings.Join(feed[i:], ""), "[", "]")
// 			return t
// 		case strings.HasSuffix(buf, "{"):
// 			t.ttype = repetition
// 			t.LHS = formatLHS(buf, "{")
// 			t.RHS, t.tail, t.err = grabBrackets(strings.Join(feed[i:], ""), "{", "}")
// 			return t
// 		case strings.HasSuffix(buf, "("):
// 			t.ttype = grouping
// 			t.LHS = formatLHS(buf, "(")
// 			t.RHS, t.tail, t.err = grabBrackets(strings.Join(feed[i:], ""), "(", ")")
// 			return t
// 		case strings.HasSuffix(buf, `"`):
// 			t.ttype = terminalstring1
// 			t.LHS, t.tail, t.err = grabGridy(strings.Join(feed[i:], ""), `"`)
// 			return t
// 		case strings.HasSuffix(buf, "'"):
// 			t.ttype = terminalstring2
// 			t.LHS, t.tail, t.err = grabGridy(strings.Join(feed[i:], ""), "'")
// 			return t
// 		case strings.HasSuffix(buf, "*"):
// 			t.ttype = terminalstring2
// 			t.LHS, t.tail, t.err = grabGridy(strings.Join(feed[i:], ""), "*")
// 			return t
// 		case strings.HasSuffix(buf, "?"):
// 			t.ttype = specialsequence
// 			t.LHS, t.tail, t.err = grabGridy(strings.Join(feed[i:], ""), "?")
// 			return t
// 		}
// 		*/
// 		t.LHS += f
// 	}
// 	return t
// }

// func formatLHS(s, suf string) string {
// 	s = strings.TrimSuffix(s, suf)
// 	s = strings.TrimSpace(s)
// 	return s
// }

// func grabGridy(data, sep string) (rhs string, rest string, err error) {
// 	if !strings.HasPrefix(data, sep) {
// 		err = fmt.Errorf("token |%v| is not opened with (%v)", data, sep)
// 		rhs = data
// 		return
// 	}
// 	dt := strings.Split(data, sep)
// 	if len(dt) < 3 {
// 		err = fmt.Errorf("token |%v| is not closed with (%v)", data, sep)
// 		rhs = data
// 		return
// 	}
// 	rhs = dt[1]
// 	rest = strings.Join(dt[2:], sep)
// 	return
// }

// func nextToken(data string) (token string, RHS string, LHS string, err error) {
// 	//token, RHS, rest := string, string, string
// 	data = OmitComments(data)
// 	switch {
// 	default:
// 		err = fmt.Errorf("unrecog")
// 		token = "UNRECONIZED"
// 		RHS = data
// 	}

// 	return
// }

// func grabTS1(data string) (RHS string, rest string, err error) {
// 	dt := strings.Split(data, "'")
// 	if len(dt) < 2 {
// 		err = fmt.Errorf("TS1 is not closed: '%v'", data)
// 		RHS = data
// 	}
// 	RHS = dt[1]
// 	rest = strings.Join(dt[2:], "'")
// 	return
// }

// func grabTS2(data string) (RHS string, rest string, err error) {
// 	dt := strings.Split(data, `"`)
// 	if len(dt) < 2 {
// 		err = fmt.Errorf("TS2 is not closed: '%v'", data)
// 		RHS = data
// 	}
// 	RHS = dt[1]
// 	rest = strings.Join(dt[2:], `"`)
// 	return
// }

// // DEFINITIONS:
// func isTerminalString1(data string) bool {
// 	return strings.HasPrefix(data, "'")
// }
// func isTerminalString2(data string) bool {
// 	return strings.HasPrefix(data, `"`)
// }
// func isDefinition(data string) bool {
// 	return strings.HasPrefix(data, "=")
// }
// func isConcatenation(data string) bool {
// 	return strings.HasPrefix(data, ",")
// }
// func isTermination(data string) bool {
// 	return strings.HasPrefix(data, ";")
// }
// func isAlternation(data string) bool {
// 	return strings.HasPrefix(data, "|")
// }
// func isOptional(data string) bool {
// 	return strings.HasPrefix(data, "[")
// }
// func isRepetition(data string) bool {
// 	return strings.HasPrefix(data, "{")
// }
// func isGrouping(data string) bool {
// 	return strings.HasPrefix(data, "(")
// }
// func isComment(data string) bool {
// 	return strings.HasPrefix(data, "(*")
// }
// func isSpecialSequance(data string) bool {
// 	return strings.HasPrefix(data, "?")
// }
// func isException(data string) bool {
// 	return strings.HasPrefix(data, "-")
// }

// //(token, lhs, rhs, tail string, error)
// // func nextToken(data string) (string, string, string, string, error) {
// // 	feed := strings.Split(data, "")

// // 	for _, f := range feed {
// // 		switch f {
// // 		case "'": //termStr
// // 			bod, pre, suf := trimClosingsGreedy(data, "'", "'")
// // 			return "TS1", pre, bod, suf, nil
// // 		case `"`:
// // 			bod, pre, suf := trimClosingsGreedy(data, `"`, `"`)
// // 			return "TS2", pre, bod, suf, nil
// // 		case "[":
// // 			bod, pre, suf := trimClosingsGreedy(data, "[", "]")
// // 			return "OPT", pre, bod, suf, nil
// // 		case "{":
// // 			bod, pre, suf := trimClosingsGreedy(data, "{", "}")
// // 			return "WHILE", pre, bod, suf, nil
// // 		case "(":
// // 			bod, pre, suf := trimClosingsGreedy(data, "(", ")")
// // 			return "grouping", pre, bod, suf, nil
// // 		case "?":
// // 			bod, pre, suf := trimClosingsGreedy(data, "?", "?")
// // 			return "special_seq", pre, bod, suf, nil
// // 		}
// // 	}
// // 	return "", "", "", "", fmt.Errorf("can't find token")
// // }

// /*

// [{'abc'}]"de"

// */

// //func parseTokens(data string) tokenList {

// //}

// func terminated(rbnf string) bool {
// 	return strings.HasSuffix(rbnf, ";")
// }

// func haveClosings(rbnf string) int {
// 	if strings.Contains(rbnf, "[") && strings.Contains(rbnf, "]") {
// 		return 1
// 	}

// 	return 0
// }

// // //////////////////////////////
// func trimClosingsLazy(str, open, close string) (string, string, string) {
// 	if !strings.Contains(str, open) || !strings.Contains(str, close) {
// 		return str, "", ""
// 	}
// 	pre, str1 := trimPrefixBefore(str, open)
// 	str2, suf := trimPrefixBefore(str1, close)
// 	return str2, pre, suf
// }

// func trimClosingsGreedy(str, open, close string) (string, string, string) {
// 	if !strings.Contains(str, open) || !strings.Contains(str, close) {
// 		return str, "", ""
// 	}
// 	pre, str1 := trimPrefixBefore(str, open)
// 	suf, str2 := trimSuffixBefore(str1, close)
// 	return str2, pre, suf
// }

// func trimPrefixBefore(str, trig string) (string, string) {
// 	arr1 := strings.Split(str, trig)
// 	return arr1[0], strings.Join(arr1[1:], trig)
// }

// func trimSuffixBefore(str, trig string) (string, string) {
// 	arr := strings.Split(str, trig)
// 	return arr[len(arr)-1], strings.Join(arr[:len(arr)-1], trig)
// }

// func trimFirstRune(s string) string {
// 	for i := range s {
// 		if i > 0 {
// 			return s[i:]
// 		}
// 	}
// 	return ""
// }

// func grabBrackets(str string, op, cl string) (capt, rest string, err error) {
// 	if op == cl {
// 		return str, "", fmt.Errorf("brackets must not match ('%v'='%v')", op, cl)
// 	}
// 	depth := 0
// 	rhs := ""
// 	for str != "" {
// 		switch {
// 		case strings.HasPrefix(str, op):
// 			str = strings.TrimPrefix(str, op)
// 			depth++
// 			if depth != 1 {
// 				rhs += op
// 				continue
// 			}
// 		case strings.HasPrefix(str, cl):
// 			str = strings.TrimPrefix(str, cl)
// 			depth--
// 			if depth != 0 {
// 				rhs += cl
// 				continue
// 			}
// 			return rhs, str, nil
// 		case depth == 0:
// 			return "", str, fmt.Errorf("brackets were not opened")
// 		}
// 		rhs += firstN(str, 1)
// 		str = trimFirstRune(str)
// 	}
// 	return rhs, str, fmt.Errorf("brackets were not closed")
// }

// func firstN(s string, n int) string {
// 	r := []rune(s)
// 	if len(r) > n {
// 		return string(r[:n])
// 	}
// 	return s
// }
