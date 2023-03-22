package log

import (
	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/rea1shane/gooooo/yaml"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// GetLogger 返回 Logger，使用基础款 Formatter
func GetLogger() *log.Logger {
	logger := log.New()
	logger.SetFormatter(GetFormatter())
	return logger
}

// GetFormatter 返回 Formatter 基础款
// 会隐藏 field list 的 key
func GetFormatter() *formatter.Formatter {
	return &formatter.Formatter{
		TimestampFormat: "2006-01-02 | 15:04:05",
		HideKeys:        true,
	}
}

// Rotate 设置 logger 滚动输出日志文件
func Rotate(logger *log.Logger, output *lumberjack.Logger) {
	logger.SetOutput(output)
}

// LoadRotateConfigFromYaml 从 yaml 文件中获取日志滚动相关配置
func LoadRotateConfigFromYaml(path string) (output *lumberjack.Logger, err error) {
	err = yaml.Load(path, &output)
	return
}
