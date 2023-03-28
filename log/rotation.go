package log

import (
	"github.com/rea1shane/gooooo/yaml"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Rotate 设置 logger 滚动输出日志文件
func Rotate(logger *log.Logger, output *lumberjack.Logger) {
	logger.SetOutput(output)
}

// LoadRotationConfigFromYaml 从 yaml 文件中获取日志滚动相关配置
func LoadRotationConfigFromYaml(path string) (output *lumberjack.Logger, err error) {
	err = yaml.Load(path, &output)
	return
}
