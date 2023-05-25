package log

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGetLogger(t *testing.T) {
	logger := NewLogger()

	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("trace", "A", "B")
	logger.Debug("debug")

	fields := make(logrus.Fields)
	fields["k1"] = "v1"
	fields["k2"] = "v2"
	loggerWithFields := logger.WithFields(fields)

	loggerWithFields.Info("info", "A", "B")
	loggerWithFields.Warn("warn")

	formatter := GetFormatter()
	formatter.HideKeys = false
	logger.SetFormatter(formatter)

	logger.
		WithField("key1", "value1").
		WithField("key2", "value2").
		Error("error")
}

func TestFatal(t *testing.T) {
	logger := NewLogger()
	logger.Fatal(errors.New("ERROR"))
	fmt.Println("end")
}
