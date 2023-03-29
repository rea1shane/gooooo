package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"testing"
	"time"
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
	c.AddFunc("* * * * * *", func() { fmt.Println("Print per second") })
	c.Start()
	time.Sleep(100 * time.Second)
	c.Stop()
}
