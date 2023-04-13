package time

import (
	"fmt"
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	t1me := "2023-03-29 11:10:05.087582 +0800 CST"
	layout := "2006-01-02 15:04:05.000000 -0700 MST"
	tm, err := time.Parse(layout, t1me)
	if err != nil {
		panic(err)
	}
	fmt.Println(tm)
}

func TestSub(t *testing.T) {
	t1 := time.Now()
	t2 := time.Now()
	fmt.Println(t2.Sub(t1).String())
}
