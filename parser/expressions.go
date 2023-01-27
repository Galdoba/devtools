package parser

/*
JOIN
JOIN *
CONCAT
CNCT
SUM

*/
func JOIN(args ...interface{}) ParserFunc {
	return func(rd *Reader) (*Result, *Error) {
		res := (*Result)(nil)
		for _, arg := range args {
			rs, err := BaseLogic(rd, arg)
			if err != nil {
				return nil, err
			}
			res = AppendResult(res, rs)
		}
		return res, nil
	} //func
}

//подумать над именем //choice?
/*
OR
CHOOSE
SELECT
SEL
ONEOF
*/
func OR(args ...interface{}) ParserFunc {
	return func(rd *Reader) (*Result, *Error) {
		pos := rd.Save()
		for _, arg := range args {
			res, err := BaseLogic(rd, arg)
			if err == nil {
				return res, err
			}
			rd.Restore(pos)
		} //for
		return nil, NewError().Add(pos, "choice invalid")
	} //func
}

//OPT 0 или 1 arg
//MAYBE
//MYB
//MIGHT
//OPT*
//NOTMORETHANONCE
//
func OPT(arg interface{}) ParserFunc {
	return func(rd *Reader) (*Result, *Error) {
		pos := rd.Save()
		res, err := BaseLogic(rd, arg)
		if err != nil {
			rd.Restore(pos)
		}
		return res, nil
	}
}

func WHILE(arg interface{}) ParserFunc {
	return func(rd *Reader) (*Result, *Error) {
		res := (*Result)(nil)
		for {
			pos := rd.Save()
			rs, err := BaseLogic(rd, arg)
			if err != nil {
				rd.Restore(pos)
				return res, nil
			}
			res = AppendResult(res, rs)
		}
	}
}

func NOT(arg interface{}) ParserFunc {
	return func(rd *Reader) (*Result, *Error) {
		pos := rd.Save()
		_, err := BaseLogic(rd, arg)
		rd.Restore(pos)
		switch err {
		default:
			err = nil
		case nil:
			err = NewError().Add(pos, "NOT3 is not true")
		}

		return nil, err
	}
}

func LIT(fn func(chr byte) bool) ParserFunc {
	return func(rd *Reader) (*Result, *Error) {
		//сравнивается chr и чтение их ридера по chr из вне
		if fn(rd.Read()) {
			return nil, nil
		}
		rd.UnRead()
		return nil, NewError()
	}
}

/*
Till - с заглатыванием символа
Untill - без заглатывания символа
While - читаем
Skip ??
*/
