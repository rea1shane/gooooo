package log

import (
	formatter "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
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
