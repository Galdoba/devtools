package parser

import (
	"fmt"
)

const (
	RuneEOF = 0
)

type ParserFunc = func(*Reader) (*Result, *Error)

type Reader struct {
	data    []byte
	pos     Position
	prevCol int
}

type Position struct {
	offset int
	line   int
	col    int
}

func (p *Position) ToString() string {
	return fmt.Sprintf("offset: %v (ln %v col %v)", p.offset, p.line+1, p.col+1)
}

func NewReader(data string) *Reader {
	r := Reader{}
	r.data = []byte(data)
	return &r
}

func (r *Reader) Read() byte {
	if r.pos.offset >= len(r.data) {
		return RuneEOF
	}
	chr := r.data[r.pos.offset]
	//смещение физической позиции
	r.pos.offset++
	//смещение абстрактной (логической) позиции
	r.prevCol = r.pos.col
	r.pos.col++
	if chr == '\n' {
		r.pos.col = 0
		r.pos.line++
	}

	return chr
}

func (rd *Reader) Data(startPos Position) string {
	return string(rd.data[startPos.offset:rd.pos.offset])

}

func (r *Reader) UnRead() {
	//смещение физической позиции
	r.pos.offset--
	//смещение абстрактной (логической) позиции
	chr := r.data[r.pos.offset]
	r.pos.col--
	if chr == '\n' {
		r.pos.col = r.prevCol
		r.pos.line--
	}
}

func (r *Reader) Save() Position {
	return r.pos
}

func (r *Reader) Restore(p Position) {
	r.pos = p
}

func (r *Reader) Reset() {
	r.pos = Position{}
}

//////////////////////////////////////////////////////////////////////////

type Error struct {
	errors []parseError
}

type parseError struct {
	message string
	pos     Position
}

func NewError() *Error {
	return &Error{}
	//return &Error{[]parseError{parseError{msg, p}}}
}

func (e *Error) Add(p Position, msg string) *Error {
	e.errors = append(e.errors, parseError{msg, p})
	return e
}

func (e *Error) ToString() string {
	str := "Errors:\n"
	l := len(e.errors) - 1
	for i := range e.errors {
		str += fmt.Sprintf("  %v: %v\n", e.errors[l-i].pos.ToString(), e.errors[l-i].message)
	}
	return str
}

////////////////////////////////////////////////////////////////////////
type Result struct {
	kvs []keyval
}

func NewResult(key, val string) *Result {
	return &Result{
		kvs: []keyval{
			{k: key, v: val},
		},
	}
}

func (res *Result) ToString() string {
	if res == nil {
		return "ok"
	}
	str := ""
	for _, kv := range res.kvs {
		str += fmt.Sprintf("key: %v  val: %v\n", kv.k, kv.v)
	}
	return str
}

func AppendResult(res, new *Result) *Result {
	if new == nil {
		return res
	}
	if res == nil {
		res = &Result{}
	}
	res.kvs = append(res.kvs, new.kvs...)
	return res
}

type keyval struct {
	k string
	v string
}

////////////////////////////////////////////////////////////////////////
func BaseLogic(rd *Reader, arg interface{}) (*Result, *Error) {
	pos := rd.Save()

	switch val := arg.(type) {
	default:
		panic(fmt.Sprintf("unallowed state %T", val))
	case int32:
		chr := rd.Read()
		v := byte(val)
		if chr != v {
			return nil, NewError().Add(pos, fmt.Sprintf("expected %q, got %q", v, chr))
		}
	case string:
		if val == "" {
			return nil, nil
		}
		for _, s := range []byte(val) {
			chr := rd.Read()
			if chr != s {
				return nil, NewError().Add(pos, fmt.Sprintf("read string: %v: expected %q, got %q", val, s, chr))
			}
		}
	case ParserFunc:
		rs, err := val(rd)
		if err != nil {
			return nil, err
		}
		return rs, nil
	} //switch
	return nil, nil
}

//////////////////////////////////////////////////////////////////////////
//-Keep сохраняет найденное
func Keep(name string, arg interface{}) ParserFunc {
	//res, err := Choose(arg, "")
	return func(rd *Reader) (*Result, *Error) {
		pos := rd.Save()
		fn := JOIN(arg)
		res, err := fn(rd)
		if err != nil {
			return nil, err.Add(pos, fmt.Sprintf("Keep %q ненашел то что искал ", name))
		}
		res = AppendResult(res, NewResult(name, rd.Data(pos)))
		return res, nil
	}

}

func Run(data string, fn ParserFunc) (*Result, error) {
	rd := NewReader(data)
	res, e := fn(rd)
	if e != nil {
		return nil, fmt.Errorf("run(): %v", e.ToString())
	}
	return res, nil
}
