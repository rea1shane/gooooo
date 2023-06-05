package log

import (
	"github.com/pkg/errors"
	"github.com/rea1shane/gooooo/data"
	"github.com/rea1shane/gooooo/os"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Rotate 设置 logrus.Logger 滚动输出日志文件
func Rotate(logger *logrus.Logger, output *lumberjack.Logger) {
	logger.SetOutput(output)
}

// NewRotationConfigFromFile 从文件中新建日志滚动配置
func NewRotationConfigFromFile(path string, format data.Format) (output *lumberjack.Logger, err error) {
	if format != data.JsonFormat && format != data.YamlFormat {
		return nil, errors.New("unsupported format, only JSON or YAML formats are supported")
	}
	err = os.Load(path, &output, format)
	return
}

// TODO 定时滚动 https://github.com/natefinch/lumberjack/issues/111
