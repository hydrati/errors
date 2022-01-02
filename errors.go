package errors

import (
	"errors"
	"fmt"
	"sync/atomic"
)

var (
	_messages atomic.Value // atomic.Value[map[int]string]
	_codes    = map[int]none{}
	_err_m    = map[error]int{}

	_must_into_error = errors.New("error: this is not a `CodeError`")
	_not_init_error  = errors.New("error: the message table is not initialized")
)

const (
	defaultErrMsg = "unknown error"
)

func Create(code int, message string) CodeError {
	if hasCode(code) {
		panic(fmt.Sprintf("error: has err code %d", code))
	}
	err := intInto(code)
	Set(code, message)
	return err
}

func Set(code int, message string) {
	m := getMessageMap()
	_codes[code] = none{}
	m[code] = message
}

func Map(code int, err error) int {
	if !hasCode(code) {
		panic(fmt.Sprintf("error: hasn't err code %d", code))
	}
	_err_m[err] = code
	return code
}

func Code(code int) CodeError {
	if !hasCode(code) {
		panic(fmt.Sprintf("error: hasn't err code %d", code))
	}
	return intInto(code)
}

func init() {
	_messages.Store(make(MsgTable))
}
