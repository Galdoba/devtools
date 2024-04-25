package tui

import (
	"fmt"
	"strings"

	"github.com/Galdoba/devtools/app/configbuilder/internal/model"
)

func Status(m *model.Model) string {
	stat := "Current Model: " + fmt.Sprintf("%v fields\n", len(m.Fields))
	report := []string{}
	for _, fld := range m.Fields {
		// stat += fld.String() + "\n"
		report = append(report, fld.String())
	}
	report = formatFieldReport(report)
	for _, r := range report {
		stat += r + "\n"
	}
	stat += "Error: " + modelErrorText(m)
	return stat
}

func modelErrorText(m *model.Model) string {
	for i, f := range m.Fields {
		if err := f.Validate(); err != nil {
			return fmt.Sprintf("field %v: %v", i+1, err.Error())
		}
	}
	return "none"
}

func formatFieldReport(sl []string) []string {
	lens := []int{0, 0, 0}
	for _, s := range sl {
		data := strings.Fields(s)
		for i := 0; i < 2; i++ {
			lens[i] = max(lens[i], len(strings.Split(data[i], "")))
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
		// word3 := data[2]
		// for len(strings.Split(word3, "")) < lens[2] {
		// 	word3 += " "
		// }

		out = append(out, word1+" "+word2+" "+strings.Join(data[2:], " "))
	}
	outF := []string{}
	if len(out) > 0 {
		outF = append(outF, "=============================================")
		outF = append(outF, out...)
		outF = append(outF, "=============================================")
	}
	return outF
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
