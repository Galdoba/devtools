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
	try := 0
	for dataNew != data {
		try++
		if try > 3 {
			return data
		}
		data = dataNew
		dataNew = debracketGreedy(data, "(*", "*)")
	}
	return data
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
*/
//(token, lhs, rhs, tail string, error)
func nextToken(data string) (string, string, string, string, error) {
	feed := strings.Split(data, "")
	for _, f := range feed {
		switch f {
		case "'": //termStr
			bod, pre, suf := trimClosingsGreedy(data, "'", "'")
			return "TS1", pre, bod, suf, nil
		case `"`:
			bod, pre, suf := trimClosingsGreedy(data, `"`, `"`)
			return "TS2", pre, bod, suf, nil
		}
	}
	return "", "", "", "", fmt.Errorf("can't find token")
}

func terminated(rbnf string) bool {
	return strings.HasSuffix(rbnf, ";")
}

func haveClosings(rbnf string) int {
	if strings.Contains(rbnf, "[") && strings.Contains(rbnf, "]") {
		return 1
	}

	return 0
}

/*
aaaaa[sssss]bbbb
0aaaaa
1sssss]bbbb
0sssss
1bbbb

aa[bb[sssss]cc]dd split(str, "[") == split(1:)
aa bb +[+ sssss]cc]dd

bb[sssss]cc]dd
bb[sssss]cc +]+ dd

bb[sssss]cc
*/

func collect(input, open, close string) string {
	bodyTailSl := strings.Split(input, open)
	bodyTail := strings.Join(bodyTailSl[1:], "")
	body := strings.Split(bodyTail, close)
	bd := strings.Join(body[:len(body)-1], "")
	fmt.Println("col", bd)
	return bd
}

/*
[def] => [def];
[[def]] => OPT(OPT(def));
*/
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
