package slices

import (
	"fmt"
	"golang.org/x/exp/slices"
	"testing"
)

func TestContains(t *testing.T) {
	ss := []string{"A", "B", "C"}
	fmt.Println(slices.Contains(ss, "A"))
	is := []int{1, 2, 3}
	fmt.Println(slices.Contains(is, 1))
	fmt.Println(slices.Contains(is, 4))
}
