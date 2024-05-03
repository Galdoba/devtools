package tui

import (
	"fmt"
	"strings"

	"github.com/Galdoba/devtools/app/configbuilder/internal/model"
	"github.com/fatih/color"
)

func Status(m *model.Model) string {
	stat := "Model: " + fmt.Sprintf("contains %v fields\n", len(m.Fields))
	stat += "Encoding: " + fmt.Sprintf("%v\n", m.Language())
	report := []string{}
	for _, fld := range m.Fields {
		// stat += fld.String() + "\n"
		report = append(report, fld.String())
	}
	report = formatFieldReport(report)
	for _, r := range report {
		stat += r + "\n"
	}

	stat += "Model Validation: " + modelErrorText(m)
	return stat
}

func modelErrorText(m *model.Model) string {
	for i, f := range m.Fields {
		if err := f.Validate(); err != nil {
			return color.HiYellowString(fmt.Sprintf("field %v: %v", i+1, err.Error()))
		}
	}
	return color.GreenString("ok")
}

func formatFieldReport(sl []string) []string {
	lens := []int{0, 0, 0, 0}
	for _, s := range sl {
		data := strings.Fields(s)

		for i := 0; i <= 3; i++ {
			val := ""
			switch i {
			case 0, 1:
				val = data[i]
			case 2:
				for _, v := range data[2:] {
					if strings.Contains(v, "//") {
						break
					}
					val += v + " "
				}

			}
			lens[i] = max(lens[i], len(strings.Split(val, "")))

		}
	}

	out := []string{}
	for _, s := range sl {
		data := strings.Fields(s)
		word1 := data[0]
		for len(strings.Split(word1, "")) < lens[0] {
			word1 += " "
		}
		word2 := data[1]
		for len(strings.Split(word2, "")) < lens[1] {
			word2 += " "
		}
		rest := strings.Join(data[2:], " ")
		encodingText, other := chopHead(rest, "//")
		for len(encodingText) < (lens[2] + 2) {
			encodingText += " "
		}
		encodingText = color.YellowString(encodingText)
		comment, defaults := chopHead(other, " [")
		if defaults == "[ : ]" {
			defaults = ""
		}

		comment = color.GreenString(comment)
		str := word1 + "  " + color.HiCyanString(word2) + "  " + encodingText + comment + defaults

		out = append(out, str)
	}
	outF := []string{}
	if len(out) > 0 {
		outF = append(outF, "=============================================")
		outF = append(outF, out...)
		outF = append(outF, "=============================================")
	}
	return outF
}

func chopHead(str, pref string) (string, string) {
	letters := strings.Split(str, "")
	head := ""
	for i, l := range letters {
		body := strings.Join(letters, "")
		if strings.HasPrefix(body[i:], pref) {
			break
		}
		head += l
	}
	rest := strings.TrimPrefix(str, head)
	return head, rest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
Current Model:
als;dk string ashdga
---------------------
Error: none

*/
