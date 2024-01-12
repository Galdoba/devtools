package cellcont

import (
	"fmt"
	"strings"
)

const (
	Align_Left = iota
	Align_Center
	Align_Right
)

type cellcontent struct {
	align         int
	columnWidth   int
	rawText       string
	textSeparated []string
}

func New(width int) (*cellcontent, error) {
	cell := &cellcontent{}
	if err := cell.SetColumnWidth(width); err != nil {
		return cell, fmt.Errorf("can't create new content: %v", err.Error())
	}
	return cell, nil
}

type Content interface {
	SetAlign(int) error
	SetColumnWidth(int) error
	SetText(string) error
}

func (c *cellcontent) SetText(text string) error {
	if c.columnWidth < 1 {
		return fmt.Errorf("can't set text: column width is not set")
	}
	c.rawText = text
	textSep := separateGlyphs(text, c.columnWidth)
	textSep = align(textSep, c.align, c.columnWidth)
	c.textSeparated = textSep
	return nil
}

func (c *cellcontent) SetAlign(a int) error {
	switch a {
	default:
		return fmt.Errorf("can't set align: align value '%v' unknown", a)
	case Align_Left, Align_Center, Align_Right:
		c.align = a
		c.textSeparated = align(c.textSeparated, c.align, c.columnWidth)
	}
	return nil
}

func (c *cellcontent) SetColumnWidth(l int) error {
	if l < 1 {
		return fmt.Errorf("can't set column len: column width is %v must be >= 1", l)
	}
	c.columnWidth = l
	return nil
}

////////////////////////////

func separateGlyphs(text string, rowLen int) []string {
	letters := strings.Split(text, "")
	if len(letters) <= rowLen {
		return []string{strings.Join(letters, "")}
	}
	out := []string{}
	s := ""
	for l := 0; l < len(letters); l++ {
		if l%rowLen == 0 && s != "" {
			out = append(out, s)
			s = ""
		}
		s += letters[l]
	}
	out = append(out, s)
	return out
}

func align(textSep []string, a int, w int) []string {
	alighFunc := make(map[int]func(string, int) string)
	alighFunc[Align_Left] = addToLenLeft
	alighFunc[Align_Center] = addToLenCenter
	alighFunc[Align_Right] = addToLenRight
	for i := range textSep {
		textSep[i] = alighFunc[a](textSep[i], w)
	}
	return textSep
}

func addToLenLeft(s string, l int) string {
	for len(s) < l {
		s += " "
	}
	return s
}

func addToLenRight(s string, l int) string {
	for len(s) < l {
		s = " " + s
	}
	return s
}

func addToLenCenter(s string, l int) string {
	s = strings.TrimSpace(s)
	for len(s) < l {
		s = " " + s + " "
	}
	if len(s) > l {
		s = strings.TrimPrefix(s, " ")
	}
	return s
}
