# errors
 Golang Errors with Error Code

## Example
```golang
package main

import (
	"context"
	"fmt"

	"github.com/hyroge/errors"
)

var (
	RouteNotFoundError = errors.Create(404, "")
	RouteServerError   = errors.Create(500, "server error")
)

type ResponseError struct {
	Msg  string `coderror:"message"`
	Code int    `coderror:"code"`
}

func init() {

	errors.Set(404, "Hello")
	errors.Map(500, context.DeadlineExceeded)
}

func main() {
	err0 := new(ResponseError)
	errors.Cause(context.DeadlineExceeded).Inject(err0)
	fmt.Println(
		RouteNotFoundError,
		RouteNotFoundError.Message(),
		RouteNotFoundError.Code(),
		RouteNotFoundError.Equals(RouteNotFoundError),
		RouteNotFoundError.Equals(errors.Code(404)),
		errors.Is(context.DeadlineExceeded, context.DeadlineExceeded),
		errors.Cause(context.DeadlineExceeded),

		err0,
	)

	err := new(ResponseError)
	RouteNotFoundError.Inject(err)
	fmt.Println(err)

}


```