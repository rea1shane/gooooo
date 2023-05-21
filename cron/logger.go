package cron

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
	"time"
)

// GenerateLogger 通过 logrus.Logger 生成 Logger
func GenerateLogger(logger *logrus.Logger, hiddenFields []string) *Logger {
	return &Logger{
		logger:       logger,
		entries:      make(map[cron.EntryID]string),
		hiddenFields: hiddenFields,
	}
}

// Logger 实现了 cron.Logger 接口
type Logger struct {
	logger       *logrus.Logger
	entries      map[cron.EntryID]string // entries 记录每个 entryID 对应的名称
	hiddenFields []string                // hiddenFields cron 会返回 entry、now（当前时间）、next（下次调度时间）等 kv，数组中的 key 对应的 field 将被隐藏，如果不想隐藏任何信息就传入空
}

// RegisterEntry 记录任务对应的名称
func (l *Logger) RegisterEntry(entryID cron.EntryID, entryName string) {
	l.entries[entryID] = entryName
}

func (l *Logger) Info(msg string, keysAndValues ...any) {
	entry := l.logger.WithFields(l.generateLoggerFields(keysAndValues...))
	entry.Info(msg)
}

func (l *Logger) Error(err error, msg string, keysAndValues ...any) {
	entry := l.logger.WithFields(l.generateLoggerFields(keysAndValues...))
	entry.Error(msg, "err: ", err)
}

// generateLoggerFields 将 keysAndValues 转换为 logrus.Fields
func (l *Logger) generateLoggerFields(kvs ...any) logrus.Fields {
	fields := make(logrus.Fields)
	fields["module"] = "cron"
	for i := 0; i < len(kvs); i += 2 {
		key := kvs[i].(string)
		value := kvs[i+1]
		// 检测元素是否在 hiddenFields 中
		if slices.Contains(l.hiddenFields, key) {
			continue
		}
		switch key {
		default:
			fields[key] = value
		case "entry":
			entryID := value.(cron.EntryID)
			if entryName, ok := l.entries[entryID]; ok {
				fields[key] = entryName
			}
		case "now", "next":
			t := value.(time.Time)
			fields[key] = t.Format("2006-01-02 15:04:05")
		}
	}
	return fields
}
