package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rea1shane/gooooo/log"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestNewHandler(t *testing.T) {
	logger := logrus.New()
	formatter := log.GetFormatter()
	formatter.FieldsOrder = []string{"StatusCode", "Latency"}
	logger.SetFormatter(formatter)
	handler := NewHandler(logger, 10*time.Microsecond, "/skip")
	run(handler)
}

func TestGinDefault(t *testing.T) {
	run(gin.Default())
}

func run(handler *gin.Engine) {
	handler.GET("/ping", func(c *gin.Context) {
		c.Writer.WriteString("pong")
		e := c.Error(errors.New("测试错误"))
		e.Meta = "测试元数据"
	})
	handler.GET("/skip", func(c *gin.Context) {
		c.Writer.WriteString("skip")
	})
	handler.Run("0.0.0.0:7777")
}
