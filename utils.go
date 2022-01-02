package errors

func Is(err error, target error) bool {
	return Equals(err, target)
}

func Into(e error) (err CodeError, ok bool) {
	err, ok = e.(CodeError)
	if !ok {
		v, ok := _err_m[e]
		if !ok {
			return nil, false
		} else {
			return Code(v), true
		}
	}
	return err, ok
}

func MustInto(e error) CodeError {
	err, ok := Into(e)
	if !ok {
		panic(_must_into_error)
	}
	return err
}

func Cause(e error) CodeError {
	return MustInto(e)
}

func intInto(code int) codeError {
	return codeError(code)
}

func getMessageMap() MsgTable {
	m, ok := _messages.Load().(MsgTable)
	if !ok {
		panic(_not_init_error)
	}
	return m
}

func hasCode(code int) bool {
	_, ok := _codes[code]
	return ok
}
