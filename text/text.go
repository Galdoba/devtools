package text

import "strings"

func EnsureASCII(s string) string {
	gl := strings.Split(s, "")
	result := ""
	for _, g := range gl {
		switch g {
		default:
			result += "?"
		case "!", "#", "$", "%", "&", "'", "(", ")", "*", "+", ",", "-", ".",
			"/", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ":", ";",
			"<", "=", ">", "?", "@", "A", "B", "C", "D", "E", "F", "G", "H",
			"I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U",
			"V", "W", "X", "Y", "Z", "[", "]", "^", "_", "`",
			"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
			"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
			"{", "|", "}", "~", " ":
			result += g
		case `"`, `\`:
			result += g
		}
	}
	return result
}

func PickGlyph(pick string, defaultGlyph string, optional ...string) string {
	for _, opt := range optional {
		if opt == pick {
			return pick
		}
	}
	return defaultGlyph
}
