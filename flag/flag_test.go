package flag

import (
	"flag"
	"fmt"
	"reflect"
	"testing"
)

var (
	a = flag.Int("a", 99, "这是 A")
	b string
)

func TestFlag(t *testing.T) {
	flag.StringVar(&b, "b", "value", "这是 B")
	flag.Parse()
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(*a)
	fmt.Println(b)
}
