package text

import "strings"

func WrapByWidth(txt string, wd int) []string {
	letters := split(txt)
	rows := []string{}
	row := ""
	for i, l := range letters {
		if i != 0 && i%wd == 0 {
			rows = append(rows, row)
			row = ""
		}
		row += l
	}
	if row != "" {
		rows = append(rows, row)
	}
	return rows
}

func WrapByWords(txt string, wd int) []string {
	words := words(txt, wd)
	//letters := split(txt)
	rows := []string{}
	row := ""
	for _, wrd := range words {
		if len(split(row))+len(wrd) <= wd {
			row += wrd
			if len(split(row)) < wd {
				row += " "
			}
			continue
		} else {
			rows = append(rows, row)
			row = ""
		}

	}
	if row != "" {
		rows = append(rows, row)
	}
	return rows
}

func SetLength(txt string, l int) string {
	if l <= 0 {
		return ""
	}
	for l > trueLen(txt) {
		txt += " "
	}
	if trueLen(txt) > l {
		ltrs := split(txt)
		n := 2
		suff := ".."
		if l < 5 {
			n = 0
			suff = ""
		}
		txt = join(ltrs[0:l-n]) + suff
	}
	return txt
}

//helpers
func split(s string) []string {
	return strings.Split(s, "")
}

func join(l []string) string {
	return strings.Join(l, "")
}

func words(s string, max int) []string {
	flds := strings.Fields(s)
	out := []string{}
	for _, f := range flds {
		parts := WrapByWidth(f, max)
		for _, p := range parts {
			out = append(out, p)
		}
	}
	return out
}

func trueLen(s string) int {
	l := split(s)
	return len(l)
}
