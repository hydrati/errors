package errors

import (
	"errors"
	"fmt"
	"strconv"
)

func (e codeError) Code() int {
	return int(e)
}

func (e codeError) Error() string {
	return strconv.FormatInt(int64(e.Code()), 10)
}

func (e codeError) Message() string {
	if msg, ok := getMessageMap()[e.Code()]; ok {
		return msg
	} else {
		return DEFAULT_ERROR_MESSAGE
	}
}

func (e codeError) Equals(a error) bool {
	if s, ok := Into(a); ok {
		if s.Code() == e.Code() {
			return true
		}
	}
	return false
}

func (e codeError) ToString() string {
	return fmt.Sprintf("error(code=%d)", e.Code())
}

func Equals(left error, right error) bool {
	l, ok := Into(left)
	r, ok2 := Into(right)
	if !ok || !ok2 {
		return errors.Is(left, right)
	}
	return l.Equals(r)
}
