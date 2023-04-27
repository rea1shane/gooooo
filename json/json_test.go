package json

import (
	"fmt"
	"testing"
)

type Test struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}

func TestLoad(t *testing.T) {
	var test *Test
	err := Load("test.json", &test)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%+v", test))
}
