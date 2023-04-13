package time

import (
	"fmt"
	"testing"
	"time"
)

func TestFormatDuration(t *testing.T) {
	d1, _ := time.ParseDuration("124h32m59s")
	d2, _ := time.ParseDuration("-1124h32m59s")
	fmt.Println(FormatDuration(d1))
	fmt.Println(FormatDuration(d2))
}
