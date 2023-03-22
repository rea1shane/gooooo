package log

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestGetLogger(t *testing.T) {
	logger := GetLogger()

	logger.SetLevel(log.TraceLevel)

	logger.Trace("trace")
	logger.Debug("debug")

	// 携带 field list
	fields := make(log.Fields)
	fields["k1"] = "v1"
	fields["k2"] = "v2"
	loggerWithFields := logger.WithFields(fields)

	loggerWithFields.Info("info")
	loggerWithFields.Warn("warn")

	logger.
		WithField("key1", "value1").
		WithField("key2", "value2").
		Error("error")
}
