package cellcont

import (
	"fmt"
	"strings"
	"testing"
)

func TestContentInput(t *testing.T) {
	testTexts := []string{
		"Short",
		"Shrt",
		"Sht",
		"not so shooort ssssssssssssssssssssssssssssssssss",
		"",
	}
	testWidth := []int{5, 10, 2, 1, 0, -1}

	for _, w := range testWidth {
		cell := New()
		err := cell.SetText("test")
		if err != nil {
			if strings.Contains(err.Error(), "column width is not set") {
				t.Logf("expected error: %v", err.Error())
			} else {
				t.Errorf("unattended error: %v", err.Error())
			}
		}
		err = cell.SetColumnWidth(w)
		if err != nil {
			switch {
			case strings.Contains(err.Error(), "must be >= 1"):
				t.Logf("expected error from input (%v): %v", w, err.Error())
				continue
			default:
				t.Errorf("unattended Error: %v", err.Error())
				//continue
			}
		}
		for _, txt := range testTexts {
			txtErr := cell.SetText(txt)
			if txtErr != nil {
				if strings.Contains(txtErr.Error(), "width is not set") {
					t.Logf("expected error: %v", err.Error())
					continue
				}
				t.Errorf("untested error: %v", txtErr.Error())
				continue
			}
			for i := -1; i <= 4; i++ {
				switch i {
				default:
					err := cell.SetAlign(i)
					t.Logf("%v", err.Error())
					continue
				case Align_Left, Align_Center, Align_Right:
					err := cell.SetAlign(i)
					if err != nil {
						t.Errorf("%v", err.Error())
					}
				}

			}
		}
	}

}

func TestContentOutput(t *testing.T) {
	for _, w := range []int{20, 10, 5, 3, 2, 1, 0, -1} {
		cell := New()
		if err := cell.SetColumnWidth(w); err != nil {
			t.Logf("expected error: %v", err.Error())
		}
		for _, inputText := range []string{
			"input text",
			"input text very looooooooooooooooong",
			"input text very looooooooooooooooong, and even loooooooooooooooooooooooooooooooooooooooooooooonger",
			"t",
			"",
		} {

			if err := cell.SetText(inputText); err != nil {
				t.Logf("expected error: %v", err.Error())
			}

			rows := []int{2, 1, 0, -1}
			for _, r := range rows {
				for _, align := range []int{Align_Center, Align_Right, Align_Left} {
					if err := cell.SetAlign(align); err != nil {
						t.Logf("expected error: %v", err.Error())
					}
					textRow := cell.TextRow(r)
					if r < 0 && textRow != "" {
						t.Errorf("expected empty string, but have '%v'", textRow)
					}
					if r >= len(cell.textSeparated) && textRow != "" {
						t.Errorf("expected empty string, but have '%v'", textRow)
					}
					t.Logf("valid: %v", fmt.Sprintf("w=%v,r=%v,a=%v textLen=%v:  '%v'\n", w, r, align, len(inputText), textRow))
				}
				wrapVals := []bool{true, false, true}
				for _, wrapval := range wrapVals {
					cell.SetWrap(wrapval)
					if cell.Wrap() != wrapval {
						t.Errorf("SetWrap(bool) did not worked")
					}
				}
				wr := cell.Wrap()
				cell.ToggleWrap()
				if wr == cell.Wrap() {
					t.Errorf("ToggleWrap(bool) did not worked")
				}
				cell.ToggleWrap()
				text := cell.Text()
				row0 := cell.TextRow(0)
				if cell.wrap && len(cell.rawText) > len(text) && len(text) > 3 {
					if !strings.HasSuffix(text, "..") {
						t.Errorf("expected '..' at the end of '%v'", text)
					}

				}
				t.Logf("text: %v", text)
				t.Logf("text[0]: %v", row0)
				//fmt.Printf("'%v'|'%v'\n", text, row0)

			}
		}
	}
}
