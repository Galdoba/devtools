package csvp

import (
	"fmt"
	"strings"
)

func NewContainer() *container {
	return &container{}
}

type container struct {
	entries []*entry
}

func (c *container) Entries() []*entry {
	return c.entries
}

type entry struct {
	fields []string
}

func (e *entry) Fields() []string {
	return e.fields
}

func NewEntry(flds ...string) *entry {
	e := entry{}
	for _, f := range flds {
		e.fields = append(e.fields, asField(f))
	}
	return &e
}

func (c *container) ensureSquare() {
	width := 0
	for _, entry := range c.entries {
		for k := range entry.fields {
			if k+1 > width {
				width = k + 1
			}
		}
	}
	for e, entry := range c.entries {
		for len(entry.fields) < width {
			entry.fields = append(entry.fields, " ")
		}
		c.entries[e] = NewEntry(entry.fields...)
	}
}

func (c *container) Field(e, f int) string {
	if e < 0 || f < 0 {
		return ""
	}
	if len(c.entries) > e {
		return ""
	}
	entry := c.entries[e]
	if len(entry.fields) > f {
		return ""
	}
	return asField(entry.fields[f])
}

func FromString(str string) (*container, error) {
	c := container{}
	parts := strings.Split(str, "\n")
	prevParts := ""
	for i := range parts {
		if e, err := EntryFromString(prevParts + parts[i]); err == nil {
			c.entries = append(c.entries, e)
			prevParts = ""
			continue
		}
		prevParts += parts[i] + "\n"

	}
	if prevParts != "" {
		return nil, fmt.Errorf("part of container was not decoded:\n%v", prevParts)
	}
	return &c, nil
}

func FieldFromString(s string) (string, error) {
	if strings.HasPrefix(s, `"`) {
		if strings.HasSuffix(s, `"`) {
			s = strings.TrimPrefix(s, `"`)
			s = strings.TrimSuffix(s, `"`)
			s = strings.ReplaceAll(s, `""`, `"`)
			return s, nil
		}
		return "", fmt.Errorf("expect double-qutes at the end '%v'", s)
	}
	return s, nil
}

func EntryFromString(s string) (*entry, error) {
	e := entry{}
	parts := strings.Split(s, ",")
	prevParts := ""
	for i := range parts {
		if f, err := FieldFromString(prevParts + parts[i]); err == nil {
			e.fields = append(e.fields, f)
			prevParts = ""
			continue
		}
		prevParts += parts[i] + ","
	}
	if prevParts != "" {
		return nil, fmt.Errorf("part of entry was not decoded:\n'%v'", prevParts)
	}
	return &e, nil

}

// func toField(s string) string {
// 	if strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) {
// 		s = strings.TrimPrefix(s, `"`)
// 		s = strings.TrimSuffix(s, `"`)
// 		//		s = strings.ReplaceAll(s, `""`, `"`)
// 		return s
// 	}
// 	// s = strings.ReplaceAll(s, `""`, `"`)
// 	return s
// }

func (c *container) String() string {
	s := ""
	for _, e := range c.entries {
		s += e.String() + "\n"
	}
	return strings.TrimSuffix(s, "\n")
}

func (c *container) TableView() string {
	//data := [][]string{}
	//TODO: вывести на экран в виде таблицы
	return ""
}

func (e *entry) String() string {
	s := ""
	for _, f := range e.fields {
		s += asField(f) + ","
	}
	return strings.TrimSuffix(s, ",")
}

func asField(field string) string {
	for _, glyph := range []string{",", "\n"} {
		if strings.Contains(field, glyph) && !strings.HasPrefix(field, `"`) && !strings.HasSuffix(field, `"`) {
			return `"` + field + `"`
		}
	}
	return field
}
