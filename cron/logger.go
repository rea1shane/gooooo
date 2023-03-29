package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

// GenerateCronLogger 从 logrus Logger 生成 cron Logger
func GenerateCronLogger(logger *log.Logger, hiddenFields []string) *CronLogger {
	return &CronLogger{
		logger:       logger,
		entries:      make(map[cron.EntryID]string),
		hiddenFields: hiddenFields,
	}
}

type CronLogger struct {
	logger       *log.Logger
	entries      map[cron.EntryID]string // entries 记录每个 entryID 对应的名称
	hiddenFields []string                // hiddenFields cron 会返回 entry、now（当前时间、next（下次调度时间）等 kv，数组中的 key 对应的 field 将被隐藏，如果不想隐藏任何信息就传入空
}

// RegisterEntry 记录任务对应的名称
func (cl *CronLogger) RegisterEntry(entryID cron.EntryID, entryName string) {
	cl.entries[entryID] = entryName
}

func (cl *CronLogger) Info(msg string, keysAndValues ...interface{}) {
	entry := cl.logger.WithFields(cl.generateLoggerFields(keysAndValues...))
	entry.Info(msg)
}

func (cl *CronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	entry := cl.logger.WithFields(cl.generateLoggerFields(keysAndValues...))
	entry.Error(msg, "err: ", err)
}

// generateLoggerFields 将 keysAndValues 转换为 logrus 的 Fields
func (cl *CronLogger) generateLoggerFields(kvs ...interface{}) log.Fields {
	fields := make(log.Fields)
	fields["module"] = "cron"
	for i := 0; i < len(kvs); i += 2 {
		key := fmt.Sprintf("%v", kvs[i])
		if cl.hide(key) {
			continue
		}
		if key == "entry" {
			entryID := kvs[i+1].(cron.EntryID)
			if entryName, ok := cl.entries[entryID]; ok {
				fields[key] = entryName
				continue
			}
		}
		fields[key] = kvs[i+1]
	}
	return fields
}

// hide 如果命中了 hiddenFields 中的元素，则不添加到 Fields 中
func (cl *CronLogger) hide(key string) bool {
	for _, hiddenField := range cl.hiddenFields {
		if key == hiddenField {
			return true
		}
	}
	return false
}
