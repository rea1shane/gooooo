package http

import (
	"github.com/rea1shane/gooooo/data"
	"io"
	"net/http"
)

// Do 执行请求并且解析返回体数据加载到 model 中。
func Do(client *http.Client, request *http.Request, model any, format data.Format) error {
	// 请求
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// 读取
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// 解析
	return data.UnmarshalBytes(responseBody, model, format)
}
