package cellcont

import (
	"strings"
	"testing"
)

func TestCont(t *testing.T) {
	testTexts := []string{
		"Short",
		"Shrt",
		"Sht",
		"not so shooort ssssssssssssssssssssssssssssssssss",
		"",
	}
	testWidth := []int{5, 10, 2, 1, 0, -1}

	for _, w := range testWidth {
		cell, err := New(w)
		if w < 1 {
			if err == nil {
				t.Errorf("expect error, but have none")
			} else {
				t.Logf("expected error: %v", err.Error())
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
