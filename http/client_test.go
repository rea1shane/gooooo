package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rea1shane/gooooo/data"
	"net/http"
	"testing"
)

type JsonData struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}

type YamlData struct {
	A string `yaml:"a"`
	B string `yaml:"b"`
	C string `yaml:"c"`
}

// TestServer 启动一个用于获取数据的服务。
func TestServer(t *testing.T) {
	handler := gin.Default()
	handler.GET("/json", func(c *gin.Context) {
		c.Writer.WriteString(`{"a": "1", "b": "2", "c": "3"}`)
	})
	handler.GET("/yaml", func(c *gin.Context) {
		c.Writer.WriteString(`a: 4
b: 5
c: 6`)
	})
	handler.Run("0.0.0.0:7777")
}

func TestDo(t *testing.T) {
	var jsonData JsonData
	r1, err := http.NewRequest(http.MethodGet, "http://localhost:7777/json", nil)
	if err != nil {
		panic(err)
	}
	err = DoAndUnmarshal(http.DefaultClient, r1, &jsonData, data.JsonFormat)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", jsonData)

	var yamlData YamlData
	r2, err := http.NewRequest(http.MethodGet, "http://localhost:7777/yaml", nil)
	if err != nil {
		panic(err)
	}
	err = DoAndUnmarshal(http.DefaultClient, r2, &yamlData, data.YamlFormat)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", yamlData)
}
