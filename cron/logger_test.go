package cron

import (
	"fmt"
	"testing"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/rea1shane/gooooo/log"
)

func TestLogger(t *testing.T) {
	formatter := log.NewFormatter()
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
