package cron

import (
	"fmt"
	"github.com/rea1shane/gooooo/log"
	"github.com/robfig/cron/v3"
	"testing"
	"time"
)

func TestCronLogger(t *testing.T) {
	cronLogger := GenerateCronLogger(log.GetLogger(), []string{
		"now",
		//"next",
	})
	c := cron.New(
		cron.WithSeconds(),
		cron.WithLogger(cronLogger),
	)
	entryID, err := c.AddFunc("* * * * * *", func() { fmt.Println("Print per second") })
	if err != nil {
		panic(err)
	}
	cronLogger.RecordEntry(entryID, "PER SEC")
	c.Start()
	time.Sleep(100 * time.Second)
	c.Stop()
}
