package log

import (
	"github.com/rea1shane/gooooo/yaml"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Rotate 设置 logrus.Logger 滚动输出日志文件
func Rotate(logger *logrus.Logger, output *lumberjack.Logger) {
	logger.SetOutput(output)
}

// LoadRotationConfigFromYaml 从 yaml 文件中获取日志滚动相关配置
func LoadRotationConfigFromYaml(path string) (output *lumberjack.Logger, err error) {
	err = yaml.Load(path, &output)
	return
}

// TODO 定时滚动 https://github.com/natefinch/lumberjack/issues/111
