package error

import (
	"fmt"
	"github.com/morikuni/failure"
	"net/http"
	"testing"
)

func TestWrapError(t *testing.T) {
	fmt.Printf("%+v", httpRequest())
}

func httpRequest() error {
	_, err := http.Get("http://localhost:8888")
	return failure.Wrap(err)
}
