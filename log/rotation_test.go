package log

import (
	"github.com/rea1shane/gooooo/data"
	"gopkg.in/natefinch/lumberjack.v2"
	"testing"
)

func TestRotate(t *testing.T) {
	logger := NewLogger()
	logger.Info("测试")
	Rotate(logger, &lumberjack.Logger{
		Filename: "logs/test.log",
	})
	logger.Info("test")
}

func TestNewRotationConfigFromFile(t *testing.T) {
	output, err := NewRotationConfigFromFile("rotation.yaml", data.YamlFormat)
	if err != nil {
		panic(err)
	}
	logger := NewLogger()
	logger.SetOutput(output)
	logger.Info("test")
}
