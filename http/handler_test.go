package http

import (
	"github.com/gin-gonic/gin"
	"github.com/rea1shane/gooooo/log"
	"testing"
)

func TestNewHandler(t *testing.T) {
	logger := log.NewLogger()
	handler := NewHandler(logger, -1)
	run(handler)
}

func TestGinDefault(t *testing.T) {
	run(gin.Default())
}

func run(handler *gin.Engine) {
	handler.GET("/ping", func(c *gin.Context) {
		c.Writer.WriteString("pong")
		//c.Error(errors.New("测试错误"))
	})
	handler.Run("0.0.0.0:7777")
}
