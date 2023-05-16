package cron

import (
	"fmt"
	"github.com/rea1shane/gooooo/log"
	"github.com/robfig/cron/v3"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	formatter := log.GetFormatter()
	formatter.FieldsOrder = []string{"module"}
	logger := log.NewLogger()
	logger.SetFormatter(formatter)
	cronLogger := GenerateLogger(logger, []string{
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
	cronLogger.RegisterEntry(entryID, "PER SEC")
	c.Start()
	time.Sleep(100 * time.Second)
	c.Stop()
}
