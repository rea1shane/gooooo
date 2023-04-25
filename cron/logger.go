package cron

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"time"
)

// GenerateCronLogger 通过 logrus.Logger 生成 cron.CronLogger
func GenerateCronLogger(logger *logrus.Logger, hiddenFields []string) *CronLogger {
	return &CronLogger{
		logger:       logger,
		entries:      make(map[cron.EntryID]string),
		hiddenFields: hiddenFields,
	}
}

type CronLogger struct {
	logger       *logrus.Logger
	entries      map[cron.EntryID]string // entries 记录每个 entryID 对应的名称
	hiddenFields []string                // hiddenFields cron 会返回 entry、now（当前时间）、next（下次调度时间）等 kv，数组中的 key 对应的 field 将被隐藏，如果不想隐藏任何信息就传入空
}

// RegisterEntry 记录任务对应的名称
func (cl *CronLogger) RegisterEntry(entryID cron.EntryID, entryName string) {
	cl.entries[entryID] = entryName
}

func (cl *CronLogger) Info(msg string, keysAndValues ...any) {
	entry := cl.logger.WithFields(cl.generateLoggerFields(keysAndValues...))
	entry.Info(msg)
}

func (cl *CronLogger) Error(err error, msg string, keysAndValues ...any) {
	entry := cl.logger.WithFields(cl.generateLoggerFields(keysAndValues...))
	entry.Error(msg, "err: ", err)
}

// generateLoggerFields 将 keysAndValues 转换为 logrus.Fields
func (cl *CronLogger) generateLoggerFields(kvs ...any) logrus.Fields {
	fields := make(logrus.Fields)
	fields["module"] = "cron"
	for i := 0; i < len(kvs); i += 2 {
		key := kvs[i].(string)
		value := kvs[i+1]
		if cl.hide(key) {
			continue
		}
		switch key {
		default:
			fields[key] = value
		case "entry":
			entryID := value.(cron.EntryID)
			if entryName, ok := cl.entries[entryID]; ok {
				fields[key] = entryName
			}
		case "now", "next":
			t := value.(time.Time)
			fields[key] = t.Format("2006-01-02 15:04:05")
		}
	}
	return fields
}

// hide 检测元素是否在 hiddenFields 中
func (cl *CronLogger) hide(key string) bool {
	for _, hiddenField := range cl.hiddenFields {
		if key == hiddenField {
			return true
		}
	}
	return false
}
