package errors

type none = struct{}
type any = interface{}

type CodeError interface {
	Code() int
	Message() string
	Equals(error) bool
	Error() string
	Inject(val interface{})
}

type MsgTable = map[int]string

type codeError int
