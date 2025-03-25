package text

import (
	"strings"
)

type wrapper struct {
	maxWidth    int
	leftOffset  int
	rightOffset int
	wrapLimit   int
	wrapBefore  []string
	wrapAfter   []string
}

func Wrap(text string, options ...WrappingOption) string {
	wr := wrapper{
		maxWidth:   -1,
		wrapAfter:  []string{" ", "-", "\t"},
		wrapBefore: []string{"(", "[", "{", "`", `"`, "<"},
	}
	for _, set := range options {
		set(&wr)
	}
	schema := schema(wr, text)
	if wr.maxWidth < 1 {
		return text
	}
	lines := composeLines(wr, schema)

	return strings.Join(lines, "\n")
}

type WrappingOption func(*wrapper)

func MaxWidth(mw int) WrappingOption {
	return func(w *wrapper) {
		w.maxWidth = mw
	}
}

func LeftOffset(lo int) WrappingOption {
	return func(w *wrapper) {
		w.leftOffset = lo
	}
}

func WrapLimit(wl int) WrappingOption {
	return func(w *wrapper) {
		w.wrapLimit = wl
	}
}

type wrappingSchema map[int]litera

type litera struct {
	text       string
	rn         rune
	maybeStart bool
	maybeEnd   bool
}

var lit_Space = litera{
	text:       " ",
	rn:         32,
	maybeStart: true,
	maybeEnd:   true,
}

var lit_NewLine = litera{
	text:       "\n",
	rn:         10,
	maybeStart: false,
	maybeEnd:   true,
}

func schema(wr wrapper, text string) wrappingSchema {
	ws := make(wrappingSchema)
	lits := strings.Split(text, "")
	for i, l := range lits {
		lit := litera{}
		lit.text = l
		lit.maybeStart = inSlice(wr.wrapBefore, l)
		lit.maybeEnd = inSlice(wr.wrapAfter, l)
		for _, r := range l {
			lit.rn = r
		}
		ws[i] = lit
	}
	return ws
}

func inSlice(sl []string, s string) bool {
	for _, text := range sl {
		if text == s {
			return true
		}
	}
	return false
}

func composeLines(wr wrapper, sc wrappingSchema) []string {
	maxWidth := wr.maxWidth
	lines := []string{}
	candidate := newCandidate(0)
	trimmed := -1
	lenText := len(sc)
	for i := 0; i < lenText; i++ {

		lit := sc[i]
		if lit.text == "\n" {
			lines = append(lines, toLine(candidate))
			candidate = newCandidate(wr.leftOffset)
			continue
		}
		candidate = append(candidate, lit)

		if len(candidate) == maxWidth {
			candidate, trimmed = candidate.trimSuffix(wr.wrapLimit)
			switch trimmed {
			case 0:
				lines = append(lines, toLine(candidate))
				candidate = newCandidate(wr.leftOffset)
			default:
				lines = append(lines, toLine(candidate))
				candidate = newCandidate(wr.leftOffset)
				i = i - trimmed
			}

		}

	}
	lines = append(lines, toLine(candidate))
	return lines
}

type candidate []litera

func newCandidate(offset int) candidate {
	cnd := candidate{}
	for i := 0; i < offset; i++ {
		cnd = append(cnd, lit_Space)
	}
	return cnd
}

func (cnd candidate) trimSuffix(trimax int) (candidate, int) {
	trimmed := 0
	origin := cnd
	for !cnd[len(cnd)-1].maybeEnd {
		if trimmed >= trimax {
			return origin, 0
		}
		cnd = cnd[:len(cnd)-1]
		trimmed++
		if len(cnd) == 0 {
			return origin, 0
		}
		if cnd[len(cnd)-1].maybeStart {
			cnd = cnd[:len(cnd)-1]
			trimmed++
			return cnd, trimmed
		}

	}
	return cnd, trimmed
}

func toLine(cnd candidate) string {
	line := ""
	for _, lit := range cnd {
		line += lit.text
	}
	return line
}
