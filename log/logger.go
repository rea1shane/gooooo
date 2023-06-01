package log

import (
	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

// NewLogger 新建一个 logrus.Logger，使用基础款 Formatter
func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(NewFormatter())
	return logger
}

// NewFormatter 新建一个带有自定义配置的 logrus.Formatter
func NewFormatter() *formatter.Formatter {
	return &formatter.Formatter{
		TimestampFormat: "2006-01-02 | 15:04:05",
		HideKeys:        true,
	}
}
