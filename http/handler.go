package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// NewHandler 新建一个使用指定 logrus.Logger 作为 logger 的 gin.Engine
func NewHandler(logger logrus.FieldLogger, latencyThreshold time.Duration, notLogged ...string) *gin.Engine {
	engine := gin.New()
	engine.Use(HandlerLogger(logger, latencyThreshold, notLogged...), gin.Recovery())
	return engine
}

// HandlerLogger 通过 logrus.Logger 生成 gin 日志中间层
// 代码参考 gin.LoggerWithConfig
func HandlerLogger(logger logrus.FieldLogger, latencyThreshold time.Duration, notLogged ...string) gin.HandlerFunc {
	var skip map[string]struct{}

	if length := len(notLogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notLogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			param := gin.LogFormatterParams{
				Request: c.Request,
				Keys:    c.Keys,
			}

			// Stop timer
			param.TimeStamp = time.Now()
			param.Latency = param.TimeStamp.Sub(start)

			param.ClientIP = c.ClientIP()
			param.Method = c.Request.Method
			param.StatusCode = c.Writer.Status()
			param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

			param.BodySize = c.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}

			param.Path = path

			log(logger, param, latencyThreshold)
		}
	}
}

// log 输出日志，包含两个 logrus.Fields："StatusCode" 和 "Latency"
// 根据 http code 设定日志等级，大于等于 500 是 error，大于等于 400 小于 500 是 warning，其余的都是 info。
// 当处理请求时长 latency 超过 latencyThreshold 时，info 等级的日志将转为 warning。latencyThreshold <= 0 时禁用此功能
func log(logger logrus.FieldLogger, param gin.LogFormatterParams, latencyThreshold time.Duration) {
	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}

	entry := logger.WithFields(logrus.Fields{
		"StatusCode": fmt.Sprintf("%3d", param.StatusCode),
		"Latency":    fmt.Sprintf("%13v", param.Latency),
	})

	msg := fmt.Sprintf("%15s | %-7s %#v",
		param.ClientIP,
		param.Method,
		param.Path,
	)
	if len(param.ErrorMessage) > 0 {
		msg = fmt.Sprintf("%s\n%s",
			msg,
			param.ErrorMessage[:len(param.ErrorMessage)-1],
		)
	}

	if param.StatusCode >= http.StatusInternalServerError {
		// 服务端错误
		entry.Error(msg)
	} else if param.StatusCode >= http.StatusBadRequest || (latencyThreshold > 0 && param.Latency > latencyThreshold) {
		// 客户端错误 || latency 超过阈值
		entry.Warn(msg)
	} else {
		entry.Info(msg)
	}
}
