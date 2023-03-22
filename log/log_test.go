package log

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
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

	// 更改 Formatter
	formatter := GetFormatter()
	formatter.HideKeys = false
	logger.SetFormatter(formatter)

	logger.
		WithField("key1", "value1").
		WithField("key2", "value2").
		Error("error")
}

func TestRotate(t *testing.T) {
	logger := GetLogger()
	logger.Info("测试")
	Rotate(logger, &lumberjack.Logger{
		Filename: "logs/test.log",
	})
	logger.Info("test")
}

func TestLoadRotateConfigFromYaml(t *testing.T) {
	output, err := LoadRotateConfigFromYaml("lumberjack.yaml")
	if err != nil {
		panic(err)
	}
	logger := GetLogger()
	logger.SetOutput(output)
	logger.Info("test")
}
