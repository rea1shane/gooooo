package error

import (
	"fmt"
	"github.com/morikuni/failure"
	"net/http"
	"testing"
)

func TestWrapError(t *testing.T) {
	fmt.Printf("%+v\n", a())
}

func request() error {
	_, err := http.Get("http://localhost:8888")
	return failure.Wrap(err)
}

const (
	notFound failure.StringCode = "NotFound"
)

func a() error {
	return failure.Wrap(&aError{
		msg: "A",
		err: b(),
	})
}

func b() error {
	return failure.Wrap(&bError{
		msg: "B",
		err: c(),
	})
}

func c() error {
	return failure.New(notFound)
}

// aError 包装接收到的其他 error
type aError struct {
	msg string
	err error
}

func (a *aError) Unwrap() error {
	return a.err
}

func (a *aError) Error() string {
	if a.err != nil {
		return a.msg + " : " + a.err.Error()
	}
	return a.msg
}

// bError 包装接收到的其他 error
type bError struct {
	msg string
	err error
}

func (b *bError) Unwrap() error {
	return b.err
}

func (b *bError) Error() string {
	if b.err != nil {
		return b.msg + " : " + b.err.Error()
	}
	return b.msg
}
