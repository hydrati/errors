package main

import (
	"context"
	"fmt"

	"github.com/hyroge/errors"
)

var (
	RouteNotFoundError = errors.Code(404)
	RouteServerError   = errors.CodeWithMsg(500, "server error")
)

func init() {
	errors.Register(errors.MsgTable{
		200: "Ok",
		500: "server error!",
	})
	errors.SetCodeMsg(404, "Hello")
}

func main() {
	fmt.Println(
		RouteNotFoundError,
		RouteNotFoundError.Message(),
		RouteNotFoundError.Code(),
		RouteNotFoundError.Equals(RouteNotFoundError),
		RouteNotFoundError.Equals(errors.Code(404)),
		errors.Is(context.DeadlineExceeded, context.DeadlineExceeded),
	)
}
