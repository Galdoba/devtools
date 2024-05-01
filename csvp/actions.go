package csvp

import "fmt"

// AppendEntry - add entry at the end
func (c *container) AppendEntry(e *entry) {
	c.entries = append(c.entries, e)
	c.ensureSquare()
}

// PrependEntry - add entry at the beginning
func (c *container) PrependEntry(e *entry) {
	c.entries = append([]*entry{e}, c.entries...)
	c.ensureSquare()
}

// InsertEntryAfter - add entry to csv container after n entries
func (c *container) InsertEntryAfter(e *entry, n int) error {
	if n >= len(c.entries) || n < 0 {
		return fmt.Errorf("index is out of bounds")
	}
	out := []*entry{}
	for i, oe := range c.entries {
		out = append(out, oe)
		if i == n {
			out = append(out, e)
		}
	}
	c.entries = out
	c.ensureSquare()
	return nil
}

// DeleteEntry - delete entry number n
func (c *container) DeleteEntry(n int) error {
	if n >= len(c.entries) || n < 0 {
		return fmt.Errorf("index is out of bounds")
	}
	out := []*entry{}
	for i, oe := range c.entries {
		if i == n {
			continue
		}
		out = append(out, oe)
	}
	c.entries = out
	c.ensureSquare()
	return nil
}

// SwitchEntries - switch two entries between them selves
func (c *container) SwitchEntries(n, m int) error {
	if n < 0 || n >= len(c.entries) {
		return fmt.Errorf("index is out of bounds")
	}
	if m < 0 || m >= len(c.entries) {
		return fmt.Errorf("index is out of bounds")
	}
	if n == m {
		return fmt.Errorf("expect different indexes to switch")
	}
	c.entries[n], c.entries[m] = c.entries[m], c.entries[n]
	c.ensureSquare()
	return nil
}

// AppendColumn - append fields to each entries. arguments used
// as value for new fields
func (c *container) AppendColumn(flds ...string) {
	for i, e := range c.entries {
		newFld := " "
		if i < len(flds) {
			newFld = flds[i]
		}
		entr := NewEntry()
		entr.fields = append(entr.fields, e.fields...)
		entr.fields = append(entr.fields, newFld)
		c.entries[i] = entr
	}
	c.ensureSquare()
}

// PrependColumn - prepend fields to each entries. arguments used
// as value for new fields
func (c *container) PrependColumn(flds ...string) {
	for i, e := range c.entries {
		newFld := " "
		if i < len(flds) {
			newFld = flds[i]
		}
		entr := NewEntry()
		entr.fields = append(entr.fields, newFld)
		entr.fields = append(entr.fields, e.fields...)
		c.entries[i] = entr
	}
	c.ensureSquare()
}

// InsertColumnAfter - add fields to each entries after col-th field. arguments used
// as value for new fields
func (c *container) InsertColumnAfter(col int, flds ...string) {
	for i, e := range c.entries {
		newFld := " "
		if i < len(flds) {
			newFld = flds[i]
		}
		entr := NewEntry()
		for j, f := range e.fields {
			entr.fields = append(entr.fields, f)
			if j == col {
				entr.fields = append(entr.fields, newFld)
			}
		}
		c.entries[i] = entr
	}
	c.ensureSquare()
}

// DeleteColumn - delete n-th field in each entry
func (c *container) DeleteColumn(col int) error {
	if len(c.entries) < 1 {
		return fmt.Errorf("nothing to delete")
	}
	if col >= len(c.entries[0].fields) || col < 0 {
		return fmt.Errorf("index is out of bounds")
	}
	for i, e := range c.entries {
		entr := NewEntry()
		for j, f := range e.fields {
			if j == col {
				continue
			}
			entr.fields = append(entr.fields, f)
		}
		c.entries[i] = entr
	}
	c.ensureSquare()
	return nil
}

// SwitchColumns - switch n-th and m-th field in each entry
func (c *container) SwitchColumns(n, m int) error {
	if len(c.entries) < 1 {
		return fmt.Errorf("container is blank")
	}
	e0 := c.entries[0]
	if len(e0.fields) >= n || len(e0.fields) >= m {
		return fmt.Errorf("index is out of bounds")
	}
	if n < 0 || m < 0 {
		return fmt.Errorf("index is out of bounds")
	}
	if n == m {
		return fmt.Errorf("expect different indexes to switch")
	}
	for a, entr := range c.entries {
		newEntr := NewEntry()
		newEntr.fields = append(newEntr.fields, entr.fields...)
		newEntr.fields[m], newEntr.fields[n] = newEntr.fields[n], newEntr.fields[m]
		c.entries[a] = newEntr
	}
	c.ensureSquare()
	return nil

}

/////////CELLS

func (c *container) GetFieldValue(row, col int) string {
	if len(c.entries) <= row || row < 0 {
		return ""
	}
	if len(c.entries[0].fields) <= col || col < 0 {
		return ""
	}
	return c.entries[row].fields[col]
}

func (c *container) SetFieldValue(row, col int, newVal string) error {
	if len(c.entries) <= row || row < 0 {
		return fmt.Errorf("bad row index provided")
	}
	if len(c.entries[0].fields) <= col || col < 0 {
		return fmt.Errorf("bad col index provided")
	}
	c.entries[row].fields[col] = newVal
	return nil
}
