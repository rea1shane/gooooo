package log

import (
	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

// NewLogger 返回一个新建的 logrus.Logger，使用基础款 Formatter
func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(GetFormatter())
	return logger
}

// GetFormatter 返回基础款 logrus.Formatter
func GetFormatter() *formatter.Formatter {
	return &formatter.Formatter{
		TimestampFormat: "2006-01-02 | 15:04:05",
		HideKeys:        true,
	}
}
