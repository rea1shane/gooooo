package yaml

import (
	"fmt"
	"testing"
)

type Test struct {
	A string `yaml:"a"`
	B string `yaml:"b"`
	C string `yaml:"c"`
}

func TestLoad(t *testing.T) {
	var test *Test
	err := Load("test.yaml", &test)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%+v", test))
}
