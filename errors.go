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

func CodeWithMsg(code int, message string) CodeError {
	err := Code(code)
	SetCodeMsg(code, message)
	return err
}

func SetCodeMsg(code int, message string) {
	m := getMessageMap()
	m[code] = message
}

func Code(code int) CodeError {
	if hasCode(code) {
		panic(fmt.Sprintf("error: has err code %d", code))
	}
	return intInto(code)
}

func Register(kv MsgTable) {
	_messages.Store(kv)
}

func init() {
	_messages.Store(make(MsgTable))
}
