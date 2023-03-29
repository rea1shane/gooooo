package cron

import (
	"fmt"
	"github.com/rea1shane/gooooo/log"
	"github.com/robfig/cron/v3"
	"testing"
	"time"
)

func TestCronLogger(t *testing.T) {
	formatter := log.GetFormatter()
	formatter.FieldsOrder = []string{"module"}
	logger := log.GetLogger()
	logger.SetFormatter(formatter)
	cronLogger := GenerateCronLogger(logger, []string{
		"now",
		"next",
	})
	c := cron.New(
		cron.WithSeconds(),
		cron.WithLogger(cronLogger),
	)
	entryID, err := c.AddFunc("* * * * * *", func() { fmt.Println("Print per second") })
	if err != nil {
		panic(err)
	}
	cronLogger.RegisterEntry(entryID, "per-sec")
	c.Start()
	time.Sleep(100 * time.Second)
	c.Stop()
}
