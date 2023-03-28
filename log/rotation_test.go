package log

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"testing"
)

func TestRotate(t *testing.T) {
	logger := GetLogger()
	logger.Info("测试")
	Rotate(logger, &lumberjack.Logger{
		Filename: "logs/test.log",
	})
	logger.Info("test")
}

func TestLoadRotationConfigFromYaml(t *testing.T) {
	output, err := LoadRotationConfigFromYaml("lumberjack.yaml")
	if err != nil {
		panic(err)
	}
	logger := GetLogger()
	logger.SetOutput(output)
	logger.Info("test")
}
