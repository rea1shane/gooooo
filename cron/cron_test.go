package cron

import (
	"fmt"
	"testing"
	"time"

	"github.com/robfig/cron/v3"
)

func TestCron(t *testing.T) {
	c := cron.New()
	c.AddFunc("* * * * *", func() { fmt.Println("Print per minute") })
	c.Start()
	time.Sleep(100 * time.Second)
	c.Stop() // 不会停止已经运行的任务
}

func TestCronWithSeconds(t *testing.T) {
	c := cron.New(cron.WithSeconds())
	id, err := c.AddFunc("* * * * * *", func() { fmt.Println("Print per second") })
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
	c.Start()
	time.Sleep(100 * time.Second)
	c.Stop()
}
