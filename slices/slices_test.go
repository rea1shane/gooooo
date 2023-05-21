package slices

import (
	"fmt"
	"golang.org/x/exp/slices"
	"testing"
)

var (
	strings = []string{"A", "B", "C"}
	ints    = []int{1, 2, 3}
)

func TestContains(t *testing.T) {
	fmt.Println(slices.Contains(strings, "A"))
	fmt.Println(slices.Contains(ints, 1))
	fmt.Println(slices.Contains(ints, 4))
}

func TestConvertToMap(t *testing.T) {
	m := Map(strings)
	_, ok := m["A"]
	fmt.Println(ok)
	_, ok = m["D"]
	fmt.Println(ok)
}
